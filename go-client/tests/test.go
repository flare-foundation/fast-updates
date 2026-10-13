package main

import (
	"flag"
	"fmt"
	"math/big"
	"os"
	"strconv"

	"fast-updates-client/config"
	"fast-updates-client/tests/test_utils"

	"github.com/flare-foundation/go-flare-common/pkg/logger"
)

func main() {
	flag.Parse()
	cfg, err := config.BuildConfig()
	if err != nil {
		fmt.Println("Config error: %w", err)
		return
	}
	logger.Set(cfg.Logger.Config)

	arg := os.Args[3]
	switch arg {
	case "deploy":
		test_utils.Deploy(cfg)
	case "register":
		numEpoch, err := strconv.Atoi(os.Args[4])
		if err != nil {
			if os.Args[4] == "" {
				numEpoch = 1
			} else {
				logger.Fatalf("Config error: %s", err)
				return
			}
		}
		err = test_utils.Register(cfg, numEpoch)
		if err != nil {
			logger.Fatalf("Registering error: %s", err)
			return
		}
	case "incentive":
		rangeIncreaseStr := os.Args[4]
		rangeIncrease, err := strconv.ParseFloat(rangeIncreaseStr, 64)
		if err != nil {
			logger.Fatalf("Incentive error: %s", err)
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
				logger.Fatalf("Could not read sample cost")
				return
			}
		}

		test_utils.Incentivize(cfg, rangeIncrease, sampleCost)
	}
}
