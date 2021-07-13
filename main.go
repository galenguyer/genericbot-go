package main

import (
	"github.com/galenguyer/genericbot/config"
	"github.com/galenguyer/genericbot/logging"
	"github.com/sirupsen/logrus"
)

func main() {
	config, err := config.Load()
	if err != nil {
		logging.Logger.WithFields(logrus.Fields{"error": err, "module": "genericbot", "method": "main"}).Fatal("fatal error loading configuration")
	}
	if err = config.Validate(); err != nil {
		logging.Logger.WithFields(logrus.Fields{"error": err, "module": "genericbot", "method": "main"}).Fatal("fatal error loading configuration")
	}
	logging.Logger.WithFields(logrus.Fields{"module": "genericbot", "method": "main"}).Info("starting genericbot")
	logging.Logger.WithFields(logrus.Fields{"module": "genericbot", "method": "main"}).Info("using prefix " + config.BotConfig.Prefix)
	logging.Logger.WithFields(logrus.Fields{"module": "genericbot", "method": "main"}).Info("using log level " + config.BotConfig.LogLevel.String())
}
