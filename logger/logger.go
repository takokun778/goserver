package logger

import (
	"os"

	"go.uber.org/zap"
)

var Log *zap.Logger

func init() {
	logger, _ := zap.NewProduction()

	if os.Getenv("ENV") != "prod" {
		logger, _ = zap.NewDevelopment()
	}

	Log = logger
}
