package provider_test

import (
	"fast-updates-client/provider"
	"math/big"
	"testing"
)

// scaleWithDiff returns a scale in the contract's 2^127 fixed-point encoding
// such that a single update multiplies a feed value by (1 + diff), together
// with an expected sample size of exactly 1 in the 2^120 encoding. With
// sampleSize = 1 the volatility band used by GetDeltas equals diff itself,
// which makes the deterministic/probabilistic regions of GetDeltas easy to
// steer in tests.
func scaleWithDiff(diff float64) (scale, sampleSize *big.Int) {
	one127 := new(big.Int).Exp(big.NewInt(2), big.NewInt(127), nil)
	diffFixed, _ := new(big.Float).Mul(big.NewFloat(diff), new(big.Float).SetInt(one127)).Int(nil)
	scale = new(big.Int).Add(one127, diffFixed)
	sampleSize = new(big.Int).Exp(big.NewInt(2), big.NewInt(120), nil)
	return scale, sampleSize
}

func TestGetDeltasDirection(t *testing.T) {
	// band is 0.1%; all diffs below are well outside it, so deltas are
	// deterministic signs
	scale, sampleSize := scaleWithDiff(0.001)

	cases := []struct {
		name           string
		chainValues    []float64
		providerValues []float64
		want           string
	}{
		{
			name:           "provider_above_chain_increments",
			chainValues:    []float64{100, 100, 100},
			providerValues: []float64{110, 120, 150},
			want:           "+++",
		},
		{
			name:           "provider_below_chain_decrements",
			chainValues:    []float64{100, 100, 100},
			providerValues: []float64{90, 50, 99},
			want:           "---",
		},
		{
			name:           "exactly_at_target_never_updates",
			chainValues:    []float64{100, 100, 100},
			providerValues: []float64{100, 100, 100},
			want:           "000",
		},
		{
			name:           "mixed_directions",
			chainValues:    []float64{100, 100, 100, 100},
			providerValues: []float64{110, 100, 90, 111},
			want:           "+0-+",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			indexes := make([]int, len(tc.chainValues))
			for i := range indexes {
				indexes[i] = i
			}
			deltas, deltasString, err := provider.GetDeltas(tc.chainValues, tc.providerValues, indexes, scale, sampleSize)
			if err != nil {
				t.Fatalf("GetDeltas: %v", err)
			}
			if deltasString != tc.want {
				t.Errorf("deltas string: want %q, got %q", tc.want, deltasString)
			}
			wantPacked := provider.StringToDeltas(tc.want)
			if string(deltas) != string(wantPacked) {
				t.Errorf("packed deltas: want %v, got %v", wantPacked, deltas)
			}
		})
	}
}

// TestGetDeltasMissingFeeds checks the index alignment when the provider only
// supports a subset of the on-chain feeds: valueIndexToFeedIndex maps each
// provided value to its on-chain feed index and every unmapped index must stay
// '0'. A misalignment here would silently push another feed's update.
func TestGetDeltasMissingFeeds(t *testing.T) {
	scale, sampleSize := scaleWithDiff(0.001)

	// on-chain feeds 0..6 exist, but only 1, 3 and 6 have provider values
	valueIndexToFeedIndex := []int{1, 3, 6}
	chainValues := []float64{100, 100, 100}
	providerValues := []float64{110, 90, 110}

	_, deltasString, err := provider.GetDeltas(chainValues, providerValues, valueIndexToFeedIndex, scale, sampleSize)
	if err != nil {
		t.Fatalf("GetDeltas: %v", err)
	}

	want := "0+0-00+"
	if deltasString != want {
		t.Errorf("deltas string: want %q, got %q", want, deltasString)
	}
}

// TestGetDeltasWithinBand checks the probabilistic region: when the relative
// difference is inside the volatility band, the delta may randomly stay '0'
// but must never have the wrong sign.
func TestGetDeltasWithinBand(t *testing.T) {
	scale, sampleSize := scaleWithDiff(0.001)

	sawZero, sawPlus := false, false
	for i := 0; i < 1000; i++ {
		// +0.05% diff, half of the 0.1% band
		_, deltasString, err := provider.GetDeltas([]float64{100}, []float64{100.05}, []int{0}, scale, sampleSize)
		if err != nil {
			t.Fatalf("GetDeltas: %v", err)
		}
		switch deltasString {
		case "0":
			sawZero = true
		case "+":
			sawPlus = true
		default:
			t.Fatalf("within-band delta must be '0' or '+', got %q", deltasString)
		}
	}
	// with p = 0.5 over 1000 draws, both outcomes are effectively guaranteed
	if !sawZero || !sawPlus {
		t.Errorf("expected both '0' and '+' outcomes within band, got zero=%v plus=%v", sawZero, sawPlus)
	}
}

func TestStringToDeltasPacking(t *testing.T) {
	cases := []struct {
		name   string
		update string
		want   []byte
	}{
		// '+' encodes as 1 and '-' as 3 in 2-bit slots packed
		// big-endian within each byte
		{name: "empty", update: "", want: []byte{}},
		{name: "single_plus", update: "+", want: []byte{0b01000000}},
		{name: "single_minus", update: "-", want: []byte{0b11000000}},
		{name: "all_slots_one_byte", update: "+--+", want: []byte{0b01111101}},
		{name: "zeros_only", update: "0000", want: []byte{0}},
		{
			name:   "not_multiple_of_four",
			update: "+-0+-0",
			want:   []byte{0b01110001, 0b11000000},
		},
		{
			name:   "nine_feeds_three_bytes",
			update: "00000000+",
			want:   []byte{0, 0, 0b01000000},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := provider.StringToDeltas(tc.update)
			if len(got) != len(tc.want) {
				t.Fatalf("packed length: want %d, got %d", len(tc.want), len(got))
			}
			for i := range tc.want {
				if got[i] != tc.want[i] {
					t.Errorf("byte %d: want %08b, got %08b", i, tc.want[i], got[i])
				}
			}
		})
	}
}
