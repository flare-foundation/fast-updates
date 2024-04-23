package main

import (
	"flag"
	"fmt"
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
	}
}
