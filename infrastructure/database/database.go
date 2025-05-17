package database

import (
	"fmt"
	"go.uber.org/zap"

	"dg-server/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Database holds the DB connection instance
type Database struct {
	DB *gorm.DB
}

// NewDatabase initializes the PostgreSQL database connection
func NewDatabase(cfg *config.Config, logger *zap.Logger) (*Database, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		cfg.Database.Host, cfg.Database.User, cfg.Database.Password,
		cfg.Database.Name, cfg.Database.Port, cfg.Database.SSLMode,
		cfg.Database.TimeZone,
	)

	// Use a GORM config that routes SQL logs through Zap
	gormCfg := &gorm.Config{
		Logger: NewZapGormLogger(logger),
	}
	db, err := gorm.Open(postgres.Open(dsn), gormCfg)
	if err != nil {
		logger.Error("failed to open DB", zap.Error(err))
		return nil, fmt.Errorf("opening DB: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		logger.Error("cannot get sql.DB", zap.Error(err))
		return nil, fmt.Errorf("getting sql.DB: %w", err)
	}

	sqlDB.SetMaxOpenConns(cfg.Database.MaxOpenConns)
	sqlDB.SetMaxIdleConns(cfg.Database.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(cfg.Database.ConnMaxLifetime)
	sqlDB.SetConnMaxIdleTime(cfg.Database.ConnMaxIdleTime)

	logger.Info("database connection established")

	if err := dbMigrate(db, logger); err != nil {
		return nil, fmt.Errorf("migrating schema: %w", err)
	}

	return &Database{DB: db}, nil
}
