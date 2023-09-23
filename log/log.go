package log

import (
	"context"
	"github.com/pkg/errors"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewLogger creates
func NewLogger(lc fx.Lifecycle) *zap.Logger {
	logConfig := zap.NewProductionConfig()
	logConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// Disable stacktrace field on logs as github.com/pkg/errors
	// adds the stacktrace to errorVerbose field already.
	logConfig.DisableStacktrace = true

	var err error
	log, err := logConfig.Build()
	if err != nil {
		panic(errors.Wrap(err, "failed to initialize zap logger"))
	}

	// Sets the created log instance as global logger
	zap.ReplaceGlobals(log)

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return log.Sync()
		},
	})

	return log
}
