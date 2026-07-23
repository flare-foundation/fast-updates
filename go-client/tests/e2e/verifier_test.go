package e2e

import (
	"bytes"
	"context"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"fast-updates-client/contracts-interface/fast_updater"
	"fast-updates-client/contracts-interface/incentive"
)

// FeedState is the chain state of all feeds at one block height, read via
// fetchCurrentFeeds (the canonical read path, applying any pending deltas).
type FeedState struct {
	Raw      []*big.Int
	Decimals []int8
}

// Values converts the raw values to floats independently of the client's
// conversion code. Raw values fit in uint32, so float64 arithmetic is exact.
func (s *FeedState) Values() []float64 {
	values := make([]float64, len(s.Raw))
	for i, raw := range s.Raw {
		values[i] = float64(raw.Int64()) * math.Pow(10, -float64(s.Decimals[i]))
	}
	return values
}

// FeedStateAt reads all feeds at the given block height (0 = latest) through
// fetchCurrentFeeds, hand-rolled as an eth_call because the method is payable.
func (env *TestEnv) FeedStateAt(blockNum uint64) *FeedState {
	fuABI, err := fast_updater.FastUpdaterMetaData.GetAbi()
	if err != nil {
		env.t.Fatalf("getting FastUpdater ABI: %v", err)
	}

	indices := make([]*big.Int, numFeeds)
	for i := range indices {
		indices[i] = big.NewInt(int64(i))
	}
	input, err := fuABI.Pack("fetchCurrentFeeds", indices)
	if err != nil {
		env.t.Fatalf("packing fetchCurrentFeeds: %v", err)
	}

	value, _ := new(big.Int).SetString(env.Cfg.Client.FetchCurrentFeedsValue, 10)
	msg := ethereum.CallMsg{
		From:  common.HexToAddress(env.Cfg.Client.FetchCurrentFeedsAddress),
		To:    &env.Contracts.FastUpdater,
		Data:  input,
		Value: value,
	}
	var atBlock *big.Int
	if blockNum != 0 {
		atBlock = new(big.Int).SetUint64(blockNum)
	}
	output, err := env.EthClient.CallContract(context.Background(), msg, atBlock)
	if err != nil {
		env.t.Fatalf("fetchCurrentFeeds at block %d: %v", blockNum, err)
	}
	unpacked, err := fuABI.Unpack("fetchCurrentFeeds", output)
	if err != nil {
		env.t.Fatalf("unpacking fetchCurrentFeeds result: %v", err)
	}

	return &FeedState{
		Raw:      *abi.ConvertType(unpacked[0], new([]*big.Int)).(*[]*big.Int),
		Decimals: *abi.ConvertType(unpacked[1], new([]int8)).(*[]int8),
	}
}

// ScaleFloat reads the incentive manager's scale and returns it as a plain
// multiplier (e.g. 1.000122): the factor a single update applies to a feed.
func (env *TestEnv) ScaleFloat() float64 {
	incentiveContract, err := incentive.NewIncentive(env.Contracts.IncentiveManager, env.EthClient)
	if err != nil {
		env.t.Fatalf("binding incentive manager: %v", err)
	}
	scale, err := incentiveContract.GetScale(&bind.CallOpts{})
	if err != nil {
		env.t.Fatalf("reading scale: %v", err)
	}
	scaleFloat, _ := scale.Float64()
	return scaleFloat / math.Pow(2, 127)
}

