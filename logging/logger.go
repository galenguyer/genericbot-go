package logging

import (
	"os"

	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var Logger = &logrus.Logger{
	Out: os.Stdout,
	Formatter: &logrus.TextFormatter{
		DisableLevelTruncation: true,
		PadLevelText:           true,
		FullTimestamp:          true,
	},
	Hooks: make(logrus.LevelHooks),
	Level: logrus.InfoLevel,
}

func init() {
	path := "./logs/genericbot.log"

	Logger.AddHook(lfshook.NewHook(lfshook.PathMap{
		logrus.InfoLevel:  path,
		logrus.WarnLevel:  path,
		logrus.ErrorLevel: path,
		logrus.FatalLevel: path,
		logrus.PanicLevel: path,
	}, &logrus.JSONFormatter{}))
}
