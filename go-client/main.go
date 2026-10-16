package main

import (
	"fast-updates-client/client"
	"fast-updates-client/config"
	"fast-updates-client/provider"
	"flag"
	"time"

	"github.com/flare-foundation/go-flare-common/pkg/logger"
)

func main() {
	flag.Parse()
	cfg, err := config.BuildConfig()
	if err != nil {
		logger.Fatalw("Invalid configuration", "error", err)
	}
	logger.Set(cfg.Logger.Config)

	valuesProvider := provider.NewHttpValueProvider(cfg.Client.ValueProviderUrl)

	client, err := client.CreateFastUpdatesClient(cfg, valuesProvider)
	if err != nil {
		logger.Fatalw("Failed to create client", "error", err)
	}

	for {
		err = client.Run(0, 0)
		if err != nil {
			logger.Errorw("Client loop failed, restarting", "error", err)
		}
		time.Sleep(200 * time.Millisecond)
	}
}
