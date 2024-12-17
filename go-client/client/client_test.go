package client_test

import (
	"context"
	"fast-updates-client/client"
	"fast-updates-client/config"
	"fast-updates-client/logger"
	"fast-updates-client/provider"
	"fast-updates-client/tests/test_utils"
	"math"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
)

func TestClient(t *testing.T) {
	chainNode := os.Getenv("CHAIN_NODE")
	chainAddress := ""
	valueProviderUrl := ""
	if chainNode == "docker_ganache" {
		chainAddress = "http://ganache:8545"
		valueProviderUrl = "http://value-provider:3101/feed-values/0"
	} else {
		// running a ganache node and an external provider that returns fixed values for testing
		logger.Info("starting a ganache chain node and data provider")
		// Can set VALUE_PROVIDER_IMPL to "fixed" or "random" to return 0.01 or random values for all feeds.
		cmd := exec.Command("bash", "-c", "docker compose up ganache  --detach && VALUE_PROVIDER_IMPL=random docker compose up value-provider --detach")
		cmd.Dir = "../tests"
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		go cmd.Run() //nolint:errcheck

		chainAddress = "http://127.0.0.1:8545"
		valueProviderUrl = "http://localhost:3101/feed-values/0"
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
		SigningPrivateKey:        "0xd49743deccbccc5dc7baa8e69e5be03298da8688a15dd202e20f15d5e0e9a9fb",
		SortitionPrivateKey:      "0xd49743deccbccc5dc7baa8e69e5be03298da8688a15dd202e20f15d5e0e9a9fb",
		SubmissionWindow:         4,
		MaxWeight:                1024,
		FetchCurrentFeedsAddress: "0xeAD9C93b79Ae7C1591b1FB5323BD777E86e150d4",
		FetchCurrentFeedsValue:   "10000",
	}
	cfgTransactions := config.TransactionsConfig{
		Accounts: []string{"0xd49743deccbccc5dc7baa8e69e5be03298da8688a15dd202e20f15d5e0e9a9fb",
			"0x23c601ae397441f3ef6f1075dcb0031ff17fb079837beadaf3c84d96c6f3e569",
			"0xee9d129c1997549ee09c0757af5939b2483d80ad649a0eda68e8b0357ad11131"},
		GasLimit:           80000000,
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

	updatesProvider := provider.NewHttpValueProvider(valueProviderUrl)

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

	feedIds, err := client.GetCurrentFeedIds()
	if err != nil {
		t.Fatal(err)
	}
	indexes := make([]int, len(feedIds))
	for i := range indexes {
		indexes[i] = i
	}
	startingFeeds, _, err := client.GetFeeds(indexes)
	if err != nil {
		t.Fatal(err)
	}

	err = client.Run(blockNum, blockNum+10)
	if err != nil {
		t.Fatal(err)
	}
	client.Stop()

	feeds, _, err := client.GetFeeds(indexes)
	if err != nil {
		t.Fatal(err)
	}
	scaleBig, err := client.GetScale()
	if err != nil {
		t.Fatal(err)
	}
	scaleFloat, _ := scaleBig.Float64()
	scale := scaleFloat / math.Pow(2, 127)

	downDockerContainers()
	if client.Stats.NumUpdates == 0 {
		t.Fatal("no updates submitted")
	}

	if client.Stats.NumSuccessfulUpdates == 0 {
		t.Fatal("no successful update")
	}

	for i, val := range feeds {
		// all updates are expected to be negative
		expectedVal := startingFeeds[i] * math.Pow(scale, -float64(client.Stats.NumSuccessfulUpdates))
		if expectedVal*0.999 > val && expectedVal*1.001 < val {
			if err != nil {
				t.Fatal("final feed values not correct:", expectedVal, val)
			}
		}
	}
}

func downDockerContainers() {
	cmd := exec.Command("bash", "-c", "docker compose down ganache value-provider")
	cmd.Dir = "../tests"
	cmd.Run() //nolint:errcheck
}
