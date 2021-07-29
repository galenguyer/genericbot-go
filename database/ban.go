package database

import (
	"context"
	"time"

	"github.com/galenguyer/genericbot/entities"
	"github.com/galenguyer/genericbot/logging"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetBans(guildId string) ([]entities.Ban, error) {
	logging.Logger.WithFields(logrus.Fields{"module": "database", "method": "Disconnect", "guild": guildId}).Info("getting bans")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var bans []entities.Ban

	cursor, err := Client.Database(guildId).Collection("bans").Find(ctx, bson.D{})
	if err != nil {
		logging.Logger.WithFields(logrus.Fields{"error": err, "module": "database", "method": "GetBans"}).Error("error getting bans")
		return nil, err
	}
	for cursor.Next(ctx) {
		var current entities.Ban
		err = cursor.Decode(&current)
		if err != nil {
			return nil, err
		}
		bans = append(bans, current)
	}
	return bans, nil
}

func SaveBan(guildId string, ban entities.Ban) (UpsertResult, error) {
	logging.Logger.WithFields(logrus.Fields{"module": "database", "method": "SaveBan", "guild": guildId, "user": ban.UserId}).Info("saving ban")

	// set up context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// insert document
	err := Client.Database(guildId).Collection("bans").FindOneAndReplace(ctx, bson.D{{Key: "_id", Value: ban.UserId}}, ban).Err()
	if err == mongo.ErrNoDocuments {
		_, err = Client.Database(guildId).Collection("bans").InsertOne(ctx, ban)
		return New, err
	}
	return Updated, err
}
