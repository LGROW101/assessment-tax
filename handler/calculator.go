package handler

import (
	"net/http"

	"github.com/LGROW101/assessment-tax/model"
	"github.com/LGROW101/assessment-tax/service"
	"github.com/labstack/echo/v4"
)

type TaxHandler struct {
	TaxCalculatorService *service.TaxCalculatorService
}

func (h *TaxHandler) CalculateTax(c echo.Context) error {
	var req model.TaxCalculationRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	resp, err := h.TaxCalculatorService.CalculateTax(&req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}
