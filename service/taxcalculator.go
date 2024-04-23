package service

import (
	"github.com/LGROW101/assessment-tax/model"
	"github.com/LGROW101/assessment-tax/repository"
)

type TaxCalculatorService interface {
	GetAllCalculations() ([]*model.TaxCalculation, error)
	CalculateTax(income, wht float64, allowances []model.Allowance) (*model.TaxCalculation, error)
}
type taxCalculatorService struct {
	taxRepo  repository.TaxRepository
	adminSvc AdminServiceInterface
}

func NewTaxCalculatorService(taxRepo repository.TaxRepository, adminRepo repository.AdminRepository) TaxCalculatorService {
	return &taxCalculatorService{
		taxRepo:  taxRepo,
		adminSvc: NewAdminService(adminRepo),
	}
}

func (s *taxCalculatorService) GetAllCalculations() ([]*model.TaxCalculation, error) {
	return s.taxRepo.GetAllCalculations()
}

func (s *taxCalculatorService) CalculateTax(totalIncome, wht float64, allowances []model.Allowance) (*model.TaxCalculation, error) {

	config, err := s.adminSvc.GetConfig() // Use the GetConfig method from the AdminServiceInterface
	if err != nil {
		return nil, err
	}
	// Set default values if not provided
	personalAllowance := config.PersonalDeduction
	donation := 0.0
	kReceipt := 0.0

	for _, allowance := range allowances {
		switch allowance.AllowanceType {
		case "donation":
			donation = allowance.Amount
			if donation > 100000 {
				donation = 100000
			}
		case "k-receipt":
			kReceipt = allowance.Amount
			if kReceipt > config.KReceipt {
				kReceipt = config.KReceipt
			}
		}
	}

	// Calculate taxable income
	taxableIncome := totalIncome - personalAllowance - donation - kReceipt

	// Calculate tax
	var tax float64
	switch {
	case taxableIncome <= 0:
		tax = 0
	case taxableIncome <= 150000:
		tax = taxableIncome * 0.05
	case taxableIncome <= 300000:
		tax = 7500 + (taxableIncome-150000)*0.05
	case taxableIncome <= 500000:
		tax = 15000 + (taxableIncome-300000)*0.10
	case taxableIncome <= 750000:
		tax = 35000 + (taxableIncome-500000)*0.15
	case taxableIncome <= 1000000:
		tax = 57500 + (taxableIncome-750000)*0.20
	case taxableIncome <= 2000000:
		tax = 107500 + (taxableIncome-1000000)*0.25
	case taxableIncome <= 5000000:
		tax = 357500 + (taxableIncome-2000000)*0.30
	default:
		tax = 1257500 + (taxableIncome-5000000)*0.35
	}

	taxPayable := tax - wht
	if taxPayable < 0 {
		taxPayable = 0
	}

	// Calculate tax levels
	taxLevel := []model.TaxRate{
		{Level: "0-150,000", Tax: 0},
		{Level: "150,001-500,000", Tax: 0},
		{Level: "500,001-1,000,000", Tax: 0},
		{Level: "1,000,001-2,000,000", Tax: 0},
		{Level: "2,000,001 ขึ้นไป", Tax: 0},
	}

	switch {
	case taxableIncome <= 150000:
		taxLevel[0].Tax = taxPayable
	case taxableIncome <= 500000:
		taxLevel[1].Tax = taxPayable
	case taxableIncome <= 1000000:
		taxLevel[2].Tax = taxPayable
	case taxableIncome <= 2000000:
		taxLevel[3].Tax = taxPayable
	default:
		taxLevel[4].Tax = taxPayable
	}

	taxCalculation := &model.TaxCalculation{
		TotalIncome:       totalIncome,
		WHT:               wht,
		PersonalAllowance: personalAllowance,
		Donation:          donation,
		KReceipt:          kReceipt,
		Tax:               taxPayable,
		TaxPayable:        taxPayable,
		TaxLevel:          taxLevel,
	}

	err = s.taxRepo.Save(taxCalculation)
	if err != nil {
		return nil, err
	}

	return taxCalculation, nil
}
