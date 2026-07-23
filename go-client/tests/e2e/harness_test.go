package e2e

import (
	"context"
	"math/big"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"fast-updates-client/client"
	"fast-updates-client/config"
	"fast-updates-client/contracts-interface/fast_updates_configuration"
	"fast-updates-client/logger"
	"fast-updates-client/tests/test_utils"
)

const (
	// pinned to the v7.9.2 manifest digest so upstream tag changes cannot
	// break the suite. Always run as linux/amd64 (see ganachePlatform):
	// the image's arm64 variant ships an amd64 node binary (exec format
	// error) and upstream is sunset, so on ARM hosts the amd64 variant
	// under emulation (Rosetta on Apple Silicon) is the working option.
	ganacheImage    = "trufflesuite/ganache:v7.9.2@sha256:7b58eeff2a5cfeb62ee19fe16ce7c5b88b9e9bae0f627719121609e9e188b730"
	ganachePlatform = "linux/amd64"
	ganachePort     = "8547"
	containerName   = "fast-updates-e2e-ganache"
	chainId         = 1337
	deployerKey     = "0xd49743deccbccc5dc7baa8e69e5be03298da8688a15dd202e20f15d5e0e9a9fb"
	// governanceKey deploys the contracts and signs feed add/remove
	// transactions. It must NOT be one of the client's transaction accounts:
	// a scenario changing feeds while the client runs would otherwise race
	// the client for the account nonce and lose its transaction.
	governanceKey  = "0x87630b2d1de0fbd5044eb6891b3d9d98c34c8d310c852f98550ba774480e47cc"
	numFeeds       = 9 // feeds added by test_utils.Deploy
	startFeedValue = 1000.0
)

// rpcHost is where the ganache container's published port is reachable. It is
// 127.0.0.1 for local runs; in CI with docker-in-docker the container port is
// published on the dind service host instead, so it can be overridden with
// the FAST_UPDATES_E2E_RPC_HOST environment variable.
func rpcHost() string {
	if host := os.Getenv("FAST_UPDATES_E2E_RPC_HOST"); host != "" {
		return host
	}
	return "127.0.0.1"
}

func nodeURL() string {
	return "http://" + rpcHost() + ":" + ganachePort
}

// blockTimeSeconds is the chain's block time: 2s by default (close to
// Flare's real cadence — sortition alone costs ~1s per block at weight 4096,
// so against a 1s chain the client falls behind the submission window and
// aborts on slower machines), overridable with FAST_UPDATES_E2E_BLOCK_TIME.
func blockTimeSeconds() int {
	value := os.Getenv("FAST_UPDATES_E2E_BLOCK_TIME")
	if value == "" {
		return 2
	}
	seconds, err := strconv.Atoi(value)
	if err != nil || seconds < 1 {
		panic("invalid FAST_UPDATES_E2E_BLOCK_TIME: " + value)
	}
	return seconds
}

// same funded accounts as tests/docker-compose.yaml
var ganacheKeys = []string{
	"0xc5e8f61d1ab959b397eecc0a37a6517b8e67a0e7cf1f4bce5591f3ed80199122",
	"0xd49743deccbccc5dc7baa8e69e5be03298da8688a15dd202e20f15d5e0e9a9fb",
	"0x23c601ae397441f3ef6f1075dcb0031ff17fb079837beadaf3c84d96c6f3e569",
	"0xee9d129c1997549ee09c0757af5939b2483d80ad649a0eda68e8b0357ad11131",
	"0x87630b2d1de0fbd5044eb6891b3d9d98c34c8d310c852f98550ba774480e47cc",
}

// TestEnv holds everything a scenario needs to drive the client and verify
// the chain afterwards.
type TestEnv struct {
	t         *testing.T
	Cfg       *config.Config
	Contracts test_utils.ContractAddresses
	EthClient *ethclient.Client
	Provider  *ScriptedProvider
}

