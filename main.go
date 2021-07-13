package main

import (
	"github.com/galenguyer/genericbot/logging"
	"github.com/sirupsen/logrus"
)

func main() {
	logging.Logger.WithFields(logrus.Fields{"module": "genericbot", "method": "main"}).Info("starting genericbot")
}
