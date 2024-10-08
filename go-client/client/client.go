package client

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"fast-updates-client/config"
	"fast-updates-client/contracts-interface/fast_updater"
	"fast-updates-client/contracts-interface/fast_updates_configuration"
	"fast-updates-client/contracts-interface/incentive"
	"fast-updates-client/contracts-interface/mock"
	"fast-updates-client/contracts-interface/submission"
	"fast-updates-client/contracts-interface/system_manager"
	"fast-updates-client/logger"
	"fast-updates-client/provider"
	"fast-updates-client/sortition"
)

type FastUpdatesClient struct {
	params              config.FastUpdateClientConfig
	chainClient         *ethclient.Client
	valuesProvider      provider.ValuesProvider
	signingAccount      *Account
	transactionAccounts []*Account
	fastUpdater         *fast_updater.FastUpdater
	fastUpdatesConfig   *fast_updates_configuration.FastUpdatesConfiguration
	submission          *submission.Submission
	flareSystemMock     *mock.Mock
	flareSystemManager  *system_manager.SystemManager
	IncentiveManager    *incentive.Incentive
	key                 *sortition.Key
	registeredEpochs    map[int64]bool
	transactionQueue    *TransactionQueue
	allFeeds            []provider.FeedId
	loggingParams       config.LoggerConfig
	Stats               UpdatesStats
}

type Account struct {
	Address    common.Address
	PrivateKey *ecdsa.PrivateKey
}

type UpdatesStats struct {
	NumUpdates           uint64
	NumSuccessfulUpdates uint64
}

const (
	refreshFeedsBlockInterval = 100
	cutoffRepeats             = 10
)

func CreateFastUpdatesClient(cfg *config.Config, valuesProvider provider.ValuesProvider) (*FastUpdatesClient, error) {
	fastUpdatesClient := FastUpdatesClient{}
	fastUpdatesClient.params = cfg.Client
	fastUpdatesClient.loggingParams = cfg.Logger
	fastUpdatesClient.valuesProvider = valuesProvider

	var err error
	fastUpdatesClient.chainClient, err = cfg.Chain.DialETH()
	if err != nil {
		return nil, fmt.Errorf("CreateFastUpdatesClient: Dial: %w", err)
	}

	fastUpdatesClient.signingAccount = &Account{}
	privateKey := cfg.Client.SigningPrivateKey
	if privateKey[:2] == "0x" {
		privateKey = privateKey[2:]
	}
	fastUpdatesClient.signingAccount.PrivateKey, err = crypto.HexToECDSA(privateKey)
	if err != nil {
		return nil, fmt.Errorf("CreateFastUpdatesClient: HexToECDSA: %w", err)
	}
	publicKey := fastUpdatesClient.signingAccount.PrivateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("CreateFastUpdatesClient: Error casting public key to ECDSA: %w", err)
	}
	fastUpdatesClient.signingAccount.Address = crypto.PubkeyToAddress(*publicKeyECDSA)

	fastUpdatesClient.transactionAccounts = make([]*Account, len(cfg.Transactions.Accounts))
	for i, accountPrivateKey := range cfg.Transactions.Accounts {
		fastUpdatesClient.transactionAccounts[i] = &Account{}

		privateKey := accountPrivateKey
		if privateKey[:2] == "0x" {
			privateKey = privateKey[2:]
		}
		fastUpdatesClient.transactionAccounts[i].PrivateKey, err = crypto.HexToECDSA(privateKey)
		if err != nil {
			return nil, fmt.Errorf("CreateFastUpdatesClient: HexToECDSA: %w", err)
		}
		publicKey := fastUpdatesClient.transactionAccounts[i].PrivateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			return nil, fmt.Errorf("CreateFastUpdatesClient: Error casting public key to ECDSA: %w", err)
		}
		fastUpdatesClient.transactionAccounts[i].Address = crypto.PubkeyToAddress(*publicKeyECDSA)
	}

	if cfg.Client.FastUpdaterAddress != "" {
		fastUpdatesClient.fastUpdater, err = fast_updater.NewFastUpdater(
			common.HexToAddress(cfg.Client.FastUpdaterAddress), fastUpdatesClient.chainClient,
		)
		if err != nil {
			return nil, fmt.Errorf("CreateFastUpdatesClient: NewFastUpdater: %w", err)
		}
	}

	fastUpdatesClient.IncentiveManager, err = incentive.NewIncentive(
		common.HexToAddress(cfg.Client.IncentiveManagerAddress), fastUpdatesClient.chainClient,
	)
	if err != nil {
		return nil, fmt.Errorf("CreateFastUpdatesClient: NewIncentive: %w", err)
	}

	fastUpdatesClient.fastUpdatesConfig, err = fast_updates_configuration.NewFastUpdatesConfiguration(
		common.HexToAddress(cfg.Client.FastUpdatesConfigurationAddress), fastUpdatesClient.chainClient,
	)
	if err != nil {
		return nil, fmt.Errorf("CreateFastUpdatesClient: NewFastUpdatesConfiguration: %w", err)
	}

	if cfg.Client.SubmissionAddress != "" {
		fastUpdatesClient.submission, err = submission.NewSubmission(
			common.HexToAddress(cfg.Client.SubmissionAddress), fastUpdatesClient.chainClient,
		)
		if err != nil {
			return nil, fmt.Errorf("CreateFastUpdatesClient: NewSubmission: %w", err)
		}
	}

	err = fastUpdatesClient.UpdateFastUpdaterContractAddress()
	if err != nil {
		return nil, err
	}

	fastUpdatesClient.flareSystemMock, err = mock.NewMock(
		common.HexToAddress(cfg.Client.MockAddress), fastUpdatesClient.chainClient,
	)
	if err != nil {
		return nil, fmt.Errorf("CreateFastUpdatesClient: NewMock: %w", err)
	}
	fastUpdatesClient.flareSystemManager, err = system_manager.NewSystemManager(
		common.HexToAddress(cfg.Client.FlareSystemManagerAddress), fastUpdatesClient.chainClient,
	)
	if err != nil {
		return nil, fmt.Errorf("CreateFastUpdatesClient: NewSystemManager: %w", err)
	}

	if cfg.Client.SortitionPrivateKey == "" {
		fastUpdatesClient.key, err = sortition.KeyGen()
		if err != nil {
			return nil, fmt.Errorf("CreateFastUpdatesClient: KeyGen: %w", err)
		}
		logger.Info("generated new private sortition key: %s", "0x"+fastUpdatesClient.key.Sk.Text(16))
	} else {
		fastUpdatesClient.key, err = sortition.KeyFromString(cfg.Client.SortitionPrivateKey)
		if err != nil {
			return nil, fmt.Errorf("CreateFastUpdatesClient: KeyGen: %w", err)
		}
	}

	fastUpdatesClient.registeredEpochs = make(map[int64]bool)

	cfg.Transactions.ChainId = big.NewInt(int64(cfg.Chain.ChainId))
	fastUpdatesClient.transactionQueue = NewTransactionQueue(fastUpdatesClient.chainClient, fastUpdatesClient.transactionAccounts, cfg.Transactions)
	go fastUpdatesClient.transactionQueue.Run()

	if fastUpdatesClient.params.MaxWeight == 0 {
		fastUpdatesClient.params.MaxWeight = 4096
	}

	return &fastUpdatesClient, nil
}

