package provider_test

import (
	"fast-updates-client/provider"
	"testing"
)

func TestGetRandomDeltas(t *testing.T) {
	randProvider := provider.NewRandomFeedsProvider()
	feedsIds := []provider.FeedId{
		{Category: 1, Name: "BTC/USD"},
		{Category: 1, Name: "ETH/USD"},
		{Category: 1, Name: "BOL/USD"},
	}
	feeds, err := randProvider.GetValues(feedsIds)
	if err != nil {
		t.Fatal(err)
	}
	if len(feeds) != len(feedsIds) {
		t.Fatalf("number of deltas in the representation string not correct")
	}
}
