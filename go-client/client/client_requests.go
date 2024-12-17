package client

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"

	"fast-updates-client/config"
	"fast-updates-client/contracts-interface/mock"
	"fast-updates-client/logger"
	"fast-updates-client/provider"
	"fast-updates-client/sortition"
	"fast-updates-client/updates"
)

const fetchCurrentFeedsMethod = "fetchCurrentFeeds"

func (client *FastUpdatesClient) CurrentBlockNumber() (uint64, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(config.CallTimeoutMillisDefault)*time.Millisecond)
	currentBlockNum, err := client.chainClient.BlockNumber(ctx)
	cancelFunc()

	return currentBlockNum, err
}

func (client *FastUpdatesClient) GetCurrentScoreCutoff() (*big.Int, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(config.CallTimeoutMillisDefault)*time.Millisecond)
	ops := &bind.CallOpts{Context: ctx}
	score, err := client.fastUpdater.CurrentScoreCutoff(ops)
	cancelFunc()

	return score, err
}

func (client *FastUpdatesClient) GetBlockScoreCutoff(blockNum *big.Int) (*big.Int, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(config.CallTimeoutMillisDefault)*time.Millisecond)
	ops := &bind.CallOpts{Context: ctx}
	score, err := client.fastUpdater.BlockScoreCutoff(ops, blockNum)
	cancelFunc()

	return score, err
}

func (client *FastUpdatesClient) GetSeed(rewardEpochId int64) (*big.Int, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(config.CallTimeoutMillisDefault)*time.Millisecond)
	ops := &bind.CallOpts{Context: ctx}
	seed, err := client.flareSystemManager.GetSeed(ops, new(big.Int).SetInt64(rewardEpochId))
	cancelFunc()

	return seed, err
}

func (client *FastUpdatesClient) GetScale() (*big.Int, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(config.CallTimeoutMillisDefault)*time.Millisecond)
	ops := &bind.CallOpts{Context: ctx}
	scale, err := client.IncentiveManager.GetScale(ops)
	cancelFunc()

	return scale, err
}

func (client *FastUpdatesClient) GetExpectedSampleSize() (*big.Int, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(config.CallTimeoutMillisDefault)*time.Millisecond)
	ops := &bind.CallOpts{Context: ctx}
	sampleSize, err := client.IncentiveManager.GetExpectedSampleSize(ops)
	cancelFunc()

	return sampleSize, err
}

func (client *FastUpdatesClient) GetCurrentRewardEpochId() (int64, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(config.CallTimeoutMillisDefault)*time.Millisecond)
	ops := &bind.CallOpts{Context: ctx}
	epoch, err := client.flareSystemManager.GetCurrentRewardEpochId(ops)
	cancelFunc()
	if err != nil {
		return 0, err
	}

	return epoch.Int64(), nil
}

func (client *FastUpdatesClient) GetMyWeight() (uint64, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(config.CallTimeoutMillisDefault)*time.Millisecond)
	ops := &bind.CallOpts{Context: ctx}
	weight, err := client.fastUpdater.CurrentSortitionWeight(ops, client.signingAccount.Address)
	cancelFunc()
	if err != nil || weight == nil {
		return 0, err
	}

	return weight.Uint64(), nil
}

func (client *FastUpdatesClient) callContractFetchCurrentFeeds(ctx context.Context, feedIndexes []*big.Int) (*provider.ValuesDecimals, error) {
	fastUpdatesAddress := common.HexToAddress(client.params.FastUpdaterAddress)

	input, err := client.fastUpdaterABI.Pack(fetchCurrentFeedsMethod, feedIndexes)
	if err != nil {
		return nil, err
	}

	msg := ethereum.CallMsg{
		To:    &fastUpdatesAddress,
		Data:  input,
		Value: client.FetchCurrentFeedsValue,
		From:  client.FetchCurrentFeedsAddress,
	}
	outputBytes, err := client.chainClient.CallContract(ctx, msg, nil)
	if err != nil {
		return nil, err
	}

	res, err := client.fastUpdaterABI.Unpack(fetchCurrentFeedsMethod, outputBytes)
	if err != nil {
		return nil, err
	}
	if len(res) != 3 {
		return nil, fmt.Errorf("ABI unpack error, length of response %d", len(res))
	}

	return &provider.ValuesDecimals{
		Feeds:     *abi.ConvertType(res[0], new([]*big.Int)).(*[]*big.Int),
		Decimals:  *abi.ConvertType(res[1], new([]int8)).(*[]int8),
		Timestamp: *abi.ConvertType(res[2], new(uint64)).(*uint64),
	}, nil
}