// CountSubmissions scans the block range for submitUpdates transactions to
// the FastUpdater contract and returns how many were mined successfully and
// how many reverted. On a healthy run every submission the client sends is
// mined successfully, so scenarios expect zero reverts.
func (env *TestEnv) CountSubmissions(startBlock, endBlock uint64) (successful, reverted int) {
	fuABI, err := fast_updater.FastUpdaterMetaData.GetAbi()
	if err != nil {
		env.t.Fatalf("getting FastUpdater ABI: %v", err)
	}
	selector := fuABI.Methods["submitUpdates"].ID

	for blockNum := startBlock; blockNum <= endBlock; blockNum++ {
		block, err := env.EthClient.BlockByNumber(context.Background(), new(big.Int).SetUint64(blockNum))
		if err != nil {
			env.t.Fatalf("getting block %d: %v", blockNum, err)
		}
		for _, tx := range block.Transactions() {
			if tx.To() == nil || *tx.To() != env.Contracts.FastUpdater ||
				len(tx.Data()) < 4 || !bytes.Equal(tx.Data()[:4], selector) {
				continue
			}
			receipt, err := env.EthClient.TransactionReceipt(context.Background(), tx.Hash())
			if err != nil {
				env.t.Fatalf("getting receipt for %s: %v", tx.Hash(), err)
			}
			if receipt.Status == types.ReceiptStatusSuccessful {
				successful++
			} else {
				reverted++
			}
		}
	}
	return successful, reverted
}

// Expectations describes how the feeds may move during a verified range.
type Expectations struct {
	// Targets holds the provider value per feed index; a feed may only move
	// toward its target.
	Targets map[int]float64
	// FrozenFeeds are feed indices whose raw on-chain value must stay
	// byte-identical over the whole range: feeds without a provider value,
	// delisted feeds, and feeds already exactly at their target.
	FrozenFeeds []int
}

// VerifyFeedMovements replays the chain block by block and checks that every
// feed with a target only ever moves toward it and that frozen feeds never
// move at all. Near the target the client's volatility band allows a little
// jitter (including overshoot by a few scale steps), so movements away from
// the target are only errors while the feed is outside that zone.
func (env *TestEnv) VerifyFeedMovements(startBlock, endBlock uint64, expect Expectations) {
	t := env.t
	scaleStep := env.ScaleFloat() - 1

	frozen := make(map[int]bool, len(expect.FrozenFeeds))
	for _, feed := range expect.FrozenFeeds {
		frozen[feed] = true
	}

	prev := env.FeedStateAt(startBlock)
	prevValues := prev.Values()
	for blockNum := startBlock + 1; blockNum <= endBlock; blockNum++ {
		cur := env.FeedStateAt(blockNum)
		curValues := cur.Values()

		for feed := 0; feed < numFeeds; feed++ {
			if frozen[feed] {
				if prev.Raw[feed].Cmp(cur.Raw[feed]) != 0 || prev.Decimals[feed] != cur.Decimals[feed] {
					t.Errorf("block %d: frozen feed %d moved (raw %s -> %s)",
						blockNum, feed, prev.Raw[feed], cur.Raw[feed])
				}
				continue
			}
			target, hasTarget := expect.Targets[feed]
			if !hasTarget {
				continue
			}

			prevDist := math.Abs(prevValues[feed] - target)
			curDist := math.Abs(curValues[feed] - target)
			// jitterZone: within a few scale steps of the target the band
			// logic may randomly push either way; epsilon covers fixed-point
			// rounding of buffered deltas (a raw unit or two per block)
			jitterZone := 4 * scaleStep * target
			epsilon := 0.25 * scaleStep * target
			if curDist > prevDist+epsilon && curDist > jitterZone {
				t.Errorf("block %d: feed %d moved away from target %.6f (%.6f -> %.6f)",
					blockNum, feed, target, prevValues[feed], curValues[feed])
			}
		}
		prev, prevValues = cur, curValues
	}
}

// AssertConverged checks that each feed with a provider target ended the run
// within the given number of scale steps of that target.
func (env *TestEnv) AssertConverged(endBlock uint64, targets map[int]float64, withinSteps float64) {
	values := env.FeedStateAt(endBlock).Values()
	scaleStep := env.ScaleFloat() - 1
	for feed, target := range targets {
		if math.Abs(values[feed]-target) > withinSteps*scaleStep*target {
			env.t.Errorf("feed %d did not converge: chain %.8f, target %.8f (allowed %.1f scale steps of %.6f%%)",
				feed, values[feed], target, withinSteps, scaleStep*100)
		}
	}
}
