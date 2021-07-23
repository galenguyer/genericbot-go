package database

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/galenguyer/genericbot/entities"
	"github.com/galenguyer/genericbot/logging"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func ConvertLegacyGuildConfig(guildId string) (*entities.GuildConfig, error) {
	logging.Logger.WithFields(logrus.Fields{"module": "database", "method": "ConvertLegacyGuildConfig", "guild": guildId}).Info("getting guild config")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var legConf struct {
		Prefix                       string
		AdminRoleIds                 []uint64
		ModRoleIds                   []uint64
		UserRoleIds                  [][]interface{} `bson:"UserRoles"`
		MutedRoleId                  uint64
		MutedUsers                   []uint64
		AutoRoleIds                  []uint64
		LoggingChannelid             uint64
		MessageLoggingIgnoreChannels []uint64
		VerifiedRole                 uint64
		VerifiedMessage              string
		JoinMessage                  string
		JoinMessageChannelId         uint64
		PointsEnabled                bool
		TrustedRoleId                uint64
		TrustedRolePointsThreshold   int
	}

	err := Client.Database(guildId).Collection("config").FindOne(ctx, bson.D{}).Decode(&legConf)
	if err == mongo.ErrNoDocuments {
		logging.Logger.WithFields(logrus.Fields{"module": "database", "method": "ConvertLegacyGuildConfig", "guild": guildId}).Info("no guild config found")
		return &entities.GuildConfig{
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
		}, nil
	} else if err != nil {
		logging.Logger.WithFields(logrus.Fields{"error": err, "module": "database", "method": "ConvertLegacyGuildConfig"}).Error("error getting guildconfig")
	}

	var newAdminRoleIds []string
	for _, role := range legConf.AdminRoleIds {
		newAdminRoleIds = append(newAdminRoleIds, fmt.Sprint(role))
	}
	var newModRoleIds []string
	for _, role := range legConf.ModRoleIds {
		newModRoleIds = append(newModRoleIds, fmt.Sprint(role))
	}
	var newUserRoleIds = make(map[string][]string)
	for _, arr := range legConf.UserRoleIds {
		switch reflect.TypeOf(arr[1]).Kind() {
		case reflect.Slice:
			key := fmt.Sprint(arr[0])
			if key == "" {
				key = "Ungrouped"
			}
			newUserRoleIds[key] = make([]string, reflect.ValueOf(arr[1]).Len())
			for i := 0; i < reflect.ValueOf(arr[1]).Len(); i++ {
				newUserRoleIds[key][i] = fmt.Sprint(reflect.ValueOf(arr[1]).Index(i))
			}
		}
	}
	var newRequiresRoles = make(map[string][]string)
	var newAutoRoleIds []string
	for _, role := range legConf.AutoRoleIds {
		newAutoRoleIds = append(newAutoRoleIds, fmt.Sprint(role))
	}
	var newMessageLoggingIgnoreChannelIds []string
	for _, channel := range legConf.MessageLoggingIgnoreChannels {
		newMessageLoggingIgnoreChannelIds = append(newMessageLoggingIgnoreChannelIds, fmt.Sprint(channel))
	}

	return &entities.GuildConfig{
		GuildId:                        guildId,
		Prefix:                         legConf.Prefix,
		AdminRoleIds:                   newAdminRoleIds,
		ModRoleIds:                     newModRoleIds,
		UserRoleIds:                    newUserRoleIds,
		RequiresRoles:                  newRequiresRoles,
		MutedRoleId:                    fmt.Sprint(legConf.MutedRoleId),
		MutedUsers:                     make(map[string]time.Time),
		AutoRoleIds:                    newAutoRoleIds,
		MessageLoggingChannelId:        fmt.Sprint(legConf.LoggingChannelid),
		MessageLoggingIgnoreChannelIds: newMessageLoggingIgnoreChannelIds,
		VerifiedRoleId:                 fmt.Sprint(legConf.VerifiedRole),
		VerificationMessage:            legConf.VerifiedMessage,
		JoinMessage:                    legConf.JoinMessage,
		JoinMessageChannelId:           fmt.Sprint(legConf.JoinMessageChannelId),
		PointsEnabled:                  legConf.PointsEnabled,
		TrustedRoleId:                  fmt.Sprint(legConf.TrustedRoleId),
		TrustedRolePointsThreshold:     legConf.TrustedRolePointsThreshold,
	}, nil
}