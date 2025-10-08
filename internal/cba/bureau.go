package cba

import (
	"context"

	"github.com/tgutu/credit-bureau-adapter/internal/cba/adapter"
	"github.com/tgutu/credit-bureau-adapter/pkg/pb/cba/v1"
)

type CreditBureauAdapter interface {
	GetCreditReport(ctx context.Context, req *cba.GetCreditReportRequest) (cba.GetCreditReportResponse, error)
	GetCreditScore(ctx context.Context, req *cba.GetCreditScoreRequest) (cba.GetCreditScoreResponse, error)
}

// Ensure EquifaxAdapter implements the CreditBureauAdapter interface.
var _ CreditBureauAdapter = (*adapter.EquifaxAdapter)(nil)

// Ensure TransUnionAdapter implements the CreditBureauAdapter interface.
var _ CreditBureauAdapter = (*adapter.TransUnionAdapter)(nil)
