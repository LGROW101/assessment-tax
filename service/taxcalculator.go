package service

import (
	"github.com/LGROW101/assessment-tax/model"
)

type TaxCalculatorService struct{}

func (s *TaxCalculatorService) CalculateTax(req *model.TaxCalculationRequest) (*model.TaxCalculationResponse, error) {
	var personalAllowance float64 = 60000
	var donationAllowance float64

	for _, allowance := range req.Allowances {
		if allowance.AllowanceType == "donation" {
			donationAllowance = allowance.Amount
			break
		}
	}

	netIncome := req.TotalIncome - personalAllowance - donationAllowance

	var tax float64
	switch {
	case netIncome <= 150000:
		tax = 0
	case netIncome <= 500000:
		tax = (netIncome - 150000) * 0.1
	case netIncome <= 1000000:
		tax = 35000 + (netIncome-500000)*0.15
	case netIncome <= 2000000:
		tax = 110000 + (netIncome-1000000)*0.2
	default:
		tax = 310000 + (netIncome-2000000)*0.35
	}

	return &model.TaxCalculationResponse{Tax: tax}, nil
}
