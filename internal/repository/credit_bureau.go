package repository

import (
	"context"

	"github.com/tgutu/credit-bureau-adapter/internal/database"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CreditBureauRepository interface {
	ListBureaus(ctx context.Context) ([]database.CreditBureau, error)
}

type CreditBureauRepositoryParams struct {
	fx.In
	DB     *gorm.DB
	Logger *zap.Logger
}

type creditBureauRepository struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewCreditBureauRepository(params CreditBureauRepositoryParams) CreditBureauRepository {
	return &creditBureauRepository{
		db:     params.DB,
		logger: params.Logger,
	}
}

func (r *creditBureauRepository) ListBureaus(ctx context.Context) ([]database.CreditBureau, error) {
	bureaus, err := gorm.G[database.CreditBureau](r.db).Find(ctx)
	if err != nil {
		r.logger.Error("failed to list bureaus", zap.Error(err))
		return nil, err
	}
	return bureaus, nil
}
