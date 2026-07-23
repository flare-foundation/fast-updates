package e2e

import (
	"context"
	"testing"
)

// runBlocks is the number of blocks each scenario drives the client for. With
// an expected sample size of 1 update per block and a scale step of ~0.0122%,
// this comfortably covers targets ±0.2-0.3% away from the starting values.
const runBlocks = 60

// fingerprintTargets gives every feed a unique combination of direction and
// magnitude, strictly alternating up/down across neighbouring indices. Any
// updates applied to a wrong feed index then move some feed against its own
// target and fail verification, instead of hiding behind a neighbour with
// the same target.
func fingerprintTargets() map[int]float64 {
	return map[int]float64{
		0: startFeedValue * 1.0020, // ~16 steps up
		1: startFeedValue * 0.9985, // ~12 steps down
		2: startFeedValue * 1.0025, // ~21 steps up
		3: startFeedValue * 0.9980, // ~16 steps down
		4: startFeedValue * 1.0015, // ~12 steps up
		5: startFeedValue * 0.9975, // ~21 steps down
		6: startFeedValue * 1.0010, // ~8 steps up
		7: startFeedValue,          // already at target, must never move
		8: startFeedValue * 0.9990, // ~8 steps down
	}
}

// requireHealthySubmissions checks that the client's successful updates all
// landed on chain and nothing it sent reverted.
func requireHealthySubmissions(t *testing.T, env *TestEnv, startBlock, endBlock, reported uint64) {
	successful, reverted := env.CountSubmissions(startBlock, endBlock)
	t.Logf("submissions in blocks %d..%d: %d successful, %d reverted", startBlock, endBlock, successful, reverted)
	if reverted != 0 {
		t.Errorf("%d submissions reverted; a healthy client mines everything it sends", reverted)
	}
	if successful != int(reported) {
		t.Errorf("client reported %d successful updates but %d submitUpdates transactions were mined",
			reported, successful)
	}
}

// TestScenarioConvergence is the happy path: every feed has a provider value
// a few dozen scale steps away from the chain value (up, down, or exactly at
// it). Feeds may only ever move toward their targets and must all end within
// a few scale steps of them.
func TestScenarioConvergence(t *testing.T) {
	env := NewTestEnv(t)

	targets := fingerprintTargets()
	for feed, target := range targets {
		env.Provider.SetTarget(feed, target)
	}

	startBlock, endBlock, stats := env.RunClient(runBlocks)
	t.Logf("client stats: %d scheduled, %d successful", stats.NumUpdates, stats.NumSuccessfulUpdates)
	if stats.NumSuccessfulUpdates < 30 {
		t.Fatalf("expected at least 30 successful updates in %d blocks, got %d",
			runBlocks, stats.NumSuccessfulUpdates)
	}
	requireHealthySubmissions(t, env, startBlock, endBlock, stats.NumSuccessfulUpdates)

	env.VerifyFeedMovements(startBlock, endBlock, Expectations{
		Targets:     targets,
		FrozenFeeds: []int{7},
	})
	env.AssertConverged(endBlock, targets, 4)
}

// TestScenarioMissingFeeds exercises the index alignment when some feeds have
// no values: feeds 2 and 8 are not supported by the provider (it returns nil
// for them) and feed 5 is delisted on-chain (its feed id zeroed, index kept).
// None of them may ever move, and the remaining feeds must still converge to
// their own targets.
func TestScenarioMissingFeeds(t *testing.T) {
	env := NewTestEnv(t)

	env.RemoveFeed(5)

	targets := fingerprintTargets()
	delete(targets, 2)
	delete(targets, 5)
	delete(targets, 8)
	for feed, target := range targets {
		env.Provider.SetTarget(feed, target)
	}

	startBlock, endBlock, stats := env.RunClient(runBlocks)
	t.Logf("client stats: %d scheduled, %d successful", stats.NumUpdates, stats.NumSuccessfulUpdates)
	if stats.NumSuccessfulUpdates < 30 {
		t.Fatalf("expected at least 30 successful updates in %d blocks, got %d",
			runBlocks, stats.NumSuccessfulUpdates)
	}
	requireHealthySubmissions(t, env, startBlock, endBlock, stats.NumSuccessfulUpdates)

	env.VerifyFeedMovements(startBlock, endBlock, Expectations{
		Targets:     targets,
		FrozenFeeds: []int{2, 5, 7, 8},
	})
	env.AssertConverged(endBlock, targets, 4)
}

