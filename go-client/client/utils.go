package client

import (
	"fast-updates-client/provider"
	"math/big"
	"time"
)

func WaitForBlock(txQueue *TransactionQueue, blockNum uint64) error {
	for {
		if txQueue.CurrentBlockNum < blockNum {
			time.Sleep(200 * time.Millisecond)
		} else {
			return nil
		}
	}
}

// RawChainValuesToFloats converts each on-chain feed value to a float64.
//
// A feed is reported as an integer value with a signed decimals exponent, where
// the real value is value * 10^(-decimals). Positive decimals therefore divide
// (12345 with decimals=3 -> 12.345) and negative decimals multiply (254000000
// with decimals=-1 -> 2.54e9, as happens for feeds priced above the int32
// ceiling in a low-value quote unit).
//
// The conversion runs over the full big.Ints via big.Rat so it stays exact for
// any feed magnitude or decimals sign, without narrowing to int64.
func RawChainValuesToFloats(rawChainValues provider.ValuesDecimals) []float64 {
	floatValues := make([]float64, len(rawChainValues.Feeds))
	for i, feedValue := range rawChainValues.Feeds {
		decimals := int64(rawChainValues.Decimals[i])
		// big.Int.Exp rejects a negative exponent, so raise 10 to |decimals|
		// and let the sign decide which side of the fraction it goes on below.
		// pow10 is always >= 1 (10^0 = 1), so it's never a zero denominator.
		absDecimals := decimals
		if absDecimals < 0 {
			absDecimals = -absDecimals
		}
		pow10 := new(big.Int).Exp(big.NewInt(10), big.NewInt(absDecimals), nil)

		ratio := new(big.Rat)
		if decimals >= 0 {
			ratio.SetFrac(feedValue, pow10) // feedValue / 10^decimals
		} else {
			ratio.SetInt(new(big.Int).Mul(feedValue, pow10)) // feedValue * 10^|decimals|
		}
		floatValues[i], _ = ratio.Float64()
	}
	return floatValues
}

func CheckBalances(balances []*big.Int, minBalance float64) bool {
	for _, balance := range balances {
		balanceFloat, _ := balance.Float64()
		if balanceFloat < minBalance {
			return false
		}
	}

	return true
}
