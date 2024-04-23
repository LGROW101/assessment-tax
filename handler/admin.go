package handler

import (
	"net/http"

	"github.com/LGROW101/assessment-tax/model"
	"github.com/LGROW101/assessment-tax/repository"
	"github.com/labstack/echo/v4"
)

// AdminHandler is an interface for admin handler
type AdminHandler interface {
	GetConfig(c echo.Context) error
	UpdateConfig(c echo.Context) error
}

type adminHandler struct {
	adminRepo repository.AdminRepository
}

// NewAdminHandler creates a new instance of AdminHandler
func NewAdminHandler(adminRepo repository.AdminRepository) AdminHandler {
	return &adminHandler{
		adminRepo: adminRepo,
	}
}

func (h *adminHandler) GetConfig(c echo.Context) error {
	config, err := h.adminRepo.GetConfig()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if config == nil {
		return echo.NewHTTPError(http.StatusNotFound, "Admin config not found")
	}

	resp := &model.AdminResponse{
		PersonalDeduction: config.PersonalDeduction,
		KReceipt:          config.KReceipt,
	}
	return c.JSON(http.StatusOK, resp)
}

func (h *adminHandler) UpdateConfig(c echo.Context) error {
	var req model.AdminRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	config, err := h.adminRepo.GetConfig()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if config == nil {
		config = &model.AdminConfig{}
		err = h.adminRepo.InsertConfig(config)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	resp := &model.AdminResponse{
		PersonalDeduction: config.PersonalDeduction,
		KReceipt:          config.KReceipt,
	}

	if req.PersonalDeduction != nil {
		config.PersonalDeduction = *req.PersonalDeduction
		resp.PersonalDeduction = *req.PersonalDeduction
	}
	if req.KReceipt != nil {
		config.KReceipt = *req.KReceipt
		resp.KReceipt = *req.KReceipt
	}

	if err := config.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = h.adminRepo.UpdateConfig(config)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}
