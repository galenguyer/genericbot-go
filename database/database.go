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

var Client = Connect()

func Connect() *mongo.Client {
	logging.Logger.WithFields(logrus.Fields{"module": "database", "method": "Connect"}).Info("beginning database connection")

	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	// we can do this safely because we've already checked the config
	conf, _ := config.Load()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(conf.DatabaseConfig.Uri))
	if err != nil {
		logging.Logger.WithFields(logrus.Fields{"error": err, "module": "database", "method": "Connect"}).Fatal("error connecting to database")
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		logging.Logger.WithFields(logrus.Fields{"error": err, "module": "database", "method": "Connect"}).Fatal("error pinging database")
	}

	logging.Logger.WithFields(logrus.Fields{"module": "database", "method": "Connect"}).Info("connected to mongodb")

	return client
}

func Disconnect() {
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	if err := Client.Disconnect(ctx); err != nil {
		logging.Logger.WithFields(logrus.Fields{"error": err, "module": "database", "method": "Disconnect"}).Fatal("error disconnecting from database")
	}

	logging.Logger.WithFields(logrus.Fields{"module": "database", "method": "Disconnect"}).Info("disconnected from database")
}

func GetGuildConfig(guildId string) (*entities.GuildConfig, error) {
	logging.Logger.WithFields(logrus.Fields{"module": "database", "method": "Disconnect", "guild": guildId}).Info("getting guild config")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var config entities.GuildConfig
	err := Client.Database(guildId).Collection("guildConfig").FindOne(ctx, bson.D{}).Decode(&config)
	if err == mongo.ErrNoDocuments {
		config = entities.GuildConfig{
			GuildId:                        guildId,
			Prefix:                         "",
			AdminRoleIds:                   []string{},
			ModRoleIds:                     []string{},
			UserRoleIds:                    make(map[string][]string),
			RequiresRoles:                  make(map[string][]string),
			MutedRoleId:                    "",
			MutedUsers:                     make(map[string]time.Time),
			AutoRoleIds:                    []string{},
			MessageLoggingChannelId:        "",
			UserLoggingChannelId:           "",
			MessageLoggingIgnoreChannelIds: []string{},
			VerifiedRoleId:                 "",
			VerificationMessage:            "",
			JoinMessage:                    "",
			JoinMessageChannelId:           "",
			PointsEnabled:                  false,
			TrustedRoleId:                  "",
			TrustedRolePointsThreshold:     -1,
		}
	} else if err != nil {
		logging.Logger.WithFields(logrus.Fields{"error": err, "module": "database", "method": "GetGuildConfig"}).Error("error getting guildconfig")
	}
	return &config, nil
}

func SaveGuildConfig(guildId string, guildConfig entities.GuildConfig) error {
	logging.Logger.WithFields(logrus.Fields{"module": "database", "method": "SaveGuildConfig", "guild": guildId}).Info("saving guild config")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := Client.Database(guildId).Collection("config").FindOneAndReplace(ctx, bson.D{{Key: "_id", Value: guildId}}, guildConfig).Err()
	if err == mongo.ErrNoDocuments {
		_, err = Client.Database(guildId).Collection("config").InsertOne(ctx, guildConfig)
		return err
	}
	return err
}
