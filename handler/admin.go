// admin
package handler

import (
	"net/http"

	"github.com/LGROW101/assessment-tax/model"
	"github.com/LGROW101/assessment-tax/repository"
	"github.com/labstack/echo/v4"
)

type AdminHandler struct {
	adminRepo repository.AdminRepository
}

func NewAdminHandler(adminRepo repository.AdminRepository) *AdminHandler {
	return &AdminHandler{
		adminRepo: adminRepo,
	}
}

func (h *AdminHandler) GetConfig(c echo.Context) error {
	config, err := h.adminRepo.GetConfig()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, config)
}

func (h *AdminHandler) UpdateConfig(c echo.Context) error {
	var req model.AdminConfig
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	config, err := h.adminRepo.GetConfig()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if req.PersonalDeduction != 0 {
		config.PersonalDeduction = req.PersonalDeduction
	}
	if req.KReceipt != 0 {
		config.KReceipt = req.KReceipt
	}

	if err := config.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = h.adminRepo.UpdateConfig(config)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, config)
}
