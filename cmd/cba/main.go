package main

import (
	"flag"
	"net/http"

	"github.com/tgutu/credit-bureau-adapter/internal/cba"
	"github.com/tgutu/credit-bureau-adapter/internal/cba/adapter"
	"github.com/tgutu/credit-bureau-adapter/internal/config"
	"github.com/tgutu/credit-bureau-adapter/internal/database"
	"github.com/tgutu/credit-bureau-adapter/internal/server"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	configFilename := flag.String("port", "cba.yaml", "Path to the configuration file")
	flag.Parse()

	configFile, err := config.NewConfig(*configFilename)
	if err != nil {
		panic(err)
	}

	app := createApp(configFile)

	app.Run()
}

func createApp(configFile *config.Config) *fx.App {
	return fx.New(
		fx.Invoke(
			func(*grpc.Server) {},
			func(*http.Server) {},
		),
		fx.Provide(func() *config.Config {
			return configFile
		}),
		fx.Provide(
			cba.NewServer,
			adapter.NewExperianAdapter,
			adapter.NewEquifaxAdapter,
			adapter.NewTransUnionAdapter,
			database.NewDatabase,
			database.NewCreditBureauRepository,
			server.NewGrpcServer,
			server.NewHTTPServer,
			zap.NewProduction,
		),
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
	)
}
