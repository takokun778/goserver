package logger

import (
	"context"
	"math/rand"
	"os"
	"time"

	"github.com/oklog/ulid"
	"go.uber.org/zap"
)

type key string

const (
	ctxLog = key("log")
)

func SetLogCtx(parent context.Context) context.Context {
	t := time.Now()

	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)

	id := ulid.MustNew(ulid.Timestamp(t), entropy)

	logger, _ := zap.NewProduction()

	if os.Getenv("ENV") != "prod" {
		logger, _ = zap.NewDevelopment()
	}

	log := logger.With(zap.String("rid", id.String()))

	return context.WithValue(parent, ctxLog, log)
}

func GetLogCtx(ctx context.Context) *zap.Logger {
	v := ctx.Value(ctxLog)

	log, ok := v.(*zap.Logger)

	if !ok {
		logger, _ := zap.NewProduction()
		return logger
	}

	return log
}
