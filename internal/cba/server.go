package cba

import (
	"context"

	"github.com/tgutu/credit-bureau-adapter/internal/cba/adapter"
	"github.com/tgutu/credit-bureau-adapter/internal/repository"
	"github.com/tgutu/credit-bureau-adapter/pkg/pb/cba/v1"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ServiceParams struct {
	fx.In
	Logger            *zap.Logger
	CreditBureauRepo  repository.CreditBureauRepository
	EquifaxAdapter    *adapter.EquifaxAdapter
	TransUnionAdapter *adapter.TransUnionAdapter
}

type server struct {
	cba.UnimplementedCreditBureauAdapterServiceServer
	logger            *zap.Logger
	creditBureauRepo  repository.CreditBureauRepository
	equifaxAdapter    *adapter.EquifaxAdapter
	transUnionAdapter *adapter.TransUnionAdapter
}

func NewServer(lc fx.Lifecycle, params ServiceParams) cba.CreditBureauAdapterServiceServer {
	return &server{
		logger:            params.Logger,
		creditBureauRepo:  params.CreditBureauRepo,
		equifaxAdapter:    params.EquifaxAdapter,
		transUnionAdapter: params.TransUnionAdapter,
	}
}

func (s *server) GetBureaus(ctx context.Context, in *cba.GetBureausRequest) (*cba.GetBureausResponse, error) {
	bureaus, err := s.creditBureauRepo.ListBureaus(ctx)
	if err != nil {
		s.logger.Error("failed to get bureaus", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "failed to get bureaus: %v", err)
	}

	var pbBureaus []*cba.Bureau
	for _, b := range bureaus {
		pbBureaus = append(pbBureaus, &cba.Bureau{
			Name: b.Name,
		})
	}

	return &cba.GetBureausResponse{Bureaus: pbBureaus}, nil
}
