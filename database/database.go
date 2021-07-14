package database

import (
	"context"
	"time"

	"github.com/galenguyer/genericbot/config"
	"github.com/galenguyer/genericbot/entities"
	"github.com/galenguyer/genericbot/logging"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Database struct {
	client *mongo.Client
}

func Connect(config config.Config) Database {
	logging.Logger.WithFields(logrus.Fields{"module": "database", "method": "Connect"}).Info("beginning database connection")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.DatabaseConfig.Uri))
	if err != nil {
		logging.Logger.WithFields(logrus.Fields{"error": err, "module": "database", "method": "Connect"}).Fatal("error connecting to database")
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			logging.Logger.WithFields(logrus.Fields{"error": err, "module": "database", "method": "Connect"}).Fatal("disconnected from database")
		}
	}()

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		logging.Logger.WithFields(logrus.Fields{"error": err, "module": "database", "method": "Connect"}).Fatal("error pinging database")
	}

	logging.Logger.WithFields(logrus.Fields{"module": "database", "method": "Connect"}).Info("connected to mongodb")

	return Database{
		client: client,
	}
}

func (d Database) GetGuildConfig(guildId string) (*entities.GuildConfig, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var config entities.GuildConfig
	err := d.client.Database(guildId).Collection("guildConfig").FindOne(ctx, bson.D{}).Decode(&config)
	if err == mongo.ErrNoDocuments {
		config = entities.GuildConfig{
			Prefix:                       "",
			AdminRoleIds:                 []string{},
			ModRoleIds:                   []string{},
			UserRoleIds:                  []string{},
			RequiresRoles:                make(map[string][]string),
			MutedRoleId:                  "",
			MutedUsers:                   make(map[string]time.Time),
			AutoRoleIds:                  []string{},
			MessageLoggingChannelId:      "",
			UserLoggingChannelId:         "",
			MessageLoggingIgnoreChannels: []string{},
			VerifiedRoleId:               "",
			VerificationMessage:          "",
			JoinMessage:                  "",
			JoinMessageChannelId:         "",
			PointsEnabled:                false,
			TrustedRoleId:                "",
			TrustedRolePointsThreshold:   -1,
		}
	} else if err != nil {
		logging.Logger.WithFields(logrus.Fields{"error": err, "module": "database", "method": "GetGuildConfig"}).Error("error getting guildconfig")
	}
	return &config, nil
}
