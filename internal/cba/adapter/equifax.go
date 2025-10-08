package adapter

import (
	"context"
	"time"

	"github.com/tgutu/credit-bureau-adapter/internal/config"
	"github.com/tgutu/credit-bureau-adapter/pkg/pb/cba/v1"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// EquifaxAdapter implements the CreditBureauAdapter interface.
type EquifaxAdapter struct {
	baseURL string
	apiKey  string
	logger  *zap.Logger
}

type EquifaxAdapterParams struct {
	fx.In
	Config *config.Config
	Logger *zap.Logger
}

func NewEquifaxAdapter(params EquifaxAdapterParams) *EquifaxAdapter {
	return &EquifaxAdapter{
		apiKey:  params.Config.Equifax.APIKey,
		baseURL: params.Config.Equifax.BaseURL,
		logger:  params.Logger,
	}
}

func (e *EquifaxAdapter) GetCreditReport(ctx context.Context, req *cba.GetCreditReportRequest) (cba.GetCreditReportResponse, error) {
	// In reality, you’d call the Equifax API here using http.Client.
	// For demonstration, we simulate a response.
	e.logger.Info("[Equifax] Fetching report",
		zap.String("first_name", req.FirstName),
		zap.String("last_name", req.LastName),
	)

	time.Sleep(100 * time.Millisecond)

	return cba.GetCreditReportResponse{
		ReportId:    "EQ-REPORT-12345",
		FullName:    req.FirstName + " " + req.LastName,
		DateOfBirth: req.DateOfBirth,
		ReportDate:  time.Now().Format("2006-01-02"),
		Accounts: []*cba.Account{
			{
				AccountType:   "Credit Card",
				AccountNumber: "1234-5678-9012-3456",
				Balance:       "R1000.00",
				CreditLimit:   "R5000.00",
				Status:        "Open",
				OpenedDate:    "2020-01-15",
			},
		},
		Inquiries: []*cba.Inquiry{
			{
				Date:        "2023-10-01",
				InquiryType: "Credit Card Application",
				Institution: "Bank A",
			},
		},
		PublicRecords: []*cba.PublicRecord{
			{
				RecordType: "Judgment",
				DateFiled:  "2022-05-20",
				Status:     "Satisfied",
			},
		},
	}, nil
}

func (e *EquifaxAdapter) GetCreditScore(ctx context.Context, req *cba.GetCreditScoreRequest) (cba.GetCreditScoreResponse, error) {
	// In reality, you’d call the Equifax API here using http.Client.
	// For demonstration, we simulate a response.
	e.logger.Info("[Equifax] Fetching score",
		zap.String("first_name", req.FirstName),
		zap.String("last_name", req.LastName),
	)

	time.Sleep(100 * time.Millisecond)

	return cba.GetCreditScoreResponse{
		FullName:    req.FirstName + " " + req.LastName,
		DateOfBirth: req.DateOfBirth,
		CreditScore: 700,
		ScoreDate:   time.Now().Format("2006-01-02"),
	}, nil
}
