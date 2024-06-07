package provider

import (
	"math"
	"math/big"
	"math/rand"

	"github.com/pkg/errors"
)

type ValuesDecimals struct {
	Feeds     []*big.Int
	Decimals  []int8
	Timestamp uint64
}

// ValuesProvider is an interface needed to provide current off-chain values.
type ValuesProvider interface {
	GetValues(feeds []FeedId) ([]*float64, error)
}

// GetDeltas calculates the deltas between the provider values and the chain values for each feed.
// It returns the deltas as a byte slice, the deltas as a string, and any error that occurred.
// rawChainValues and providerValues contain values for supported feeds.
func GetDeltas(chainValues []float64, providerValues []float64, valueIndexToFeedIndex []int, scale, sampleSize *big.Int) ([]byte, string, error) {
	if len(chainValues) != len(providerValues) {
		return nil, "", errors.New("chain and provider values length mismatch")
	}
	expectedSampleSize, _ := sampleSize.Float64()
	expectedSampleSize /= math.Pow(2, 120)
	// calculate the approx. expected change that all updates in one block do
	scaleDiff, _ := new(big.Int).Sub(scale, new(big.Int).Exp(big.NewInt(2), big.NewInt(127), nil)).Float64()
	scaleDiff = (scaleDiff / math.Pow(2, 127)) * expectedSampleSize

	lastFeedIndex := valueIndexToFeedIndex[len(valueIndexToFeedIndex)-1]
	deltasList := make([]byte, lastFeedIndex+1)
	for index := range deltasList {
		deltasList[index] = '0'
	}

	for i := 0; i < len(chainValues); i++ {
		delta := byte('0')
		diff := math.Abs(providerValues[i]-chainValues[i]) / chainValues[i]

		if diff > scaleDiff {
			if providerValues[i] > chainValues[i] {
				delta = '+'
			} else if providerValues[i] < chainValues[i] {
				delta = '-'
			}
		} else {
			r := rand.Float64()
			if r < diff/scaleDiff {
				if providerValues[i] > chainValues[i] {
					delta = '+'
				} else if providerValues[i] < chainValues[i] {
					delta = '-'
				}
			}
		}

		deltasList[valueIndexToFeedIndex[i]] = delta
	}

	deltasString := string(deltasList)
	deltas := StringToDeltas(deltasString)

	return deltas, deltasString, nil
}

// StringToDeltas converts a string representation of updates into a byte slice of deltas.
// Each character in the input string represents a delta value, where '+' represents an increment of 1
// and '-' represents an increment of 3. The deltas are packed into a byte slice, with each byte
// containing 4 delta values.
func StringToDeltas(update string) []byte {
	deltas := make([]byte, 0)
	k := 0
	var delta byte
	for i, part := range update {
		k = i % 4
		if k == 0 {
			delta = 0
		}
		if part == '+' {
			delta += 1 << (2 * (3 - k))
		}
		if part == '-' {
			delta += 3 << (2 * (3 - k))
		}
		if k == 3 {
			deltas = append(deltas, delta)
		}
	}
	if len(update)%4 != 0 {
		deltas = append(deltas, delta)
	}

	return deltas
}

// Sorts values according to feeds order.
func sortFeedValues(feeds []FeedId, feedValues []FeedValue) []*float64 {
	var values []*float64
	for _, feed := range feeds {
		for _, v := range feedValues {
			if v.Feed == feed {
				values = append(values, v.Value)
				break
			}
		}
	}
	return values
}
