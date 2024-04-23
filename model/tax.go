package model

import (
	"errors"

	"gorm.io/gorm"
)

type TaxCalculation struct {
	ID                uint        `gorm:"primaryKey"`
	TotalIncome       float64     `db:"totalIncome"`
	WHT               float64     `db:"wht"`
	PersonalAllowance float64     `db:"personal_allowance"`
	Donation          float64     `db:"donation"`
	KReceipt          float64     `db:"k_receipt"`
	Tax               float64     `db:"tax"`
	TaxPayable        float64     `db:"tax_payable"`
	Allowances        []Allowance `json:"allowances"`
}
type TaxRate struct {
	Level string  `json:"level"`
	Tax   float64 `json:"tax"`
}

type Allowance struct {
	AllowanceType string  `json:"allowanceType"`
	Amount        float64 `json:"amount"`
}

type TaxCalculationResponse struct {
	Tax float64 `json:"tax"`
}

func (t *TaxCalculation) BeforeSave(tx *gorm.DB) (err error) {
	// Validate fields before saving
	if t.TotalIncome < 0 {
		return errors.New("income must be positive")
	}

	return nil
}