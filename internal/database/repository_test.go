package database_test

import (
	"context"
	"testing"

	"github.com/tgutu/credit-bureau-adapter/internal/database"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
	"go.uber.org/zap"
)

func TestCreditBureauRepository(t *testing.T) {
	var repo database.CreditBureauRepository

	app := fxtest.New(t,
		fx.Provide(
			database.NewDatabase,
			database.NewCreditBureauRepository,
			zap.NewExample,
		),
		fx.Invoke(func(s database.CreditBureauRepository) {
			repo = s // Capture the service instance for testing
		}),
	)

	app.RequireStart()
	defer app.RequireStop() // Ensure cleanup even if the test fails

	// Now you can test the captured service
	ctx := context.Background()

	// Example test case
	bureaus, err := repo.ListBureaus(ctx)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(bureaus) == 0 {
		t.Errorf("Expected some bureaus, got none")
	}

	bureau, err := repo.GetBureauByName(ctx, "experian")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if bureau == nil || bureau.Name != "experian" {
		t.Errorf("Expected bureau 'experian', got %v", bureau)
	}
}