// NewTestEnv starts a fresh ganache chain in docker (see blockTimeSeconds
// for the block time), deploys the fast updates contracts (with the mocked
// Flare system) and registers the test voter for the current and next reward
// epoch.
func NewTestEnv(t *testing.T) *TestEnv {
	if testing.Short() {
		t.Skip("skipping e2e scenario test in -short mode (requires docker)")
	}
	if _, err := exec.LookPath("docker"); err != nil {
		t.Skip("docker not available, skipping e2e scenario test")
	}

	startGanache(t)

	cfg := &config.Config{
		Client: config.FastUpdateClientConfig{
			SigningPrivateKey:   deployerKey,
			SortitionPrivateKey: deployerKey,
			SubmissionWindow:    int(test_utils.SUBMISSION_WINDOW),
			// the mock registers our voter as the only one, so the contract
			// normalizes us to the full 4096 virtual-provider weight; with an
			// expected sample size of 1 this yields ~1 update per block
			MaxWeight:                4096,
			FetchCurrentFeedsAddress: "0xeAD9C93b79Ae7C1591b1FB5323BD777E86e150d4",
			FetchCurrentFeedsValue:   "10000",
		},
		Transactions: config.TransactionsConfig{
			Accounts: []string{
				"0xd49743deccbccc5dc7baa8e69e5be03298da8688a15dd202e20f15d5e0e9a9fb",
				"0x23c601ae397441f3ef6f1075dcb0031ff17fb079837beadaf3c84d96c6f3e569",
				"0xee9d129c1997549ee09c0757af5939b2483d80ad649a0eda68e8b0357ad11131",
			},
			GasLimit:                 8000000,
			BaseFeeMultiplier:        4,
			MaxPriorityFeeMultiplier: 2,
		},
		Chain:  config.ChainConfig{NodeURL: nodeURL(), ChainId: chainId},
		Logger: config.LoggerConfig{Level: "INFO", Console: true, File: "../../logger/logs/e2e_test.log"},
	}
	config.GlobalConfigCallback.Call(*cfg)

	ethClient, err := ethclient.Dial(nodeURL())
	if err != nil {
		t.Fatalf("dialing ganache: %v", err)
	}

	logger.Info("deploying contracts")
	// deploy from the governance account (test_utils.Deploy uses account 0),
	// keeping it disjoint from the client's transaction accounts
	deployCfg := *cfg
	deployCfg.Transactions.Accounts = append([]string{governanceKey}, cfg.Transactions.Accounts[1:]...)
	contracts := test_utils.Deploy(&deployCfg)
	cfg.Client.FastUpdaterAddress = contracts.FastUpdater.Hex()
	cfg.Client.FastUpdatesConfigurationAddress = contracts.FastUpdatesConfiguration.Hex()
	cfg.Client.FlareSystemManagerAddress = contracts.Mock.Hex()
	cfg.Client.MockAddress = contracts.Mock.Hex()
	cfg.Client.IncentiveManagerAddress = contracts.IncentiveManager.Hex()

	logger.Info("registering voter for this and next epoch")
	if err := test_utils.Register(cfg, 2); err != nil {
		t.Fatalf("registering voter: %v", err)
	}

	return &TestEnv{
		t:         t,
		Cfg:       cfg,
		Contracts: contracts,
		EthClient: ethClient,
		Provider:  NewScriptedProvider(),
	}
}

