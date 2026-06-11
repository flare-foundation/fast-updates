package client_test

import (
	"fast-updates-client/client"
	"fast-updates-client/config"
	"fast-updates-client/provider"
	"math/big"
	"testing"
	"time"
)

func TestWaitForBlock_ReturnsWhenReached(t *testing.T) {
	q := client.NewTransactionQueue(nil, nil, config.TransactionsConfig{})
	q.CurrentBlockNum = 100
	start := time.Now()
	if err := client.WaitForBlock(q, 50); err != nil {
		t.Fatalf("WaitForBlock: %v", err)
	}
	if time.Since(start) > time.Second {
		t.Errorf("WaitForBlock should return immediately when current block already reached")
	}
}

func TestWaitForBlock_WaitsThenReturns(t *testing.T) {
	q := client.NewTransactionQueue(nil, nil, config.TransactionsConfig{})
	q.CurrentBlockNum = 0
	done := make(chan error, 1)
	go func() {
		done <- client.WaitForBlock(q, 5)
	}()
	time.Sleep(100 * time.Millisecond)
	q.CurrentBlockNum = 10
	select {
	case err := <-done:
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	case <-time.After(3 * time.Second):
		t.Fatal("WaitForBlock did not return after target block reached")
	}
}

func TestRawChainValuesToFloats_TableDriven(t *testing.T) {
	cases := []struct {
		name     string
		feeds    []*big.Int
		decimals []int8
		want     []float64
	}{
		{
			name:     "two_decimals",
			feeds:    []*big.Int{big.NewInt(100), big.NewInt(200), big.NewInt(300)},
			decimals: []int8{2, 2, 2},
			want:     []float64{1.0, 2.0, 3.0},
		},
		{
			name:     "mixed_decimals",
			feeds:    []*big.Int{big.NewInt(12345), big.NewInt(1)},
			decimals: []int8{3, 0},
			want:     []float64{12.345, 1.0},
		},
		{
			// 8 decimals, like a typical wei-ish feed.
			name:     "high_positive_decimals",
			feeds:    []*big.Int{big.NewInt(150_000_000)},
			decimals: []int8{8},
			want:     []float64{1.5},
		},
		{
			// Negative decimals (int8 signed exponent) scale the value up:
			// real value = feedValue * 10^|decimals|. Arises for feeds priced
			// above the int32 ceiling in the quote unit, e.g. a BTC pair quoted
			// in a low-value unit.
			name:     "negative_decimals",
			feeds:    []*big.Int{big.NewInt(254_000_000), big.NewInt(500_000_000)},
			decimals: []int8{-1, -2},
			want:     []float64{2_540_000_000.0, 50_000_000_000.0},
		},
		{
			// Zero value on the negative-decimals branch still yields 0.
			name:     "negative_decimals_zero_value",
			feeds:    []*big.Int{big.NewInt(0)},
			decimals: []int8{-2},
			want:     []float64{0.0},
		},
		{
			name:     "zero_value",
			feeds:    []*big.Int{big.NewInt(0)},
			decimals: []int8{2},
			want:     []float64{0.0},
		},
		{
			name:     "empty",
			feeds:    []*big.Int{},
			decimals: []int8{},
			want:     []float64{},
		},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			in := provider.ValuesDecimals{Feeds: tc.feeds, Decimals: tc.decimals}
			got := client.RawChainValuesToFloats(in)
			if len(got) != len(tc.want) {
				t.Fatalf("len: want %d got %d", len(tc.want), len(got))
			}
			for i := range tc.want {
				if got[i] != tc.want[i] {
					t.Errorf("idx %d: want %g got %g", i, tc.want[i], got[i])
				}
			}
		})
	}
}

// A feed value above the int64 range must not truncate or flip sign — the old
// Int64()-based conversion turned 2^64-1 into -1.
func TestRawChainValuesToFloats_LargeValueNoTruncation(t *testing.T) {
	huge, _ := new(big.Int).SetString("18446744073709551615", 10) // 2^64 - 1
	in := provider.ValuesDecimals{
		Feeds:    []*big.Int{huge},
		Decimals: []int8{0},
	}
	got := client.RawChainValuesToFloats(in)
	if len(got) != 1 {
		t.Fatalf("expected one result, got %d", len(got))
	}
	if got[0] < 1.8e19 {
		t.Errorf("large value truncated/flipped: got %g, want ~1.8e19", got[0])
	}
}

// High decimals must not panic — the old code computed 10^decimals via Int64(),
// which truncated to 0 for decimals >= 64 and made big.NewRat divide by zero.
func TestRawChainValuesToFloats_HighDecimalsNoPanic(t *testing.T) {
	in := provider.ValuesDecimals{
		Feeds:    []*big.Int{big.NewInt(1_000_000_000)},
		Decimals: []int8{100},
	}
	got := client.RawChainValuesToFloats(in) // must not panic
	if len(got) != 1 {
		t.Fatalf("expected one result, got %d", len(got))
	}
	if got[0] <= 0 || got[0] >= 1 {
		t.Errorf("expected a tiny positive value, got %g", got[0])
	}
}

func TestCheckBalances(t *testing.T) {
	cases := []struct {
		name       string
		balances   []*big.Int
		minBalance float64
		want       bool
	}{
		{
			name:       "all_above",
			balances:   []*big.Int{big.NewInt(100), big.NewInt(200)},
			minBalance: 50,
			want:       true,
		},
		{
			name:       "one_below",
			balances:   []*big.Int{big.NewInt(100), big.NewInt(10)},
			minBalance: 50,
			want:       false,
		},
		{
			name:       "exact_min_not_enough",
			balances:   []*big.Int{big.NewInt(50)},
			minBalance: 50,
			want:       true, // strictly less-than, equality passes
		},
		{
			name:       "empty",
			balances:   []*big.Int{},
			minBalance: 100,
			want:       true,
		},
		{
			name:       "zero_balance_fails",
			balances:   []*big.Int{big.NewInt(0)},
			minBalance: 1,
			want:       false,
		},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := client.CheckBalances(tc.balances, tc.minBalance)
			if got != tc.want {
				t.Errorf("want %v got %v", tc.want, got)
			}
		})
	}
}
