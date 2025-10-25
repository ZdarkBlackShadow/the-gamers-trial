package config

import (
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

var Log = logrus.New()

func InitLogger() {
	levelStr := strings.ToLower(os.Getenv("LOG_LEVEL"))
	if levelStr == "" {
		levelStr = "info"
	}

	level, err := logrus.ParseLevel(levelStr)
	if err != nil {
		level = logrus.InfoLevel
	}

	Log.SetLevel(level)
	Log.SetOutput(os.Stdout)
	Log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
		ForceColors:     true,
		PadLevelText:    true,
	})

	Log.Infof("Logger initialized with level: %s ✅", level)
}
