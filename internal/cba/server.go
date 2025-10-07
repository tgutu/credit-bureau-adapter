package cba

import (
	"context"

	"github.com/tgutu/credit-bureau-adapter/internal/repository"
	"github.com/tgutu/credit-bureau-adapter/pkg/pb/cba/v1"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
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

func (s *server) GetBureaus(ctx context.Context, in *emptypb.Empty) (*cba.GetBureausResponse, error) {
	bureaus, err := s.creditBureauRepo.ListBureaus(ctx)
	if err != nil {
		s.logger.Error("failed to get bureaus", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "failed to get bureaus: %v", err)
	}

	var pbBureaus []*cba.GetBureausResponse_Bureau
	for _, b := range bureaus {
		pbBureaus = append(pbBureaus, &cba.GetBureausResponse_Bureau{
			Id:   b.ID,
			Name: b.Name,
		})
	}

	return &cba.GetBureausResponse{Bureaus: pbBureaus}, nil
}
