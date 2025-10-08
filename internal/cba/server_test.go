package cba

import (
	"context"
	"testing"

	"github.com/tgutu/credit-bureau-adapter/internal/apicode"
	"github.com/tgutu/credit-bureau-adapter/internal/cba/adapter"
	"github.com/tgutu/credit-bureau-adapter/internal/config"
	"github.com/tgutu/credit-bureau-adapter/internal/database"
	"github.com/tgutu/credit-bureau-adapter/pkg/pb/cba/v1"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
	"go.uber.org/zap"
)

func TestCreditBureauAdapterService(t *testing.T) {
	var svc cba.CreditBureauAdapterServiceServer

	configFile, err := config.NewConfig("testdata/cba.yaml")
	if err != nil {
		t.Fatalf("failed to load config: %v", err)
	}

	app := fxtest.New(t,
		fx.Provide(
			NewServer,
			adapter.NewExperianAdapter,
			adapter.NewEquifaxAdapter,
			adapter.NewTransUnionAdapter,
			database.NewDatabase,
			database.NewCreditBureauRepository,
			zap.NewExample,
		),
		fx.Provide(func() *config.Config {
			return configFile
		}),
		fx.Invoke(func(s cba.CreditBureauAdapterServiceServer) {
			svc = s // Capture the service instance for testing
		}),
	)

	app.RequireStart()
	defer app.RequireStop() // Ensure cleanup even if the test fails

	// Now you can test the captured service
	ctx := context.Background()

	// Test GetBureaus
	t.Run("GetBureaus", func(t *testing.T) {
		bureausResp, err := svc.GetBureaus(ctx, &cba.GetBureausRequest{})
		if err != nil {
			t.Errorf("GetBureaus failed: %v", err)
		} else if len(bureausResp.Bureaus) == 0 {
			t.Error("GetBureaus returned no bureaus")
		}
	})

	// Test GetCreditReport with a valid bureau
	t.Run("GetCreditReport_ValidBureau", func(t *testing.T) {
		r, err := svc.GetCreditReport(ctx, &cba.GetCreditReportRequest{FirstName: "John", LastName: "Doe", BureauName: "experian"})
		if err != nil {
			t.Errorf("GetCreditReport failed: %v", err)
		}

		if r.FullName != "John Doe" {
			t.Errorf("Expected FullName 'John Doe', got '%s'", r.FullName)
		}
	})

	// Test GetCreditReport with an invalid bureau
	t.Run("GetCreditReport_InvalidBureau", func(t *testing.T) {
		_, err := svc.GetCreditReport(ctx, &cba.GetCreditReportRequest{BureauName: "invalid_bureau"})
		if err == nil {
			t.Error("GetCreditReport with invalid bureau should have failed")
		}

		if err != nil && err.Error() != apicode.ErrCreditRepoGetBureauByNameFailed.Error() {
			t.Errorf("Expected error '%v', got '%v'", apicode.ErrCreditRepoGetBureauByNameFailed.Error(), err)
		}
	})

	// Test GetCreditScore with a valid bureau
	t.Run("GetCreditScore_ValidBureau", func(t *testing.T) {
		scoreResp, err := svc.GetCreditScore(ctx, &cba.GetCreditScoreRequest{FirstName: "John", LastName: "Doe", BureauName: "experian"})
		if err != nil {
			t.Errorf("GetCreditScore failed: %v", err)
		}

		if scoreResp.FullName != "John Doe" {
			t.Errorf("Expected FullName 'John Doe', got '%s'", scoreResp.FullName)
		}
	})

	// Test GetCreditScore with an invalid bureau
	t.Run("GetCreditScore_InvalidBureau", func(t *testing.T) {
		_, err := svc.GetCreditScore(ctx, &cba.GetCreditScoreRequest{BureauName: "invalid_bureau"})
		if err == nil {
			t.Error("GetCreditScore with invalid bureau should have failed")
		}

		if err != nil && err.Error() != apicode.ErrCreditRepoGetBureauByNameFailed.Error() {
			t.Errorf("Expected error '%v', got '%v'", apicode.ErrCreditRepoGetBureauByNameFailed.Error(), err)
		}
	})
}
