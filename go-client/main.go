package main

import (
	"fast-updates-client/client"
	"fast-updates-client/config"
	"fast-updates-client/logger"
	"fast-updates-client/provider"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	flag.Parse()
	cfg, err := config.BuildConfig()
	if err != nil {
		fmt.Printf("Config error: %v\n", err)
		os.Exit(1)
	}
	config.GlobalConfigCallback.Call(cfg)

	valuesProvider := provider.NewHttpValueProvider(cfg.Client.ValueProviderUrl)
	if err != nil {
		logger.Fatal("Error: %s", err)
		return
	}

	client, err := client.CreateFastUpdatesClient(cfg, valuesProvider)
	if err != nil {
		logger.Fatal("Error: %s", err)
		return
	}

	for {
		err = client.Run(0, 0)
		if err != nil {
			logger.Error("Error: %s", err)
			logger.Info("Restarting")
		}
		time.Sleep(200 * time.Millisecond)
	}
}
