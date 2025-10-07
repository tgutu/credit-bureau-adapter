package server

import (
	"context"
	"fmt"
	"net"

	"github.com/tgutu/credit-bureau-adapter/internal/config"
	"github.com/tgutu/credit-bureau-adapter/pkg/pb/cba/v1"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type GrpcParams struct {
	fx.In
	Config *config.Config
	CBA    cba.CreditBureauAdapterServiceServer
}

func NewGrpcServer(lc fx.Lifecycle, logger *zap.Logger, params GrpcParams) *grpc.Server {
	srv := grpc.NewServer()

	cba.RegisterCreditBureauAdapterServiceServer(srv, params.CBA)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Info("Starting gRPC server")

			ln, err := net.Listen("tcp", fmt.Sprintf(":%d", params.Config.GRPC.Port))
			if err != nil {
				return err
			}

			go func() {
				if err := srv.Serve(ln); err != nil {
					logger.Error("Failed to Serve gRPC", zap.Error(err))
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("Gracefully stopping gRPC server")

			srv.GracefulStop()

			return nil
		},
	})

	return srv
}