func (client *FastUpdatesClient) Run(startBlock, endBlock uint64) error {
	var blockNum uint64
	var err error
	if startBlock != 0 {
		blockNum = startBlock
	} else {
		blockNum, err = client.CurrentBlockNumber()
		if err != nil {
			return fmt.Errorf("Run: CurrentBlockNumber: %w", err)
		}
	}

	epoch, err := client.GetCurrentRewardEpochId()
	if err != nil {
		return fmt.Errorf("Run: GetCurrentRewardEpochId: %w", err)
	}

	seed, err := client.GetSeed(epoch)
	if err != nil {
		return fmt.Errorf("Run: GetCurrentSeed: %w", err)
	}

	weight, err := client.GetMyWeight()
	if err != nil {
		logger.Error("error getting weight %s", fmt.Errorf("Run: GetMyWeight: %w", err))
	}
	weight = min(weight, uint64(client.params.MaxWeight))
	logger.Info("staring block %d, epoch %d, my weight %d", blockNum, epoch, weight)

	client.allFeeds, err = client.GetCurrentFeedIds()
	if err != nil {
		return fmt.Errorf("Run: GetCurrentFeedIds: %w", err)
	}
	logger.Info("Fetched feed ids: %v", client.allFeeds)

	for {
		err = client.UpdateFastUpdaterContractAddress()
		if err != nil {
			logger.Error("Failed attempt in updating the FastUpdater contract address: %s", err)
		}

		if blockNum%refreshFeedsBlockInterval == 0 {
			client.allFeeds, err = client.GetCurrentFeedIds()
			if err != nil {
				return fmt.Errorf("Run: GetCurrentFeedIds: %w", err)
			}
		}

		epochCheck, err := client.GetCurrentRewardEpochId()
		if err != nil {
			return fmt.Errorf("Run: GetCurrentRewardEpochId: %w", err)
		}

		// todo: this check should be changed to something that happens a few seconds before the epoch changes
		// to avoid failed transactions
		if epochCheck > epoch {
			client.transactionQueue.EmptyQueue()

			blockNum, err = client.CurrentBlockNumber()
			if err != nil {
				return fmt.Errorf("Run: CurrentBlockNumber: %w", err)
			}
			epoch = epochCheck
			weight, err = client.GetMyWeight()
			if err != nil {
				logger.Error("error getting weight %s", fmt.Errorf("Run: GetMyWeight: %w", err))
			}
			weight = min(weight, uint64(client.params.MaxWeight))
			if weight == 0 {
				return fmt.Errorf("Run: Not registered in epoch: %d", epoch)
			}

			seed, err = client.GetSeed(epoch)
			if err != nil {
				return fmt.Errorf("Run: GetCurrentSeed: %w", err)
			}
			logger.Info("new epoch, my weight weight %d, current block %d", weight, blockNum)
		}

		cutoff, err := client.GetBlockScoreCutoffWithRepeats(blockNum)
		if err != nil {
			return fmt.Errorf("Run: GetBlockScoreCutoffWithRepeats: %w", err)
		}

		updateProofs, err := sortition.FindUpdateProofs(client.key, seed, cutoff, big.NewInt(int64(blockNum)), weight)
		if err != nil {
			return fmt.Errorf("Run: FindUpdateProofs: %w", err)
		}
		for _, updateProof := range updateProofs {
			logger.Info("scheduling update for block %d replicate %d", updateProof.BlockNumber, updateProof.Replicate)
			client.SubmitUpdates(updateProof)
			client.Stats.NumUpdates++
		}

		if len(updateProofs) > 0 {
			balances, err := client.GetBalances()
			if err != nil {
				logger.Error("could not obtain balances: %s", err)
			}
			if !CheckBalances(balances, client.loggingParams.MinBalance) {
				logger.Warn("account balance low: %s", balances)
			}
		}

		if client.loggingParams.FeedValuesLog != 0 && blockNum%uint64(client.loggingParams.FeedValuesLog) == 0 {
			_, chainValues, providerValues, err := client.getOnlineOfflineValues()
			if err != nil {
				logger.Error("failed obtaining feed values %s", err)
			} else {
				logger.Info("chain feeds values in block %d: %v", blockNum, chainValues)
				logger.Info("provider feeds values: %v", providerValues)
			}
		}

		// do not calculate in advance more than specified
		err = WaitForBlock(client.transactionQueue, blockNum)
		if err != nil {
			return fmt.Errorf("Run: WaitForBlock: %w", err)
		}
		blockNum++
		if endBlock != 0 && blockNum > endBlock {
			client.transactionQueue.WaitToEmptyQueue()
			return nil
		}
	}
}

