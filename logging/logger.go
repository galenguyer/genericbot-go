package logging

import (
	"os"
	"runtime"

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

func Trace() runtime.Frame {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	return frame
}
