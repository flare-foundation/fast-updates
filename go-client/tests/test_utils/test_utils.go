package test_utils

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"math"
	"strconv"

	"fast-updates-client/client"
	"fast-updates-client/config"
	"fast-updates-client/contracts-interface/fast_updater"
	"fast-updates-client/contracts-interface/fast_updates_configuration"
	"fast-updates-client/contracts-interface/fee_calculator"
	"fast-updates-client/contracts-interface/incentive"
	"fast-updates-client/contracts-interface/mock"
	"fast-updates-client/logger"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	BASE_SAMPLE_SIZE, _      = new(big.Int).SetString("0x01000000000000000000000000000000", 0)
	BASE_RANGE, _            = new(big.Int).SetString("0x00000800000000000000000000000000", 0)
	SAMPLE_INCREASE_LIMIT, _ = new(big.Int).SetString("0x00100000000000000000000000000000", 0)
	RANGE_INCREASE_LIMIT, _  = new(big.Int).SetString("0x00008000000000000000000000000000", 0)
	SAMPLE_INCREASE_PRICE    = big.NewInt(5)
	RANGE_INCREASE_PRICE     = new(big.Int).Exp(big.NewInt(10), big.NewInt(24), nil)
	DURATION                 = big.NewInt(8)
	EPOCH_LEN                = big.NewInt(1000)
	// starting feeds hardcoded in mocked contract to be 100000
	FEEDS_INDICES     = []*big.Int{big.NewInt(0), big.NewInt(1), big.NewInt(2), big.NewInt(3), big.NewInt(4), big.NewInt(5), big.NewInt(6), big.NewInt(7), big.NewInt(8)}
	SUBMISSION_WINDOW = uint8(5)
	BACKLOG_LEN       = big.NewInt(20)
)

type ContractAddresses struct {
	FastUpdater              common.Address
	FastUpdatesConfiguration common.Address
	IncentiveManager         common.Address
	Mock                     common.Address
}

func Register(cfg *config.Config, numEpochs int) error {
	client, err := client.CreateFastUpdatesClient(cfg, nil)
	if err != nil {
		return err
	}

	epoch, err := client.GetCurrentRewardEpochId()
	if err != nil {
		return err
	}

	for i := 0; i < numEpochs; i++ {
		client.Register(epoch + int64(i))
	}

	client.WaitToEmptyRequests()
	client.Stop()

	return nil
}

