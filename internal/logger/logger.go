package logger

import (
	"go.uber.org/zap"
)

var Log *zap.Logger

// InitLogger initializes the global logger
func InitLogger() {
	var err error
	Log, err = zap.NewProduction()
	if err != nil {
		panic(err)
	}
}

// Sync flushes any buffered log entries
func Sync() {
	_ = Log.Sync()
}
