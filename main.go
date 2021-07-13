package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/galenguyer/genericbot/config"
	"github.com/galenguyer/genericbot/logging"
	"github.com/servusdei2018/shards"
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

	manager, err := shards.New("Bot " + config.BotConfig.Token)
	if err != nil {
		logging.Logger.WithFields(logrus.Fields{"error": err, "module": "genericbot", "method": "main"}).Error("error creating shard manager")
	}

	manager.AddHandler(onShardConnect)

	err = manager.Start()
	if err != nil {
		logging.Logger.WithFields(logrus.Fields{"error": err, "module": "genericbot", "method": "main"}).Error("error starting shard manager")
	}

	// wait for ^C to exit
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	logging.Logger.WithFields(logrus.Fields{"module": "genericbot", "method": "main"}).Infof("stopping shard manager with %d shards", manager.ShardCount)
	manager.Shutdown()
	logging.Logger.WithFields(logrus.Fields{"module": "genericbot", "method": "main"}).Info("stopped shard manager")
}

func onShardConnect(s *discordgo.Session, evt *discordgo.Connect) {
	logging.Logger.WithFields(logrus.Fields{"module": "genericbot", "method": "onShardConnect", "shard": s.ShardID}).Info("shard connected")
}