func Deploy(cfg *config.Config) ContractAddresses {
	client, err := ethclient.Dial(cfg.Chain.NodeURL)
	if err != nil {
		logger.Fatal("Error: %s", err)
	}
	privateKey := cfg.Transactions.Accounts[0]
	if privateKey[:2] == "0x" {
		privateKey = privateKey[2:]
	}
	privateKeyECDSA, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		logger.Fatal("Error: %s", err)
	}

	publicKey := privateKeyECDSA.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		logger.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		logger.Fatal("Error: %s", err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		logger.Fatal("Error: %s", err)
	}

	opts, err := bind.NewKeyedTransactorWithChainID(privateKeyECDSA, big.NewInt(int64(cfg.Chain.ChainId)))
	if err != nil {
		logger.Fatal("Error: %s", err)
	}
	opts.Nonce = big.NewInt(int64(nonce))
	opts.Value = big.NewInt(int64(cfg.Transactions.Value))
	opts.GasLimit = uint64(cfg.Transactions.GasLimit)
	opts.GasPrice = gasPrice

	mockAddress, _, _, err := mock.DeployMock(opts, client, big.NewInt(1), EPOCH_LEN)
	if err != nil {
		logger.Fatal("Error: %s", err)
	}
	logger.Info("mock contract address %s", mockAddress.Hex())
	opts.Nonce.Add(opts.Nonce, big.NewInt(1))

	incentiveAddress, _, _, err := incentive.DeployIncentive(
		opts, client, fromAddress, fromAddress, fromAddress, BASE_SAMPLE_SIZE, BASE_RANGE,
		SAMPLE_INCREASE_LIMIT, RANGE_INCREASE_LIMIT, SAMPLE_INCREASE_PRICE, RANGE_INCREASE_PRICE, DURATION)
	if err != nil {
		logger.Fatal("Error: %s", err)
	}
	logger.Info("incentiveManager address %s", incentiveAddress.Hex())

	opts.Nonce.Add(opts.Nonce, big.NewInt(1))

	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		logger.Fatal("Error: %s", err)
	}

	fastUpdaterAddress, tx, _, err := fast_updater.DeployFastUpdater(opts, client, fromAddress, fromAddress,
		fromAddress, fromAddress, uint32(block.Time()), 90, SUBMISSION_WINDOW)
	if err != nil {
		logger.Fatal("Error: %s", err)
	}
	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		logger.Fatal("Error: %s", err)
	}
	if receipt.Status == 0 {
		reason, err := GetFailingMessage(*client, tx.Hash())
		if err != nil {
			logger.Fatal("Error: %s", err)
		}
		logger.Fatal("Error: Transaction fail: %s", reason)
	}
	logger.Info("fastUpdater address %s", fastUpdaterAddress.Hex())
	opts.Nonce.Add(opts.Nonce, big.NewInt(1))

	fastUpdatesConfigurationAddress, tx, _, err := fast_updates_configuration.DeployFastUpdatesConfiguration(opts, client, fromAddress, fromAddress, fromAddress)
	if err != nil {
		logger.Fatal("Error: %s", err)
	}
	logger.Info("fastUpdatesConfiguration address %s", fastUpdatesConfigurationAddress.Hex())
	receipt, err = bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		logger.Fatal("Error: %s", err)
	}
	if receipt.Status == 0 {
		reason, err := GetFailingMessage(*client, tx.Hash())
		if err != nil {
			logger.Fatal("Error: %s", err)
		}
		logger.Fatal("Error: Transaction fail: %s", reason)
	}
	opts.Nonce.Add(opts.Nonce, big.NewInt(1))

	feeCalculatorAddress, tx, _, err := fee_calculator.DeployFeeCalculator(opts, client, fromAddress, fromAddress, fromAddress, big.NewInt(1))
	if err != nil {
		logger.Fatal("Error: %s", err)
	}
	logger.Info("feeCalculator address %s", feeCalculatorAddress.Hex())
	receipt, err = bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		logger.Fatal("Error: %s", err)
	}
	if receipt.Status == 0 {
		reason, err := GetFailingMessage(*client, tx.Hash())
		if err != nil {
			logger.Fatal("Error: %s", err)
		}
		logger.Fatal("Error: Transaction fail: %s", reason)
	}
	opts.Nonce.Add(opts.Nonce, big.NewInt(1))

	feeCalculatorContract, err := fee_calculator.NewFeeCalculator(feeCalculatorAddress, client)
	if err != nil {
		logger.Fatal("Error: %s", err)
	}

	addressesHash := []string{
		"0x12e7f85251b6a8cc2a2841f61f59a88110842aebcb7b0156dd0c10bd473fcb7a",
		"0x6be6257da65c607a560a35b4efea3c17b461c71f51e72de30b7c1e124e6b8153",
	}
	addressesBytes := make([][32]byte, len(addressesHash))
	for i := 0; i < len(addressesHash); i++ {
		var buf [32]byte
		b, err := hex.DecodeString(addressesHash[i][2:])
		if err != nil {
			log.Fatal(err)
		}
		copy(buf[:], b)
		addressesBytes[i] = buf
	}
	addresses := []common.Address{
		fromAddress,
		fastUpdatesConfigurationAddress,
	}
	tx, err = feeCalculatorContract.UpdateContractAddresses(opts, addressesBytes, addresses)
	if err != nil {
		logger.Fatal("Error: %s", err)
	}
	receipt, err = bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		logger.Fatal("Error: %s", err)
	}
	if receipt.Status == 0 {
		reason, err := GetFailingMessage(*client, tx.Hash())
		if err != nil {
			logger.Fatal("Error: %s", err)
		}
		logger.Fatal("Error: Transaction fail: %s", reason)
	}
	opts.Nonce.Add(opts.Nonce, big.NewInt(1))

	fastUpdaterContract, err := fast_updater.NewFastUpdater(fastUpdaterAddress, client)
	if err != nil {
		logger.Fatal("Error: %s", err)
	}

	addressesHash = []string{
		"0x12e7f85251b6a8cc2a2841f61f59a88110842aebcb7b0156dd0c10bd473fcb7a",
		"0x2b5425460b937e96e509004540fff99ad6ec17948dba96effce0ba122b8bb899",
		"0x7ae386e71020f3892e238530238dee40111e0bff57a096544e6b6806e26e8ab0",
		"0x7de5495162bf7c2e65e3e8356a8981e85633d651c850dcb5b6e0c0b8a878a195",
		"0x6be6257da65c607a560a35b4efea3c17b461c71f51e72de30b7c1e124e6b8153",
		"0x597295c852f29045b82e8864e15b8a3e2c0da8de0e4fbdd3ec498197e11d6a5e",
		"0x94e7b7fc9128e3dd39886c5e5a0ad2700a92e886b7bb5f98d3dd4f5dddb2272e",
	}

	addressesBytes = make([][32]byte, len(addressesHash))
	for i := 0; i < len(addressesHash); i++ {
		var buf [32]byte
		b, err := hex.DecodeString(addressesHash[i][2:])
		if err != nil {
			log.Fatal(err)
		}
		copy(buf[:], b)
		addressesBytes[i] = buf
	}
	addresses = []common.Address{
		fromAddress,
		mockAddress,
		incentiveAddress,
		mockAddress,
		fastUpdatesConfigurationAddress,
		mockAddress,
		feeCalculatorAddress,
	}
	tx, err = fastUpdaterContract.UpdateContractAddresses(opts, addressesBytes, addresses)
	if err != nil {
		logger.Fatal("Error: %s", err)
	}
	receipt, err = bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		logger.Fatal("Error: %s", err)
	}
	if receipt.Status == 0 {
		reason, err := GetFailingMessage(*client, tx.Hash())
		if err != nil {
			logger.Fatal("Error: %s", err)
		}
		logger.Fatal("Error: Transaction fail: %s", reason)
	}
	opts.Nonce.Add(opts.Nonce, big.NewInt(1))

	fastUpdatesConfigurationContract, err := fast_updates_configuration.NewFastUpdatesConfiguration(fastUpdatesConfigurationAddress, client)
	if err != nil {
		logger.Fatal("Error: %s", err)
	}

	addressesHash = []string{
		"0x12e7f85251b6a8cc2a2841f61f59a88110842aebcb7b0156dd0c10bd473fcb7a",
		"0x0cf0bcabf35e9f54dc06269101d6c97535ba08da6ca99a9c5df65a4dd717919c",
	}
	addressesBytes = make([][32]byte, len(addressesHash))
	for i := 0; i < len(addressesHash); i++ {
		var buf [32]byte
		b, err := hex.DecodeString(addressesHash[i][2:])
		if err != nil {
			log.Fatal(err)
		}
		copy(buf[:], b)
		addressesBytes[i] = buf
	}

	addresses = []common.Address{
		fromAddress,
		fastUpdaterAddress,
	}

	tx, err = fastUpdatesConfigurationContract.UpdateContractAddresses(opts, addressesBytes, addresses)
	if err != nil {
		logger.Fatal("Error: %s", err)
	}
	receipt, err = bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		logger.Fatal("Error: %s", err)
	}
	if receipt.Status == 0 {
		reason, err := GetFailingMessage(*client, tx.Hash())
		if err != nil {
			logger.Fatal("Error: %s", err)
		}
		logger.Fatal("Error: Transaction fail: %s", reason)
	}
	opts.Nonce.Add(opts.Nonce, big.NewInt(1))

	incentiveContract, err := incentive.NewIncentive(incentiveAddress, client)
	if err != nil {
		logger.Fatal("Error: %s", err)
	}

	addressesHash = []string{
		"0x12e7f85251b6a8cc2a2841f61f59a88110842aebcb7b0156dd0c10bd473fcb7a",
		"0x2b5425460b937e96e509004540fff99ad6ec17948dba96effce0ba122b8bb899",
		"0x3ea59489aed5d8b20b9ba382cadf6c13290f9ee8edecbebabb7712ee6ad2b232",
		"0x23c7c13c4a1e2bdc03481c359a1b7a7bbd0b2f5bd53eab76cf21b9e145f735eb",
		"0x0cf0bcabf35e9f54dc06269101d6c97535ba08da6ca99a9c5df65a4dd717919c",
		"0x6be6257da65c607a560a35b4efea3c17b461c71f51e72de30b7c1e124e6b8153",
	}
	addressesBytes = make([][32]byte, len(addressesHash))
	for i := 0; i < len(addressesHash); i++ {
		var buf [32]byte
		b, err := hex.DecodeString(addressesHash[i][2:])
		if err != nil {
			log.Fatal(err)
		}
		copy(buf[:], b)
		addressesBytes[i] = buf
	}

	addresses = []common.Address{
		fromAddress,
		fromAddress,
		fromAddress,
		fromAddress,
		fastUpdaterAddress,
		fromAddress,
	}

	tx, err = incentiveContract.UpdateContractAddresses(opts, addressesBytes, addresses)
	if err != nil {
		logger.Fatal("Error: %s", err)
	}
	receipt, err = bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		logger.Fatal("Error: %s", err)
	}
	if receipt.Status == 0 {
		reason, err := GetFailingMessage(*client, tx.Hash())
		if err != nil {
			logger.Fatal("Error: %s", err)
		}
		logger.Fatal("Error: Transaction fail: %s", reason)
	}
	opts.Nonce.Add(opts.Nonce, big.NewInt(1))

	feedsConfigurations := make([]fast_updates_configuration.IFastUpdatesConfigurationFeedConfiguration, len(FEEDS_INDICES))
	for i := 0; i < len(FEEDS_INDICES); i++ {
		feedId := [21]byte{}
		feedId[20] = byte(i + 1)
		feedsConfigurations[i] = fast_updates_configuration.IFastUpdatesConfigurationFeedConfiguration{FeedId: feedId, RewardBandValue: 2000, InflationShare: big.NewInt(200)}
	}

	tx, err = fastUpdatesConfigurationContract.AddFeeds(opts, feedsConfigurations)
	if err != nil {
		logger.Fatal("Error: %s", err)
	}
	receipt, err = bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		logger.Fatal("Error: %s", err)
	}
	if receipt.Status == 0 {
		reason, err := GetFailingMessage(*client, tx.Hash())
		if err != nil {
			logger.Fatal("Error: %s", err)
		}
		logger.Fatal("Error: Transaction fail: %s", reason)
	}

	opts.Nonce.Add(opts.Nonce, big.NewInt(1))

	numFeedsCheck, err := fastUpdatesConfigurationContract.GetNumberOfFeeds(nil)
	if err != nil {
		logger.Fatal("Error: %s", err)
	}

	if int(numFeedsCheck.Int64()) != len(FEEDS_INDICES) {
		logger.Fatal("Error: Wrong number of feeds: %d != %d", numFeedsCheck.Int64(), len(FEEDS_INDICES))
	}

	logger.Info("Contracts deployed, wait " + strconv.Itoa(int(SUBMISSION_WINDOW)) + " blocks for the daemon to set thresholds")
	for i := 0; i < int(SUBMISSION_WINDOW+1); i++ {
		tx, err = fastUpdaterContract.Daemonize(opts)
		if err != nil {
			logger.Fatal("Error: %s", err)
		}
		rec, err := bind.WaitMined(context.Background(), client, tx)
		if err != nil {
			logger.Fatal("Error: %s", err)
		}
		logger.Info("block mined %s", rec.BlockNumber.String())
		// reason, err := GetFailingMessage(*client, tx.Hash())

		opts.Nonce.Add(opts.Nonce, big.NewInt(1))
	}

	contracts := ContractAddresses{
		FastUpdater:              fastUpdaterAddress,
		FastUpdatesConfiguration: fastUpdatesConfigurationAddress,
		IncentiveManager:         incentiveAddress,
		Mock:                     mockAddress,
	}

	return contracts
}

