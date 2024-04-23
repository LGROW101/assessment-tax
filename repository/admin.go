// admin
package repository

import (
	"database/sql"

	"github.com/LGROW101/assessment-tax/model"
)

type AdminRepository interface {
	GetConfig() (*model.AdminConfig, error)
	UpdateConfig(config *model.AdminConfig) error
}

type adminRepository struct {
	db *sql.DB
}

func NewAdminRepository(db *sql.DB) AdminRepository {
	return &adminRepository{db: db}
}

func (r *adminRepository) GetConfig() (*model.AdminConfig, error) {
	query := `
        SELECT id, personal_deduction, k_receipt, created_at, updated_at
        FROM admin_configs
        ORDER BY id DESC
        LIMIT 1
    `
	row := r.db.QueryRow(query)
	var config model.AdminConfig
	err := row.Scan(&config.ID, &config.PersonalDeduction, &config.KReceipt, &config.CreatedAt, &config.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &config, nil
}

func (r *adminRepository) UpdateConfig(config *model.AdminConfig) error {
	query := `
        UPDATE admin_configs
        SET personal_deduction = $1, k_receipt = $2
        WHERE id = (SELECT id FROM admin_configs ORDER BY id DESC LIMIT 1)
    `
	_, err := r.db.Exec(query, config.PersonalDeduction, config.KReceipt)
	return err
}
