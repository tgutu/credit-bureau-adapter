package cba

import (
	"github.com/tgutu/credit-bureau-adapter/pkg/pb/cba/v1"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type ServiceParams struct {
	fx.In
	Logger *zap.Logger
}

type server struct {
	cba.UnimplementedCreditBureauAdapterServiceServer
	logger *zap.Logger
}

func NewServer(lc fx.Lifecycle, params ServiceParams) cba.CreditBureauAdapterServiceServer {
	return &server{
		logger: params.Logger,
	}
}
