package repository

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CreditBureauRepository interface {
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