func (client *FastUpdatesClient) GetFeeds(feedIndexes []int) ([]float64, uint64, error) {
	feedIndexesBigInt := make([]*big.Int, len(feedIndexes))
	for i, index := range feedIndexes {
		feedIndexesBigInt[i] = big.NewInt(int64(index))
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(config.CallTimeoutMillisDefault)*time.Millisecond)
	feedValues, err := client.callContractFetchCurrentFeeds(ctx, feedIndexesBigInt)
	cancelFunc()
	if err != nil {
		return nil, 0, err
	}

	floatValues := RawChainValuesToFloats(*feedValues)
	return floatValues, feedValues.Timestamp, err
}

func (client *FastUpdatesClient) GetCurrentFeedIds() ([]provider.FeedId, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(config.CallTimeoutMillisDefault)*time.Millisecond)
	ops := &bind.CallOpts{Context: ctx}
	rawFeedIs, err := client.fastUpdatesConfig.GetFeedIds(ops)
	cancelFunc()
	if err != nil {
		return nil, errors.Wrap(err, "error getting feed ids")
	}

	feedIds := make([]provider.FeedId, len(rawFeedIs))
	for i, price := range rawFeedIs {
		feedIds[i] = provider.FeedId{
			Category: price[0],
			Name:     strings.TrimRight(string(price[1:]), "\x00"),
		}
	}

	return feedIds, nil
}

func (client *FastUpdatesClient) Register(epoch int64) {
	compReq := ComputationRequest{Function: func(txOpts *bind.TransactOpts) error { return client.register(epoch, txOpts) }}
	client.transactionQueue.InputChan <- compReq
}

func (client *FastUpdatesClient) GetBalances() ([]*big.Int, error) {
	balances := make([]*big.Int, len(client.transactionAccounts))
	for i, account := range client.transactionAccounts {
		balance, err := client.chainClient.BalanceAt(context.Background(),
			account.Address, nil)
		if err != nil {
			return nil, err
		}

		balances[i] = balance
	}

	return balances, nil
}

// only on mocked
func (client *FastUpdatesClient) register(epoch int64, txOpts *bind.TransactOpts) error {
	policy := mock.FlareSystemMockPolicy{Pk1: client.key.Pk.X.Bytes(), Pk2: client.key.Pk.Y.Bytes(), Weight: uint16(1000)}

	logger.Info("registering %s for epoch %d", client.signingAccount.Address, epoch)
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(config.CallTimeoutMillisDefault)*time.Millisecond)
	txOpts.Context = ctx
	tx, err := client.flareSystemMock.RegisterAsVoter(txOpts, big.NewInt(int64(epoch)), client.signingAccount.Address, policy)
	cancelFunc()
	if err != nil {
		return err
	}

	ctx, cancelFunc = context.WithTimeout(context.Background(), time.Duration(config.TxTimeoutMillisDefault)*time.Millisecond)
	receipt, err := bind.WaitMined(ctx, client.chainClient, tx)
	cancelFunc()
	if err != nil {
		return err
	}
	if receipt.Status == 0 {
		return fmt.Errorf("transaction failed")
	}
	client.registeredEpochs[epoch] = true

	return nil
}

func (client *FastUpdatesClient) SubmitUpdates(updateProof *sortition.UpdateProof) {
	beforeBlock := updateProof.BlockNumber.Uint64() + uint64(client.params.SubmissionWindow) - 1
	compReq := ComputationRequest{
		Function: func(txOpts *bind.TransactOpts) error {
			return client.submitUpdates(updateProof, txOpts)
		},
		BeforeBlock: beforeBlock,
		AfterBlock:  updateProof.BlockNumber.Uint64(),
	}

	client.transactionQueue.InputChan <- compReq
}

