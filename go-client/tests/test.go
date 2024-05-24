package main

import (
	"flag"
	"fmt"
	"math/big"
	"os"
	"strconv"

	"fast-updates-client/config"
	"fast-updates-client/logger"
	"fast-updates-client/tests/test_utils"
)

func main() {
	flag.Parse()
	cfg, err := config.BuildConfig()
	if err != nil {
		fmt.Println("Config error: %w", err)
		return
	}
	config.GlobalConfigCallback.Call(cfg)

	arg := os.Args[3]
	if arg == "deploy" {
		test_utils.Deploy(cfg)
	} else if arg == "register" {
		numEpoch, err := strconv.Atoi(os.Args[4])
		if err != nil {
			if os.Args[4] == "" {
				numEpoch = 1
			} else {
				logger.Fatal("Config error: %s", err)
				return
			}
		}
		err = test_utils.Register(cfg, numEpoch)
		if err != nil {
			logger.Fatal("Registering error: %s", err)
			return
		}
	} else if arg == "incentive" {
		rangeIncreaseStr := os.Args[4]
		rangeIncrease, err := strconv.ParseFloat(rangeIncreaseStr, 64)
		if err != nil {
			logger.Fatal("Incentive error: %s", err)
			return
		}

		sampleCostStr := os.Args[5]
		var sampleCost *big.Int
		var check bool
		if sampleCostStr == "" {
			sampleCost = big.NewInt(0)
		} else {
			sampleCost, check = new(big.Int).SetString(sampleCostStr, 10)
			if !check {
				logger.Fatal("Could not read sample cost")
				return
			}
		}

		test_utils.Incentivize(cfg, rangeIncrease, sampleCost)
	}
}
