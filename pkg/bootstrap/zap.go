package bootstrap

import (
	"log"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func MustConfigureLogger() *zap.Logger {
	var cfg zap.Config
	switch os.Getenv("APP_ENV") {
	case "prod", "production":
		cfg = zap.NewDevelopmentConfig()
		cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	case "dev", "development":
		cfg = zap.NewProductionConfig()
	default:
		log.Fatalf("unknown APP_ENV: %s\n", os.Getenv("APP_ENV"))
	}

	var err error
	if cfg.Level, err = zap.ParseAtomicLevel(os.Getenv("LOG_LEVEL")); err != nil {
		log.Fatalf("cannot parse LOG_LEVEL: %s: %v\n", os.Getenv("LOG_LEVEL"), err)
	}

	return zap.Must(cfg.Build())
}