func (client *FastUpdatesClient) getOnlineOfflineValues() ([]int, []float64, []float64, error) {
	providerRawValues, err := client.valuesProvider.GetValues(client.allFeeds)
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "error getting feed values")
	}

	supportedFeedIndexes := []int{}
	providerValues := []float64{}
	for i, value := range providerRawValues {
		if value != nil {
			supportedFeedIndexes = append(supportedFeedIndexes, i)
			providerValues = append(providerValues, *value)
		}
	}

	// get current prices from on-chain
	chainValues, _, err := client.GetFeeds(supportedFeedIndexes)
	if err != nil {
		return nil, nil, nil, err
	}

	return supportedFeedIndexes, chainValues, providerValues, nil
}

func (client *FastUpdatesClient) submitUpdates(updateProof *sortition.UpdateProof, txOpts *bind.TransactOpts) error {
	// get feed values from providers and calculate deltas relative to on-chain values
	supportedFeedIndexes, chainValues, providerValues, err := client.getOnlineOfflineValues()
	if err != nil {
		return err
	}

	// get current expectedSampleSize
	sampleSize, err := client.GetExpectedSampleSize()
	if err != nil {
		return err
	}

	logger.Info("chain feeds values in block %d (before update): %v", client.transactionQueue.CurrentBlockNum, chainValues)
	logger.Info("provider feeds values: %v", providerValues)

	// get current scale
	scale, err := client.GetScale()
	if err != nil {
		return err
	}

	// calculate deltas for provider and on-chain prices
	deltas, deltasString, err := provider.GetDeltas(chainValues, providerValues, supportedFeedIndexes, scale, sampleSize)
	if err != nil {
		return err
	}

	// prepare update
	update, err := updates.PrepareUpdates(updateProof, deltas, client.signingAccount.PrivateKey)
	if err != nil {
		return err
	}

	// submit update
	logger.Info("submitting update for block %d replicate %d: %s", updateProof.BlockNumber, updateProof.Replicate, deltasString)
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(config.CallTimeoutMillisDefault)*time.Millisecond)
	txOpts.Context = ctx

	var tx *types.Transaction
	// submit directly to Fast Updates contract if the submission contract is not available
	if client.submission == nil {
		tx, err = client.fastUpdater.SubmitUpdates(txOpts, *update)
		cancelFunc()
		if err != nil {
			return err
		}
	} else {
		strippedPackedCall, err := updates.PrepareUpdatesSubmission(update)
		if err != nil {
			return err
		}
		tx, err = client.submission.SubmitAndPass(txOpts, strippedPackedCall)
		cancelFunc()
		if err != nil {
			return err
		}
	}

	// wait for the transaction to be confirmed
	ctx, cancelFunc = context.WithTimeout(context.Background(), time.Duration(config.TxTimeoutMillisDefault)*time.Millisecond)
	receipt, err := bind.WaitMined(ctx, client.chainClient, tx)
	cancelFunc()
	if err != nil {
		return err
	}
	if receipt.Status == 0 {
		return fmt.Errorf("transaction failed")
	}
	logger.Info("successful update for block %d replicate %d in block %d", updateProof.BlockNumber, updateProof.Replicate, receipt.BlockNumber.Int64())
	client.Stats.NumSuccessfulUpdates++

	// get current prices from on-chain
	chainValues, _, err = client.GetFeeds(supportedFeedIndexes)
	if err != nil {
		return err
	}
	logger.Info("chain feeds values in block %d (after update): %v", receipt.BlockNumber.Int64(), chainValues)

	return nil
}

func (client *FastUpdatesClient) GetFastUpdaterContractAddress() (common.Address, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(config.CallTimeoutMillisDefault)*time.Millisecond)
	ops := &bind.CallOpts{Context: ctx}
	fastUpdaterContractAddress, err := client.submission.SubmitAndPassContract(ops)
	cancelFunc()
	if err != nil {
		return common.Address{}, err
	}

	return fastUpdaterContractAddress, nil
}
