package cba

import (
	"github.com/tgutu/credit-bureau-adapter/internal/repository"
	"github.com/tgutu/credit-bureau-adapter/pkg/pb/cba/v1"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type ServiceParams struct {
	fx.In
	Logger           *zap.Logger
	CreditBureauRepo repository.CreditBureauRepository
}

type server struct {
	cba.UnimplementedCreditBureauAdapterServiceServer
	logger           *zap.Logger
	creditBureauRepo repository.CreditBureauRepository
}

func NewServer(lc fx.Lifecycle, params ServiceParams) cba.CreditBureauAdapterServiceServer {
	return &server{
		logger:           params.Logger,
		creditBureauRepo: params.CreditBureauRepo,
	}
}
