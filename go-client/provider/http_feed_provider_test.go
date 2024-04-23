package provider_test

import (
	"fast-updates-client/client"
	"fast-updates-client/provider"
	"math/big"
	"testing"
)

func TestRawChainFeedsToFloats(t *testing.T) {
	rawFeeds := provider.ValuesDecimals{
		Feeds:    []*big.Int{big.NewInt(100), big.NewInt(200), big.NewInt(300)},
		Decimals: []int8{2, 2, 2},
	}

	expected := []float64{1.0, 2.0, 3.0}
	result := client.RawChainValuesToFloats(rawFeeds)

	if len(result) != len(expected) {
		t.Errorf("Expected length of result to be %d, but got %d", len(expected), len(result))
	}

	for i := range expected {
		if result[i] != expected[i] {
			t.Errorf("Expected result[%d] to be %f, but got %f", i, expected[i], result[i])
		}
	}
}

func TestGetProviderFeeds(t *testing.T) {
	// Add test case
	feedsIds := []provider.FeedId{
		{Category: 1, Name: "BTC/USD"},
		{Category: 1, Name: "ETH/USD"},
		{Category: 1, Name: "BOL/USD"},
	}
	// TODO: Mock endpoint and fix test
	feedProvider := provider.NewHttpValueProvider("http://localhost:3101/")

	feeds, err := feedProvider.GetValues(feedsIds)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if len(feeds) != len(feedsIds) {
		t.Errorf("Expected length of feeds to be %d, but got %d", len(feedsIds), len(feeds))
	}
}
