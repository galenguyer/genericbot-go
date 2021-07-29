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

func GetCustomCommands(guildId string) ([]entities.CustomCommand, error) {
	logging.Logger.WithFields(logrus.Fields{"module": "database", "method": "Disconnect", "guild": guildId}).Info("getting commands")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var commands []entities.CustomCommand

	cursor, err := Client.Database(guildId).Collection("commands").Find(ctx, bson.D{})
	if err != nil {
		logging.Logger.WithFields(logrus.Fields{"error": err, "module": "database", "method": "GetCustomCommands"}).Error("error getting commands")
		return nil, err
	}
	for cursor.Next(ctx) {
		var current entities.CustomCommand
		err = cursor.Decode(&current)
		if err != nil {
			return nil, err
		}
		commands = append(commands, current)
	}
	return commands, nil
}

func SaveCustomCommand(guildId string, command entities.CustomCommand) (UpsertResult, error) {
	logging.Logger.WithFields(logrus.Fields{"module": "database", "method": "SaveCustomCommand", "guild": guildId, "command": command.Name}).Info("saving customCommand")

	// set up context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// insert document
	err := Client.Database(guildId).Collection("commands").FindOneAndReplace(ctx, bson.D{{Key: "_id", Value: command.Name}}, command).Err()
	if err == mongo.ErrNoDocuments {
		_, err = Client.Database(guildId).Collection("commands").InsertOne(ctx, command)
		return New, err
	}
	return Updated, err
}
