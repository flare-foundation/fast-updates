package client_test

import (
	"context"
	"fast-updates-client/client"
	"fast-updates-client/config"
	"fast-updates-client/logger"
	"fast-updates-client/provider"
	"fast-updates-client/tests/test_utils"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
)

func TestClient(t *testing.T) {
	chainNode := os.Getenv("CHAIN_NODE")
	chainAddress := ""
	valueProviderBaseUrl := ""
	if chainNode == "docker_ganache" {
		chainAddress = "http://ganache:8545"
		valueProviderBaseUrl = "http://value-provider:3101"
	} else {
		// running a ganache node
		logger.Info("starting a ganache chain node")
		// cmd := exec.Command("bash", "-c", "docker run --publish 8544:8545 trufflesuite/ganache:latest --chain.hardfork=\"london\" --miner.blockTime=5 --wallet.accounts \"0xc5e8f61d1ab959b397eecc0a37a6517b8e67a0e7cf1f4bce5591f3ed80199122, 10000000000000000000000\" \"0xd49743deccbccc5dc7baa8e69e5be03298da8688a15dd202e20f15d5e0e9a9fb, 10000000000000000000000\" \"0x23c601ae397441f3ef6f1075dcb0031ff17fb079837beadaf3c84d96c6f3e569, 10000000000000000000000\" \"0xee9d129c1997549ee09c0757af5939b2483d80ad649a0eda68e8b0357ad11131, 10000000000000000000000\"")
		cmd := exec.Command("bash", "-c", "docker compose up ganache")
		cmd.Dir = "../tests"
		// cmd.Stdout = os.Stdout
		// cmd.Stderr = os.Stderr
		go cmd.Run() //nolint:errcheck
		chainAddress = "http://127.0.0.1:8545"

		// runs an external provider that returns fixed values for testing
		runValueProvider()
		valueProviderBaseUrl = "http://localhost:3101"
	}

	// set chain parameters
	cfgChain := config.ChainConfig{NodeURL: chainAddress, ChainId: 1337}
	// wait for the chain node to be ready
	for {
		time.Sleep(10 * time.Second)
		client, err := ethclient.Dial(cfgChain.NodeURL)
		if err != nil {
			continue
		}
		_, err = client.BlockByNumber(context.Background(), nil)
		if err == nil {
			break
		}
	}

	// set configuration parameters
	cfgClient := config.FastUpdateClientConfig{
		SigningPrivateKey:   "0xd49743deccbccc5dc7baa8e69e5be03298da8688a15dd202e20f15d5e0e9a9fb",
		SortitionPrivateKey: "0xd49743deccbccc5dc7baa8e69e5be03298da8688a15dd202e20f15d5e0e9a9fb",
		SubmissionWindow:    5,
		MaxWeight:           1024,
	}
	cfgTransactions := config.TransactionsConfig{
		Accounts: []string{"0xd49743deccbccc5dc7baa8e69e5be03298da8688a15dd202e20f15d5e0e9a9fb",
			"0x23c601ae397441f3ef6f1075dcb0031ff17fb079837beadaf3c84d96c6f3e569",
			"0xee9d129c1997549ee09c0757af5939b2483d80ad649a0eda68e8b0357ad11131"},
		GasLimit:           8000000,
		GasPriceMultiplier: 1.2,
	}
	cfgLog := config.LoggerConfig{Level: "DEBUG", Console: true, File: "../logger/logs/flare-ftso-indexer_test.log"}

	cfg := config.Config{Client: cfgClient, Chain: cfgChain, Logger: cfgLog, Transactions: cfgTransactions}
	config.GlobalConfigCallback.Call(cfg)

	logger.Info("deploying contracts")
	contracts := test_utils.Deploy(&cfg)

	cfg.Client.FastUpdaterAddress = contracts.FastUpdater.Hex()
	cfg.Client.FastUpdatesConfigurationAddress = contracts.FastUpdatesConfiguration.Hex()
	cfg.Client.FlareSystemManagerAddress = contracts.Mock.Hex()
	cfg.Client.MockAddress = contracts.Mock.Hex()
	cfg.Client.IncentiveManagerAddress = contracts.IncentiveManager.Hex()

	updatesProvider := provider.NewHttpValueProvider(valueProviderBaseUrl)

	client, err := client.CreateFastUpdatesClient(&cfg, updatesProvider)
	if err != nil {
		t.Fatal(err)
	}

	logger.Info("registering for this and next epoch")

	err = test_utils.Register(&cfg, 2)
	if err != nil {
		t.Fatal("Registering error: %w", err)
	}

	blockNum, err := client.CurrentBlockNumber()
	if err != nil {
		t.Fatal(err)
	}

	err = client.Run(blockNum, blockNum+10)
	if err != nil {
		t.Fatal(err)
	}
	client.Stop()

	if chainNode != "docker_ganache" {
		time.Sleep(time.Second)
		// stopping a ganache node
		cmd := exec.Command("bash", "-c", "docker compose stop ganache")
		cmd.Dir = "../tests"
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err = cmd.Run() //nolint:errcheck
		if err != nil {
			t.Fatal(err)
		}
		stopValueProvider()
	}
}

// Can set VALUE_PROVIDER_IMPL to "fixed" or "random" to return 0.01 or random values for all feeds.
func runValueProvider() {
	cmd := exec.Command("bash", "-c", "VALUE_PROVIDER_IMPL=random docker compose up value-provider")
	cmd.Dir = "../tests"
	go cmd.Run() //nolint:errcheck
}

func stopValueProvider() {
	cmd := exec.Command("bash", "-c", "docker compose stop value-provider")
	cmd.Dir = "../tests"
	go cmd.Run() //nolint:errcheck
}
