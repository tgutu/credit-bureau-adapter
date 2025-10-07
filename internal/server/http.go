package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/tgutu/credit-bureau-adapter/internal/config"
	"github.com/tgutu/credit-bureau-adapter/pkg/pb/cba/v1"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type HttpParams struct {
	fx.In
	Config *config.Config
	Logger *zap.Logger
}

func NewHTTPServer(lc fx.Lifecycle, params HttpParams) *http.Server {
	logger := params.Logger.Named("http")
	gwmux := runtime.NewServeMux()

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", params.Config.HTTP.Port),
		Handler: gwmux,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Info("Starting HTTP server", zap.String("addr", srv.Addr))

			// Create a client connection to the gRPC server
			// The gRPC-Gateway proxies the requests over HTTP
			conn, err := grpc.NewClient(
				fmt.Sprintf(":%d", params.Config.GRPC.Port),
				grpc.WithTransportCredentials(insecure.NewCredentials()),
			)
			if err != nil {
				logger.Fatal("Failed to dial server:", zap.Error(err))
			}

			if err := cba.RegisterCreditBureauAdapterServiceHandler(ctx, gwmux, conn); err != nil {
				return fmt.Errorf("failed to register api key gateway: %w", err)
			}

			go func() {
				if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					logger.Error("Failed to ListenAndServe HTTP", zap.Error(err))
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("Shutting down HTTP server")

			if err := srv.Shutdown(ctx); err != nil {
				logger.Error("Failed to shutdown HTTP server", zap.Error(err))
				return err
			}

			return nil
		},
	})

	return srv
}
