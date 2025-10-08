package adapter

import (
	"context"
	"time"

	"github.com/tgutu/credit-bureau-adapter/internal/config"
	"github.com/tgutu/credit-bureau-adapter/pkg/pb/cba/v1"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// TransUnionAdapter is implementation of CreditBureauAdapter for TransUnio
type TransUnionAdapter struct {
	baseURL string
	apiKey  string
	logger  *zap.Logger
}

type TransUnionAdapterParams struct {
	fx.In
	Config *config.Config
	Logger *zap.Logger
}

func NewTransUnionAdapter(params TransUnionAdapterParams) *TransUnionAdapter {
	return &TransUnionAdapter{
		apiKey:  params.Config.TransUnion.APIKey,
		baseURL: params.Config.TransUnion.BaseURL,
		logger:  params.Logger,
	}
}

func (t *TransUnionAdapter) GetCreditReport(ctx context.Context, req *cba.GetCreditReportRequest) (cba.GetCreditReportResponse, error) {
	// In reality, you’d call the TransUnion API here using http.Client.
	// For demonstration, we simulate a response.
	t.logger.Info("[TransUnion] Fetching report",
		zap.String("first_name", req.FirstName),
		zap.String("last_name", req.LastName),
	)

	time.Sleep(100 * time.Millisecond)

	return cba.GetCreditReportResponse{
		ReportId:    "TU-REPORT-67890",
		FullName:    req.FirstName + " " + req.LastName,
		DateOfBirth: req.DateOfBirth,
		ReportDate:  time.Now().Format("2006-01-02"),
		Accounts: []*cba.Account{
			{
				AccountType:   "Mortgage",
				AccountNumber: "9876-5432-1098-7654",
				Balance:       "R200000.00",
				CreditLimit:   "R300000.00",
				Status:        "Open",
				OpenedDate:    "2018-05-20",
			},
		},
		Inquiries: []*cba.Inquiry{
			{
				Date:        time.Now().AddDate(0, -1, 0).Format("2006-01-02"),
				InquiryType: "Credit Card",
				Institution: "Some Bank",
			},
		},
	}, nil
}

func (t *TransUnionAdapter) GetCreditScore(ctx context.Context, req *cba.GetCreditScoreRequest) (cba.GetCreditScoreResponse, error) {
	// In reality, you’d call the TransUnion API here using http.Client.
	// For demonstration, we simulate a response.
	t.logger.Info("[TransUnion] Fetching score",
		zap.String("first_name", req.FirstName),
		zap.String("last_name", req.LastName),
	)

	time.Sleep(50 * time.Millisecond)

	return cba.GetCreditScoreResponse{
		FullName:    req.FirstName + " " + req.LastName,
		CreditScore: 680,
		ScoreDate:   time.Now().Format("2006-01-02"),
	}, nil
}
