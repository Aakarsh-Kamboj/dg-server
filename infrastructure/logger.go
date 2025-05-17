package infrastructure

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger holds the Zap logger instance
var Logger *zap.Logger

// InitLogger initializes the logger based on environment
func InitLogger(env string) {
	var cfg zap.Config

	// Ensure logs/ directory exists
	logDir := "logs"
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		err := os.Mkdir(logDir, 0755)
		if err != nil {
			panic("Failed to create logs directory: " + err.Error())
		}
	}

	if env == "DEV" {
		cfg = zap.NewDevelopmentConfig() // Development config (human-readable logs)
		cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		cfg.OutputPaths = []string{"stdout", "logs/dev.log"} // Log to console + file
	} else {
		cfg = zap.NewProductionConfig()                       // Production config (JSON logs)
		cfg.OutputPaths = []string{"stdout", "logs/prod.log"} // Log to console + file
	}

	var err error
	Logger, err = cfg.Build()
	if err != nil {
		panic("Failed to initialize logger: " + err.Error())
	}

	// defer func(Logger *zap.Logger) {
	// 	err := Logger.Sync()
	// 	if err != nil {
	// 		panic("Failed to sync logger: " + err.Error())
	// 	}
	// }(Logger)
	Logger.Info("Logger initialized", zap.String("environment", env))
}

// GetLogger returns the logger instance
func GetLogger() *zap.Logger {
	if Logger == nil {
		InitLogger("DEV") // Default to development mode
	}
	return Logger
}

// ProvideLogger provides a logger instance for dependency injection
func ProvideLogger() *zap.Logger {
	return GetLogger()
}