// TestScenarioFeedReplacedMidRun removes a feed and adds a feed with a new id
// while the client is running. The contract reuses the zeroed index for the
// new feed (FastUpdatesConfiguration keeps removed indices in unusedIndices),
// while the client keeps working on its stale feed list until its periodic
// refresh (every 100 blocks). The new feed's target points the opposite way
// from the old feed's, so the slot must move down once the client picks up
// the new feed id — proving updates follow the feed id, not a stale index.
func TestScenarioFeedReplacedMidRun(t *testing.T) {
	env := NewTestEnv(t)

	const (
		replacedSlot  = 2
		newFeedIdByte = 10 // provider key DeployedFeedId(9)
	)
	targets := fingerprintTargets() // old slot-2 feed target: +0.25%
	for feed, target := range targets {
		env.Provider.SetTarget(feed, target)
	}
	// the replacement feed starts at 1000 (mock publisher default) and wants
	// to go down, opposite to the old feed at the same slot
	newFeedTarget := startFeedValue * 0.9990
	env.Provider.SetTarget(newFeedIdByte-1, newFeedTarget)

	head, err := env.EthClient.BlockNumber(context.Background())
	if err != nil {
		t.Fatalf("getting head block: %v", err)
	}
	changeAt := head + 25
	// first feed-list refresh (client refreshes every 100 blocks) after the
	// feed replacement
	refreshBlock := (changeAt/100 + 1) * 100
	endTarget := refreshBlock + 35

	startBlock, wait := env.RunClientAsync(endTarget - head)

	env.WaitForChainBlock(changeAt)
	removeBlock := env.RemoveFeed(replacedSlot)
	addBlock := env.AddFeed(newFeedIdByte)
	if got := env.FeedIdAt(replacedSlot); got != newFeedIdByte {
		t.Fatalf("expected new feed id %d at index %d, got %d — did addFeeds stop reusing removed indices?",
			newFeedIdByte, replacedSlot, got)
	}
	t.Logf("feed %d removed in block %d, replacement added in block %d, client refreshes at block %d",
		replacedSlot, removeBlock, addBlock, refreshBlock)

	endBlock, stats := wait()
	t.Logf("client stats: %d scheduled, %d successful", stats.NumUpdates, stats.NumSuccessfulUpdates)
	if stats.NumSuccessfulUpdates < 60 {
		t.Fatalf("expected at least 60 successful updates in %d blocks, got %d",
			endTarget-head, stats.NumSuccessfulUpdates)
	}
	requireHealthySubmissions(t, env, startBlock, endBlock, stats.NumSuccessfulUpdates)

	// before the replacement: business as usual
	env.VerifyFeedMovements(startBlock, removeBlock-1, Expectations{
		Targets:     targets,
		FrozenFeeds: []int{7},
	})

	// between the replacement and shortly after the feed refresh the replaced
	// slot's direction is ambiguous: the stale client pushes it toward the
	// OLD feed's target until the refresh (plus in-flight submissions up to a
	// submission window later), then toward the NEW one — so the slot has no
	// target here and only the other feeds are checked
	staleTargets := make(map[int]float64, len(targets))
	for feed, target := range targets {
		staleTargets[feed] = target
	}
	delete(staleTargets, replacedSlot)
	staleWindowEnd := refreshBlock + uint64(env.Cfg.Client.SubmissionWindow) + 1
	env.VerifyFeedMovements(addBlock+1, staleWindowEnd, Expectations{
		Targets:     staleTargets,
		FrozenFeeds: []int{7},
	})

	// after the refresh the client must track the new feed id: the replaced
	// slot now moves toward the new feed's (lower) target
	postTargets := make(map[int]float64, len(targets))
	for feed, target := range targets {
		postTargets[feed] = target
	}
	postTargets[replacedSlot] = newFeedTarget
	env.VerifyFeedMovements(staleWindowEnd+1, endBlock, Expectations{
		Targets:     postTargets,
		FrozenFeeds: []int{7},
	})

	// the new feed must actually be moving toward its own target
	atRefresh := env.FeedStateAt(staleWindowEnd).Values()[replacedSlot]
	atEnd := env.FeedStateAt(endBlock).Values()[replacedSlot]
	if atEnd >= atRefresh {
		t.Errorf("replaced slot %d did not move toward the new feed target after refresh: %.6f -> %.6f (target %.6f)",
			replacedSlot, atRefresh, atEnd, newFeedTarget)
	}

	// all other feeds converged to their targets long before the end
	delete(postTargets, replacedSlot)
	env.AssertConverged(endBlock, postTargets, 4)
}