func GetFailingMessage(client ethclient.Client, hash common.Hash) (string, error) {
	tx, _, err := client.TransactionByHash(context.Background(), hash)
	if err != nil {
		return "", err
	}

	from, err := types.Sender(types.NewEIP155Signer(tx.ChainId()), tx)
	if err != nil {
		return "", err
	}

	msg := ethereum.CallMsg{
		From:     from,
		To:       tx.To(),
		Gas:      tx.Gas(),
		GasPrice: tx.GasPrice(),
		Value:    tx.Value(),
		Data:     tx.Data(),
	}

	res, err := client.CallContract(context.Background(), msg, nil)
	if err != nil {
		return "", err
	}

	return string(res), nil
}

func Incentivize(cfg *config.Config, rangeIncrease float64, sampleCost *big.Int) {
	client, err := ethclient.Dial(cfg.Chain.NodeURL)
	if err != nil {
		logger.Fatal("Error: %s", err)
	}
	privateKey := cfg.Transactions.Accounts[0]
	if privateKey[:2] == "0x" {
		privateKey = privateKey[2:]
	}
	privateKeyECDSA, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		logger.Fatal("Error: %s", err)
	}

	publicKey := privateKeyECDSA.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		logger.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		logger.Fatal("Error: %s", err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		logger.Fatal("Error: %s", err)
	}

	opts, err := bind.NewKeyedTransactorWithChainID(privateKeyECDSA, big.NewInt(int64(cfg.Chain.ChainId)))
	if err != nil {
		logger.Fatal("Error: %s", err)
	}
	opts.Nonce = big.NewInt(int64(nonce))
	opts.Value = big.NewInt(int64(cfg.Transactions.Value))
	opts.GasLimit = uint64(cfg.Transactions.GasLimit)
	opts.GasPrice = gasPrice

	incentiveInterface, err := incentive.NewIncentive(common.HexToAddress(cfg.Client.IncentiveManagerAddress), client)
	if err != nil {
		logger.Fatal("Error: %s", err)
	}

	rangeInt, err := incentiveInterface.GetRange(nil)
	if err != nil {
		logger.Fatal("Error: %s", err)
	}
	rangeFloat, _ := rangeInt.Float64()
	rangeFloat = rangeFloat / math.Pow(2, 120)
	increase := rangeFloat * rangeIncrease
	increaseBigFloat := new(big.Float).SetFloat64(increase)
	logger.Info("current range %f, increasing for %f", rangeFloat, increase)
	increaseBigFloat.Mul(increaseBigFloat, big.NewFloat(math.Pow(2, 120)))
	increaseBigInt, _ := increaseBigFloat.Int(nil)

	rangeIncreasePrice, err := incentiveInterface.RangeIncreasePrice(nil)
	if err != nil {
		logger.Fatal("Error: %s", err)
	}
	price := new(big.Int).Mul(rangeIncreasePrice, increaseBigInt)
	price.Rsh(price, 120)
	priceFloat, _ := price.Float64()

	logger.Info("price for range increase %f wei", priceFloat)
	maxRange := new(big.Int).Exp(big.NewInt(2), big.NewInt(127), nil)

	sampleSizeInt, err := incentiveInterface.GetExpectedSampleSize(nil)
	if err != nil {
		logger.Fatal("Error: %s", err)
	}
	sampleFloat, _ := sampleSizeInt.Float64()
	sampleFloat = sampleFloat / math.Pow(2, 120)
	sampleCostFloat, _ := sampleCost.Float64()
	logger.Info("current sample size %f, increasing for price of %f wei", sampleFloat, sampleCostFloat)

	offer := incentive.IFastUpdateIncentiveManagerIncentiveOffer{RangeIncrease: increaseBigInt, RangeLimit: maxRange}
	opts.Value = new(big.Int).Add(price, sampleCost)
	tx, err := incentiveInterface.OfferIncentive(opts, offer)
	if err != nil {
		logger.Fatal("Error: %s", err)
	}
	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		logger.Fatal("Error: %s", err)
	}
	if receipt.Status == 0 {
		reason, err := GetFailingMessage(*client, tx.Hash())
		if err != nil {
			logger.Fatal("Error: %s", err)
		}
		logger.Fatal("Error: Transaction fail: %s", reason)
	}
	opts.Nonce.Add(opts.Nonce, big.NewInt(1))

	rangeInt, err = incentiveInterface.GetRange(nil)
	if err != nil {
		logger.Fatal("Error: %s", err)
	}
	rangeFloat, _ = rangeInt.Float64()
	rangeFloat = rangeFloat / math.Pow(2, 120)
	sampleSizeInt, err = incentiveInterface.GetExpectedSampleSize(nil)
	if err != nil {
		logger.Fatal("Error: %s", err)
	}
	sampleFloat, _ = sampleSizeInt.Float64()
	sampleFloat = sampleFloat / math.Pow(2, 120)
	logger.Info("range after transaction %f, sample size after transaction %f", rangeFloat, sampleFloat)

}
