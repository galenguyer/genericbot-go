package database

import (
	"context"
	"strconv"
	"time"

	"github.com/galenguyer/genericbot/entities"
	"github.com/galenguyer/genericbot/logging"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetGuildConfig(guildId string) (*entities.GuildConfig, error) {
	logging.Logger.WithFields(logrus.Fields{"module": "database", "method": "Disconnect", "guild": guildId}).Info("getting guild config")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var config entities.GuildConfig
	findOptions := options.FindOne()
	findOptions.SetSort(bson.D{{Key: "_id", Value: -1}})
	err := Client.Database(guildId).Collection("config").FindOne(ctx, bson.D{}, findOptions).Decode(&config)

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

func SaveGuildConfig(guildId string, messageId string, guildConfig entities.GuildConfig) error {
	logging.Logger.WithFields(logrus.Fields{"module": "database", "method": "SaveGuildConfig", "guild": guildId}).Info("saving guild config")

	// set up context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// update id field for guildconfig
	guildConfig.MessageId, _ = strconv.ParseUint(messageId, 10, 64)

	// insert document
	_, err := Client.Database(guildId).Collection("config").InsertOne(ctx, guildConfig)
	return err
}
