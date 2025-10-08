package repository

import (
	"context"

	"github.com/tgutu/credit-bureau-adapter/internal/apicode"
	"github.com/tgutu/credit-bureau-adapter/internal/database"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CreditBureauRepository interface {
	GetBureauByName(ctx context.Context, name string) (*database.CreditBureau, error)
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
		return nil, apicode.ErrCreditRepoListBureausFailed
	}
	return bureaus, nil
}

func (r *creditBureauRepository) GetBureauByName(ctx context.Context, name string) (*database.CreditBureau, error) {
	bureau, err := gorm.G[database.CreditBureau](r.db).Where("name = ?", name).First(ctx)
	if err != nil {
		r.logger.Error("failed to get bureau by name", zap.String("name", name), zap.Error(err))
		return nil, apicode.ErrCreditRepoGetBureauByNameFailed
	}
	return &bureau, nil
}