func startGanache(t *testing.T) {
	// remove any leftover container from an earlier (possibly aborted) run
	exec.Command("docker", "rm", "-f", containerName).Run() //nolint:errcheck

	// bind only to loopback locally; when the RPC host is overridden (CI with
	// docker-in-docker) the port must be published on all interfaces of the
	// docker host instead
	publish := "127.0.0.1:" + ganachePort + ":8545"
	if os.Getenv("FAST_UPDATES_E2E_RPC_HOST") != "" {
		publish = ganachePort + ":8545"
	}
	blockTime := blockTimeSeconds()
	// no --rm: if the container dies on startup its logs must stay available
	// for the diagnostics below (cleanup removes it explicitly)
	args := []string{
		"run", "-d", "--name", containerName,
		"--platform", ganachePlatform,
		"-p", publish,
		ganacheImage,
		"--chain.hardfork", "london",
		"--gasLimit", "100000000",
		"--miner.blockTime=" + strconv.Itoa(blockTime),
	}
	for _, key := range ganacheKeys {
		args = append(args, "--wallet.accounts", key+",10000000000000000000000")
	}

	out, err := exec.Command("docker", args...).CombinedOutput()
	if err != nil {
		t.Fatalf("starting ganache container: %v\n%s", err, out)
	}
	t.Cleanup(func() {
		exec.Command("docker", "rm", "-f", containerName).Run() //nolint:errcheck
	})

	// wait until the RPC endpoint serves blocks
	deadline := time.Now().Add(60 * time.Second)
	for {
		if time.Now().After(deadline) {
			t.Fatalf("ganache did not become ready within 60s; container logs:\n%s", containerLogs())
		}
		time.Sleep(500 * time.Millisecond)

		running, _ := exec.Command("docker", "inspect", "-f", "{{.State.Running}}", containerName).Output()
		if strings.TrimSpace(string(running)) == "false" {
			logs := containerLogs()
			if strings.Contains(logs, "exec format error") {
				// ganache has no working arm64 build (its arm64 variant ships
				// an amd64 node binary and upstream is sunset), so the suite
				// needs the host to run linux/amd64 images
				t.Skipf("host cannot execute the linux/amd64 ganache image (no amd64 emulation).\n"+
					"Docker Desktop on Apple Silicon supports this out of the box (enable Rosetta); "+
					"on Linux ARM hosts install binfmt emulation, e.g.:\n"+
					"  docker run --privileged --rm tonistiigi/binfmt --install amd64\n"+
					"container logs:\n%s", logs)
			}
			t.Fatalf("ganache container exited during startup; logs:\n%s", logs)
		}

		ethClient, err := ethclient.Dial(nodeURL())
		if err != nil {
			continue
		}
		_, err = ethClient.BlockNumber(context.Background())
		ethClient.Close()
		if err == nil {
			return
		}
	}
}

func containerLogs() string {
	logs, err := exec.Command("docker", "logs", containerName).CombinedOutput()
	if err != nil {
		return "(could not fetch container logs: " + err.Error() + ")"
	}
	return string(logs)
}

// RunClient runs the fast updates client for the given number of blocks
// starting at the current head and returns the block range [start, end] that
// the verifier should scan (end is the head after all queued transactions
// were mined).
func (env *TestEnv) RunClient(numBlocks uint64) (startBlock, endBlock uint64, stats client.UpdatesStats) {
	startBlock, wait := env.RunClientAsync(numBlocks)
	endBlock, stats = wait()
	return startBlock, endBlock, stats
}

// RunClientAsync starts the fast updates client in the background so the
// scenario can manipulate the chain (e.g. replace feeds) while it runs. It
// returns the client's start block and a wait function that blocks until the
// run finishes and returns the end block and the client stats.
func (env *TestEnv) RunClientAsync(numBlocks uint64) (uint64, func() (uint64, client.UpdatesStats)) {
	fastUpdatesClient, err := client.CreateFastUpdatesClient(env.Cfg, env.Provider)
	if err != nil {
		env.t.Fatalf("creating client: %v", err)
	}

	startBlock, err := fastUpdatesClient.CurrentBlockNumber()
	if err != nil {
		env.t.Fatalf("getting current block: %v", err)
	}

	runErr := make(chan error, 1)
	go func() {
		runErr <- fastUpdatesClient.Run(startBlock, startBlock+numBlocks)
	}()

	wait := func() (uint64, client.UpdatesStats) {
		err := <-runErr
		fastUpdatesClient.Stop()
		if err != nil {
			env.t.Fatalf("running client: %v", err)
		}
		endBlock, err := fastUpdatesClient.CurrentBlockNumber()
		if err != nil {
			env.t.Fatalf("getting end block: %v", err)
		}
		return endBlock, fastUpdatesClient.Stats
	}
	return startBlock, wait
}

// WaitForChainBlock blocks until the chain head reaches the given height,
// failing the test if it does not get there in roughly twice the expected
// time (e.g. the ganache container died and blocks stopped).
func (env *TestEnv) WaitForChainBlock(blockNum uint64) {
	head, err := env.EthClient.BlockNumber(context.Background())
	if err != nil {
		env.t.Fatalf("getting head block: %v", err)
	}
	blocksToWait := int64(blockNum) - int64(head)
	timeout := time.Duration(2*(blocksToWait+15)*int64(blockTimeSeconds())) * time.Second
	deadline := time.Now().Add(timeout)

	for {
		head, err := env.EthClient.BlockNumber(context.Background())
		if err == nil && head >= blockNum {
			return
		}
		if time.Now().After(deadline) {
			env.t.Fatalf("chain did not reach block %d within %s (head %d, last error %v) — did the ganache container die?",
				blockNum, timeout, head, err)
		}
		time.Sleep(200 * time.Millisecond)
	}
}

