package config

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm/logger"
)

type GormLogger struct {
	LogLevel logger.LogLevel
}

func (l GormLogger) LogMode(level logger.LogLevel) logger.Interface {
	l.LogLevel = level
	return l
}

func (l GormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if Log.Level >= logrus.InfoLevel {
		Log.Infof(msg, data...)
	}
}

func (l GormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if Log.Level >= logrus.WarnLevel {
		Log.Warnf(msg, data...)
	}
}

func (l GormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if Log.Level >= logrus.ErrorLevel {
		Log.Errorf(msg, data...)
	}
}

func (l GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	sql, rows := fc()

	entry := Log.WithFields(logrus.Fields{
		"duration": elapsed,
		"rows":     rows,
	})

	if err != nil {
		entry.WithError(err).Error(sql)
	} else if elapsed > 200*time.Millisecond {
		entry.Warn(sql)
	} else if Log.Level >= logrus.DebugLevel {
		entry.Debug(sql)
	}
}
