package adapter

import (
	"context"
	"time"

	"github.com/tgutu/credit-bureau-adapter/internal/config"
	"github.com/tgutu/credit-bureau-adapter/pkg/pb/cba/v1"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type ExperianAdapter struct {
	baseURL string
	apiKey  string
	logger  *zap.Logger
}

type ExperianAdapterParams struct {
	fx.In
	Config *config.Config
	Logger *zap.Logger
}

func NewExperianAdapter(params ExperianAdapterParams) *ExperianAdapter {
	return &ExperianAdapter{
		baseURL: params.Config.Experian.BaseURL,
		apiKey:  params.Config.Experian.APIKey,
		logger:  params.Logger,
	}
}

func (e *ExperianAdapter) GetCreditReport(ctx context.Context, req *cba.GetCreditReportRequest) (cba.GetCreditReportResponse, error) {
	// In reality, you’d call the Experian API here using http.Client.
	// For demonstration, we simulate a response.
	e.logger.Info("[Experian] Fetching report",
		zap.String("first_name", req.FirstName),
		zap.String("last_name", req.LastName),
	)

	time.Sleep(100 * time.Millisecond)

	return cba.GetCreditReportResponse{
		ReportId:    "EX-REPORT-54321",
		FullName:    req.FirstName + " " + req.LastName,
		DateOfBirth: req.DateOfBirth,
		ReportDate:  time.Now().Format("2006-01-02"),
		Accounts: []*cba.Account{
			{
				AccountType:   "Personal Loan",
				AccountNumber: "5555-6666-7777-8888",
				Balance:       "R5000.00",
				CreditLimit:   "R15000.00",
				Status:        "Open",
				OpenedDate:    "2019-11-10",
			},
		},
		Inquiries: []*cba.Inquiry{
			{
				Date:        time.Now().AddDate(0, -2, 0).Format("2006-01-02"),
				InquiryType: "Personal Loan",
				Institution: "Another Bank",
			},
		},
	}, nil
}

func (e *ExperianAdapter) GetCreditScore(ctx context.Context, req *cba.GetCreditScoreRequest) (cba.GetCreditScoreResponse, error) {
	// In reality, you’d call the Experian API here using http.Client.
	// For demonstration, we simulate a response.
	e.logger.Info("[Experian] Fetching score",
		zap.String("first_name", req.FirstName),
		zap.String("last_name", req.LastName),
	)

	time.Sleep(50 * time.Millisecond)

	return cba.GetCreditScoreResponse{
		FullName:    req.FirstName + " " + req.LastName,
		DateOfBirth: req.DateOfBirth,
		CreditScore: 720,
		ScoreDate:   time.Now().Format("2006-01-02"),
	}, nil
}
