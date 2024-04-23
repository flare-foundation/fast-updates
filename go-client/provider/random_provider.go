package provider

import (
	"math/rand"
)

type RandomFeedsProvider struct {
}

func NewRandomFeedsProvider() RandomFeedsProvider {
	return RandomFeedsProvider{}
}

func (provider *RandomFeedsProvider) GetValues(feeds []FeedId) ([]float64, error) {
	randomFeeds := make([]float64, len(feeds))
	for i := 0; i < len(feeds); i++ {
		randomFeeds[i] = rand.Float64()
	}

	return randomFeeds, nil
}