func (client *FastUpdatesClient) GetBlockScoreCutoffWithRepeats(blockNum uint64) (*big.Int, error) {
	var cutoff *big.Int
	var err error
	for i := 0; i < cutoffRepeats; i++ {
		cutoff, err = client.GetBlockScoreCutoff(big.NewInt(int64(blockNum)))
		if err == nil {
			break
		}
		if blockNum <= client.transactionQueue.CurrentBlockNum-uint64(client.params.SubmissionWindow)+1 || i == cutoffRepeats-1 {
			return nil, fmt.Errorf("Run: GetCurrentScoreCutoff for block %d: %w", blockNum, err)
		}
		time.Sleep(200 * time.Millisecond)
	}

	return cutoff, err
}

func (client *FastUpdatesClient) WaitToEmptyRequests() {
	client.transactionQueue.WaitToEmptyQueue()
}

func (client *FastUpdatesClient) Stop() {
	client.transactionQueue.StopQueue()
}

func (client *FastUpdatesClient) UpdateFastUpdaterContractAddress() error {
	if client.submission == nil {
		return nil
	}

	newAddress, err := client.GetFastUpdaterContractAddress()
	if err != nil {
		return err
	}
	if newAddress != common.HexToAddress(client.params.FastUpdaterAddress) {
		client.fastUpdater, err = fast_updater.NewFastUpdater(
			newAddress, client.chainClient,
		)
		if err != nil {
			return err
		}

		client.params.FastUpdaterAddress = newAddress.Hex()
		logger.Info("Updated the FastUpdater address to %s", client.params.FastUpdaterAddress)
	}

	return nil
}
