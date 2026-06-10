package config

import (
	"flag"
	"fmt"
	"math/big"
	"net/url"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

var (
	CallTimeoutMillisDefault int                          = 2000   // todo
	TxTimeoutMillisDefault   int                          = 120000 // todo
	GlobalConfigCallback     ConfigCallback[GlobalConfig] = ConfigCallback[GlobalConfig]{}
	CfgFlag                                               = flag.String("config", "config.toml", "Configuration file (toml format)")
)

type GlobalConfig interface {
	LoggerConfig() LoggerConfig
}

type Config struct {
	Logger       LoggerConfig           `toml:"logger"`
	Chain        ChainConfig            `toml:"chain"`
	Client       FastUpdateClientConfig `toml:"client"`
	Transactions TransactionsConfig     `toml:"transactions"`
}

type LoggerConfig struct {
	Level         string  `toml:"level"` // valid values are: DEBUG, INFO, WARN, ERROR, DPANIC, PANIC, FATAL (zap)
	File          string  `toml:"file"`
	MaxFileSize   int     `toml:"max_file_size"` // In megabytes
	Console       bool    `toml:"console"`
	MinBalance    float64 `toml:"min_balance"`
	FeedValuesLog int     `toml:"feed_values_log"`
}

type ChainConfig struct {
	NodeURL string `toml:"node_url" envconfig:"NODE_URL"`
	ChainId int    `toml:"chain_id"`
	ApiKey  string `toml:"api_key" envconfig:"API_KEY"`
}

type FastUpdateClientConfig struct {
	FastUpdaterAddress              string `toml:"fast_updater_address"`
	FastUpdatesConfigurationAddress string `toml:"fast_updates_configuration_address"`
	SubmissionAddress               string `toml:"submission_address"`
	IncentiveManagerAddress         string `toml:"incentive_manager_address"`
	FlareSystemManagerAddress       string `toml:"flare_system_manager"`
	MockAddress                     string `toml:"mock_address"`
	SigningPrivateKey               string `toml:"signing_private_key" envconfig:"SIGNING_PRIVATE_KEY"`
	SortitionPrivateKey             string `toml:"sortition_private_key" envconfig:"SORTITION_PRIVATE_KEY"`
	SubmissionWindow                int    `toml:"submission_window"`
	MaxWeight                       int    `toml:"max_weight"`
	ValueProviderUrl                string `toml:"value_provider_url"`
	FetchCurrentFeedsValue          string `toml:"fetch_current_feeds_value"`
	FetchCurrentFeedsAddress        string `toml:"fetch_current_feeds_address"`
}

type TransactionsConfig struct {
	Accounts []string `toml:"accounts" envconfig:"ACCOUNTS"`
	GasLimit int      `toml:"gas_limit"`
	Value    int      `toml:"value"`

	// Submissions are always EIP-1559 (type 2) transactions: a type-2 fee cap stays
	// valid as the base fee rises, so a single submission stays includable through a
	// base-fee spike within the tight submission window. (A legacy type-0 tx commits
	// to one fixed gas price; once the base fee climbs past it the tx is no longer
	// includable and misses the window.) Note submitAndPass is a flare-node
	// "prioritised contract call": the node refunds everything above a fixed nominal
	// fee, so a generous cap costs nothing extra and the values below are tuned for
	// inclusion robustness, not cost. The fee is shaped as:
	//   gasTipCap = clamp(baseFee * MaxPriorityFeeMultiplier, MinimalMaxPriorityFee, MaximalMaxPriorityFee)
	//   gasFeeCap = baseFee * BaseFeeMultiplier + gasTipCap
	// Defaults are applied when unset/<=0 (matching flare-system-client):
	//   BaseFeeMultiplier=4, MaxPriorityFeeMultiplier=2,
	//   MinimalMaxPriorityFee=100 Gwei, MaximalMaxPriorityFee=5000 Gwei.
	BaseFeeMultiplier        int64 `toml:"base_fee_multiplier"`
	MaxPriorityFeeMultiplier int64 `toml:"max_priority_fee_multiplier"`
	MinimalMaxPriorityFee    int64 `toml:"minimal_max_priority_fee"` // in wei
	MaximalMaxPriorityFee    int64 `toml:"maximal_max_priority_fee"` // in wei

	ChainId *big.Int
}

func newConfig() *Config {
	return &Config{}
}

func BuildConfig() (*Config, error) {
	cfgFileName := *CfgFlag

	cfg := newConfig()
	err := ParseConfigFile(cfg, cfgFileName)
	if err != nil {
		return nil, err
	}
	err = ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	err = validateConfig(cfg)
	if err != nil {
		return nil, errors.Wrap(err, "config validation failed")
	}

	return cfg, nil
}

func ParseConfigFile(cfg *Config, fileName string) error {
	content, err := os.ReadFile(fileName)
	if err != nil {
		return fmt.Errorf("error opening config file: %w", err)
	}

	_, err = toml.Decode(string(content), cfg)
	if err != nil {
		return fmt.Errorf("error parsing config file: %w", err)
	}
	return nil
}

func ReadEnv(cfg *Config) error {
	err := envconfig.Process("", cfg)
	if err != nil {
		return fmt.Errorf("error reading env config: %w", err)
	}

	return nil
}

func (c Config) LoggerConfig() LoggerConfig {
	return c.Logger
}

func validateConfig(cfg *Config) error {
	_, err := url.ParseRequestURI(cfg.Client.ValueProviderUrl)
	if err != nil {
		return errors.Wrap(err, "invalid URL specified for ValueProviderUrl: %w")
	}

	if len(cfg.Transactions.Accounts) == 0 {
		return errors.New("no submission accounts provided")
	}

	return nil
}

// Dial the chain node and return an ethclient.Client.
func (chain *ChainConfig) DialETH() (*ethclient.Client, error) {
	rpcURL, err := chain.getRPCURL()
	if err != nil {
		return nil, err
	}

	return ethclient.Dial(rpcURL)
}

// Get the full RPC URL which may be passed to ethclient.Dial. Includes API key
// as query param if it is configured.
func (chain *ChainConfig) getRPCURL() (string, error) {
	u, err := url.Parse(chain.NodeURL)
	if err != nil {
		return "", err
	}

	if chain.ApiKey == "" {
		return u.String(), nil
	}

	q := u.Query()
	q.Set("x-apikey", chain.ApiKey)
	u.RawQuery = q.Encode()

	return u.String(), nil
}
