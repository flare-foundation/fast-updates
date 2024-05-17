package client

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"

	"fast-updates-client/config"
	"fast-updates-client/contracts-interface/mock"
	"fast-updates-client/logger"
	"fast-updates-client/provider"
	"fast-updates-client/sortition"
	"fast-updates-client/updates"
)

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

func (client *FastUpdatesClient) GetPrices(feedIndexes []int) ([]float64, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(config.CallTimeoutMillisDefault)*time.Millisecond)
	ops := &bind.CallOpts{Context: ctx}

	feedIndexesBigInt := make([]*big.Int, len(feedIndexes))
	for i, index := range feedIndexes {
		feedIndexesBigInt[i] = big.NewInt(int64(index))
	}

	feedValues, err := client.fastUpdater.FetchCurrentFeeds(ops, feedIndexesBigInt)
	cancelFunc()

	floatValues := RawChainValuesToFloats(feedValues)
	return floatValues, err
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

func (client *FastUpdatesClient) getOnlineOfflineValues() ([]int, []float64, []float64,  error) {
	// 0 value indicates unsupported feed. TODO: need to differentiate between 0 and absent value better.
	providerRawValues, err := client.valuesProvider.GetValues(client.allFeeds)
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "error getting feed values")
	}

	supportedFeedIndexes := []int{}
	providerValues := []float64{}
	for i, value := range providerRawValues {
		if value != 0 {
			supportedFeedIndexes = append(supportedFeedIndexes, i)
			providerValues = append(providerValues, float64(value))
		}
	}

	// get current prices from on-chain
	chainValues, err := client.GetPrices(supportedFeedIndexes)
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

	logger.Info("chain feeds values in block %d (before update): %v", client.transactionQueue.CurrentBlockNum, chainValues)
	logger.Info("provider feeds values: %v", providerValues)

	// get current scale
	scale, err := client.GetScale()
	if err != nil {
		return err
	}

	// calculate deltas for provider and on-chain prices
	deltas, deltasString, err := provider.GetDeltas(chainValues, providerValues, supportedFeedIndexes, scale)
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

	// get current prices from on-chain
	chainValues, err = client.GetPrices(supportedFeedIndexes)
	if err != nil {
		return err
	}
	logger.Info("chain feeds values in block %d (after update): %v", receipt.BlockNumber.Int64(), chainValues)

	return nil
}
