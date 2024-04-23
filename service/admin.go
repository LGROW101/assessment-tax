// service/admin.go
package service

import (
	"github.com/LGROW101/assessment-tax/model"
	"github.com/LGROW101/assessment-tax/repository"
)

// AdminServiceInterface defines the methods that need to be implemented by the AdminService
type AdminServiceInterface interface {
	GetConfig() (*model.AdminConfig, error)
	UpdateConfig(config *model.AdminConfig) error
}

type AdminService struct {
	adminRepo repository.AdminRepository
}

// NewAdminService returns a new instance of AdminService
func NewAdminService(adminRepo repository.AdminRepository) AdminServiceInterface {
	return &AdminService{
		adminRepo: adminRepo,
	}
}

func (s *AdminService) GetConfig() (*model.AdminConfig, error) {
	// Logika untuk mengambil konfigurasi admin dari repository
	return s.adminRepo.GetConfig()
}

func (s *AdminService) UpdateConfig(config *model.AdminConfig) error {
	// Logika untuk memperbarui konfigurasi admin di repository
	return s.adminRepo.UpdateConfig(config)
}
