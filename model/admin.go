package model

import (
	"errors"
	"time"
)

type AdminConfig struct {
	ID                uint      `gorm:"primaryKey" db:"id"`
	PersonalDeduction float64   `db:"personal_deduction"`
	KReceipt          float64   `db:"k_receipt"`
	CreatedAt         time.Time `db:"created_at"`
	UpdatedAt         time.Time `db:"updated_at"`
}

func (c *AdminConfig) Validate() error {
	// Validate fields
	if c.PersonalDeduction <= 0 {
		return errors.New("personal allowance max must be positive")
	}
	if c.KReceipt <= 0 {
		return errors.New("k-receipt max must be positive")
	}
	if c.PersonalDeduction < 0 {
		return errors.New("personal deduction must be non-negative")
	}
	if c.KReceipt < 0 {
		return errors.New("k-receipt must be non-negative")
	}
	return nil
}
