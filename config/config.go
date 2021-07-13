package config

import (
	"errors"
	"os"

	"github.com/galenguyer/genericbot/logging"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

var (
	ErrNoTokenSet     = errors.New("no token set in config.yaml or BOT_TOKEN")
	ErrUnchangedToken = errors.New("token is still CHANGE_ME in config.yaml")
)

type Config struct {
	BotConfig struct {
		Token    string       `yaml:"token" envconfig:"BOT_TOKEN"`
		Prefix   string       `yaml:"prefix" envconfig:"BOT_PREFIX" default:">"`
		LogLevel logrus.Level `yaml:"log-level" envconfig:"BOT_LOG_LEVEL"`
	} `yaml:"bot"`
}

// Load is called to return a Config struct constructed from the
// config file first and the environment as overrides
func Load() (*Config, error) {
	var config Config
	err := loadFile(&config)
	if err != nil {
		logging.Logger.WithFields(logrus.Fields{"error": err, "module": "config", "method": "Load"}).Error("error loading configuration file")
		return nil, err
	}
	err = loadEnv(&config)
	if err != nil {
		logging.Logger.WithFields(logrus.Fields{"error": err, "module": "config", "method": "Load"}).Error("error loading configuration from environment")
		return nil, err
	}
	if err = config.Validate(); err != nil {
		logging.Logger.WithFields(logrus.Fields{"error": err, "module": "config", "method": "Load"}).Error("error validating configuration")
		return nil, err
	}
	return &config, nil
}

func (config *Config) Validate() error {
	if config.BotConfig.Token == "" {
		return ErrNoTokenSet
	}
	if config.BotConfig.Token == "CHANGE_ME" {
		return ErrUnchangedToken
	}
	return nil
}

func loadFile(c *Config) error {
	f, err := os.Open("config.yaml")
	if err != nil {
		logging.Logger.WithFields(logrus.Fields{"error": err, "module": "config", "method": "loadFile"}).Error("error opening config file")
		return err
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(c)
	if err != nil {
		logging.Logger.WithFields(logrus.Fields{"error": err, "module": "config", "method": "loadFile"}).Error("error decoding config.yaml")
		return err
	}
	return nil
}

func loadEnv(c *Config) error {
	err := envconfig.Process("", c)
	if err != nil {
		logging.Logger.WithFields(logrus.Fields{"error": err, "module": "config", "method": "loadEnv"}).Error("error processing environment variables")
		return err
	}
	return nil
}
