package database

import (
	"context"
	"time"

	"github.com/galenguyer/genericbot/config"
	"github.com/galenguyer/genericbot/logging"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type UpsertResult int

const (
	New     UpsertResult = 0
	Updated UpsertResult = 1
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
