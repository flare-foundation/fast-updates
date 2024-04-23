package config

import (
	"flag"
	"fmt"
	"math/big"
	"net/url"
	"os"

	"github.com/BurntSushi/toml"
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
	Level       string `toml:"level"` // valid values are: DEBUG, INFO, WARN, ERROR, DPANIC, PANIC, FATAL (zap)
	File        string `toml:"file"`
	MaxFileSize int    `toml:"max_file_size"` // In megabytes
	Console     bool   `toml:"console"`
}

type ChainConfig struct {
	NodeURL string `toml:"node_url"`
	ChainId int    `toml:"chain_id"`
}

type FastUpdateClientConfig struct {
	FastUpdaterAddress              string `toml:"fast_updater_address"`
	FastUpdatesConfigurationAddress string `toml:"fast_updates_configuration_address"`
	SubmissionAddress               string `toml:"submission_address"`
	IncentiveManagerAddress         string `toml:"incentive_manager_address"`
	FlareSystemManagerAddress       string `toml:"flare_system_manager"`
	MockAddress                     string `toml:"mock_address"`
	PrivateKey                      string `toml:"private_key" envconfig:"PRIVATE_KEY"`
	SortitionPrivateKey             string `toml:"sortition_private_key" envconfig:"SORTITION_PRIVATE_KEY"`
	AdvanceBlocks                   int    `toml:"advance_blocks"`
	SubmissionWindow                int    `toml:"submission_window"`
	MaxWeight                       int    `toml:"max_weight"`
	ValueProviderUrl                string `toml:"value_provider_base_url"`
}

type TransactionsConfig struct {
	Accounts           []string `toml:"accounts" envconfig:"ACCOUNTS"`
	GasLimit           int      `toml:"gas_limit"`
	Value              int      `toml:"value"`
	GasPriceMultiplier float64  `toml:"gas_price_multiplier"`
	ChainId            *big.Int
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

	_, err = url.ParseRequestURI(cfg.Client.ValueProviderUrl)
	if err != nil {
		return nil, errors.Wrap(err, "invalid URL specified for ValueProviderUrl: %w")
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
