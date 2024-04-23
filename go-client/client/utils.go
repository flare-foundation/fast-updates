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

// rawChainPricesToFloats converts prices with decimals to float64 values.
func RawChainValuesToFloats(rawChainValues provider.ValuesDecimals) []float64 {
	floatValues := make([]float64, len(rawChainValues.Feeds))
	for i, feedValue := range rawChainValues.Feeds {
		decimals := big.NewInt(int64(rawChainValues.Decimals[i]))
		exp := decimals.Exp(big.NewInt(10), decimals, nil)
		floatValues[i], _ = big.NewRat(feedValue.Int64(), exp.Int64()).Float64()
	}
	return floatValues
}