func (env *TestEnv) waitMinedSuccessful(tx *types.Transaction, what string) uint64 {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	receipt, err := bind.WaitMined(ctx, env.EthClient, tx)
	if err != nil {
		env.t.Fatalf("waiting for %s tx %s: %v (transaction dropped or not mined within 60s)",
			what, tx.Hash(), err)
	}
	if receipt.Status == 0 {
		env.t.Fatalf("%s transaction reverted", what)
	}
	return receipt.BlockNumber.Uint64()
}

// RemoveFeed delists the feed with the given deployment index through the
// FastUpdatesConfiguration contract, zeroing its feed id on-chain while
// keeping the indices of the remaining feeds unchanged. It returns the block
// the removal was mined in.
func (env *TestEnv) RemoveFeed(feedIndex int) uint64 {
	configContract := env.configContract()

	feedId := [21]byte{}
	feedId[20] = byte(feedIndex + 1)

	tx, err := configContract.RemoveFeeds(env.transactOpts(), [][21]byte{feedId})
	if err != nil {
		env.t.Fatalf("removing feed %d: %v", feedIndex, err)
	}
	return env.waitMinedSuccessful(tx, "remove feed")
}

// AddFeed adds a feed whose 21-byte id has the given last byte (matching the
// id scheme of test_utils.Deploy) through the FastUpdatesConfiguration
// contract. The contract places it at the lowest previously-zeroed index, or
// appends it if there are no holes, and initializes its value from the mocked
// feed publisher. It returns the block the addition was mined in.
func (env *TestEnv) AddFeed(idByte byte) uint64 {
	configContract := env.configContract()

	feedId := [21]byte{}
	feedId[20] = idByte

	feedConfig := fast_updates_configuration.IFastUpdatesConfigurationFeedConfiguration{
		FeedId: feedId, RewardBandValue: 2000, InflationShare: big.NewInt(200),
	}
	tx, err := configContract.AddFeeds(env.transactOpts(),
		[]fast_updates_configuration.IFastUpdatesConfigurationFeedConfiguration{feedConfig})
	if err != nil {
		env.t.Fatalf("adding feed with id byte %d: %v", idByte, err)
	}
	return env.waitMinedSuccessful(tx, "add feed")
}

// FeedIdAt returns the last byte of the feed id currently configured at the
// given index (0 for a zeroed/removed slot).
func (env *TestEnv) FeedIdAt(feedIndex int) byte {
	feedIds, err := env.configContract().GetFeedIds(&bind.CallOpts{})
	if err != nil {
		env.t.Fatalf("getting feed ids: %v", err)
	}
	if feedIndex >= len(feedIds) {
		env.t.Fatalf("feed index %d out of range (%d feeds)", feedIndex, len(feedIds))
	}
	return feedIds[feedIndex][20]
}

func (env *TestEnv) configContract() *fast_updates_configuration.FastUpdatesConfiguration {
	configContract, err := fast_updates_configuration.NewFastUpdatesConfiguration(
		env.Contracts.FastUpdatesConfiguration, env.EthClient)
	if err != nil {
		env.t.Fatalf("binding configuration contract: %v", err)
	}
	return configContract
}

func (env *TestEnv) transactOpts() *bind.TransactOpts {
	privateKey, err := crypto.HexToECDSA(governanceKey[2:])
	if err != nil {
		env.t.Fatalf("parsing deployer key: %v", err)
	}
	opts, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(chainId))
	if err != nil {
		env.t.Fatalf("creating transactor: %v", err)
	}

	from := crypto.PubkeyToAddress(privateKey.PublicKey)
	nonce, err := env.EthClient.PendingNonceAt(context.Background(), from)
	if err != nil {
		env.t.Fatalf("getting nonce: %v", err)
	}
	gasPrice, err := env.EthClient.SuggestGasPrice(context.Background())
	if err != nil {
		env.t.Fatalf("getting gas price: %v", err)
	}
	opts.Nonce = big.NewInt(int64(nonce))
	opts.GasPrice = gasPrice
	opts.GasLimit = 8000000

	return opts
}
