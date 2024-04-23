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

	// ตรวจสอบว่าภาษีที่คำนวณได้มากกว่าภาษีที่หัก ณ ที่จ่ายหรือไม่
	if tax > req.WHT {
		// ถ้ามากกว่า ให้ลดภาษีที่ต้องจ่ายลงด้วยจำนวนภาษีที่หัก ณ ที่จ่าย
		tax -= req.WHT
	} else {
		// ถ้าน้อยกว่าหรือเท่ากัน แสดงว่าได้หักภาษี ณ ที่จ่ายครบแล้ว ไม่ต้องจ่ายภาษีเพิ่ม
		tax = 0
	}

	return &model.TaxCalculationResponse{Tax: tax}, nil
}
