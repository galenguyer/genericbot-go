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

func GetUser(guildId string, userId string) (*entities.User, error) {
	logging.Logger.WithFields(logrus.Fields{"module": "database", "method": "GetUser", "guild": guildId, "user": userId}).Info("getting user")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user entities.User
	err := Client.Database(guildId).Collection("users").FindOne(ctx, bson.D{{Key: "_id", Value: userId}}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	} else if err != nil {
		logging.Logger.WithFields(logrus.Fields{
			"error":  err,
			"module": "database",
			"method": "GetUser",
			"guild":  guildId,
			"user":   userId,
		}).Error("error searching for user")
		return nil, err
	}
	return &user, nil
}

func SaveUser(guildId string, user entities.User) (UpsertResult, error) {
	logging.Logger.WithFields(logrus.Fields{"module": "database", "method": "SaveUser", "guild": guildId, "user": user.Id}).Info("saving user")

	// set up context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// insert document
	err := Client.Database(guildId).Collection("users").FindOneAndReplace(ctx, bson.D{{Key: "_id", Value: user.Id}}, user).Err()
	if err == mongo.ErrNoDocuments {
		_, err = Client.Database(guildId).Collection("users").InsertOne(ctx, user)
		return New, err
	}
	return Updated, err
}
