package bootstrap

import (
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var ZapLogger = fx.WithLogger(func(decorable *zap.Logger) fxevent.Logger {
	logger := &fxevent.ZapLogger{Logger: decorable}
	logger.UseLogLevel(zapcore.DebugLevel)

	return logger
})
