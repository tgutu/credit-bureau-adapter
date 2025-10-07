package database

import (
	"context"
	"fmt"

	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DatabaseParams struct {
	fx.In
	Logger *zap.Logger
}

func NewDatabase(lc fx.Lifecycle, params DatabaseParams) (*gorm.DB, error) {

	logger := logger.New(
		zap.NewStdLog(params.Logger),
		logger.Config{
			SlowThreshold:             0,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		},
	)

	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{
		Logger:         logger,
		PrepareStmt:    true,
		TranslateError: true,
	})

	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("database not reachable: %w", err)
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			params.Logger.Info("Connected to the database")

			if err := db.AutoMigrate(&CreditBureau{}); err != nil {
				return fmt.Errorf("failed to migrate database: %w", err)
			}
			return nil
		},
		OnStop: func(ctx context.Context) error {
			params.Logger.Info("Closing database connection")
			return sqlDB.Close()
		},
	})

	return db, nil
}
