package cba

import (
	"context"

	"github.com/tgutu/credit-bureau-adapter/internal/apicode"
	"github.com/tgutu/credit-bureau-adapter/internal/cba/adapter"
	"github.com/tgutu/credit-bureau-adapter/internal/database"
	"github.com/tgutu/credit-bureau-adapter/pkg/pb/cba/v1"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type ServiceParams struct {
	fx.In
	Logger            *zap.Logger
	CreditBureauRepo  database.CreditBureauRepository
	ExperianAdapter   *adapter.ExperianAdapter
	EquifaxAdapter    *adapter.EquifaxAdapter
	TransUnionAdapter *adapter.TransUnionAdapter
}

type server struct {
	cba.UnimplementedCreditBureauAdapterServiceServer
	logger            *zap.Logger
	creditBureauRepo  database.CreditBureauRepository
	experianAdapter   *adapter.ExperianAdapter
	equifaxAdapter    *adapter.EquifaxAdapter
	transUnionAdapter *adapter.TransUnionAdapter
}

func NewServer(lc fx.Lifecycle, params ServiceParams) cba.CreditBureauAdapterServiceServer {
	return &server{
		logger:            params.Logger,
		creditBureauRepo:  params.CreditBureauRepo,
		experianAdapter:   params.ExperianAdapter,
		equifaxAdapter:    params.EquifaxAdapter,
		transUnionAdapter: params.TransUnionAdapter,
	}
}

func (s *server) GetBureaus(ctx context.Context, in *cba.GetBureausRequest) (*cba.GetBureausResponse, error) {
	bureaus, err := s.creditBureauRepo.ListBureaus(ctx)
	if err != nil {
		s.logger.Error("failed to get bureaus", zap.Error(err))
		return nil, apicode.ErrCreditRepoListBureausFailed
	}

	var pbBureaus []*cba.Bureau
	for _, b := range bureaus {
		pbBureaus = append(pbBureaus, &cba.Bureau{
			Name: b.Name,
		})
	}

	return &cba.GetBureausResponse{Bureaus: pbBureaus}, nil
}

func (s *server) GetCreditReport(ctx context.Context, in *cba.GetCreditReportRequest) (*cba.GetCreditReportResponse, error) {
	bureau, err := s.creditBureauRepo.GetBureauByName(ctx, in.BureauName)
	if err != nil {
		s.logger.Error("failed to get bureau by name", zap.String("name", in.BureauName), zap.Error(err))
		return nil, apicode.ErrCreditRepoGetBureauByNameFailed
	}

	adapter, err := s.getAdapterByName(bureau.Name)
	if err != nil {
		return nil, err
	}

	report, err := adapter.GetCreditReport(ctx, in)
	if err != nil {
		s.logger.Error("failed to get credit report", zap.String("bureau", in.BureauName), zap.Error(err))
		return nil, err
	}

	return &report, nil
}

func (s *server) GetCreditScore(ctx context.Context, in *cba.GetCreditScoreRequest) (*cba.GetCreditScoreResponse, error) {
	bureau, err := s.creditBureauRepo.GetBureauByName(ctx, in.BureauName)
	if err != nil {
		s.logger.Error("failed to get bureau by name", zap.String("name", in.BureauName), zap.Error(err))
		return nil, apicode.ErrCreditRepoGetBureauByNameFailed
	}

	adapter, err := s.getAdapterByName(bureau.Name)
	if err != nil {
		return nil, err
	}

	report, err := adapter.GetCreditScore(ctx, in)
	if err != nil {
		s.logger.Error("failed to get credit score", zap.String("bureau", in.BureauName), zap.Error(err))
		return nil, err
	}

	return &report, nil
}

func (s *server) getAdapterByName(name string) (CreditBureauAdapter, error) {
	switch name {
	case "experian":
		return s.experianAdapter, nil
	case "equifax":
		return s.equifaxAdapter, nil
	case "transunion":
		return s.transUnionAdapter, nil
	default:
		s.logger.Error("unsupported bureau", zap.String("name", name))
		return nil, apicode.ErrCbaUnsupportedBureau
	}
}
