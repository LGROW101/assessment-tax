// Calculator
package handler

import (
	"net/http"

	"github.com/LGROW101/assessment-tax/model"
	"github.com/LGROW101/assessment-tax/service"
	"github.com/labstack/echo/v4"
)

type CalculatorHandler struct {
	taxCalculatorService service.TaxCalculatorService
}

func NewCalculatorHandler(taxCalculatorService service.TaxCalculatorService) *CalculatorHandler {
	return &CalculatorHandler{
		taxCalculatorService: taxCalculatorService,
	}
}

type CalculateTaxRequest struct {
	TotalIncome     float64           `json:"totalIncome"`
	WHT             float64           `json:"wht"`
	Allowances      []model.Allowance `json:"allowances"`
	IncludeTaxLevel bool              `json:"includeTaxLevel"`
}

func (h *CalculatorHandler) CalculateTax(c echo.Context) error {
	var req CalculateTaxRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Check if the request body is valid
	if req.TotalIncome < 0 || req.WHT < 0 || len(req.Allowances) == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
	}
	taxCalculation, err := h.taxCalculatorService.CalculateTax(req.TotalIncome, req.WHT, req.Allowances)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	response := map[string]interface{}{
		"tax": taxCalculation.TaxPayable,
	}

	if req.IncludeTaxLevel {
		response["taxLevel"] = taxCalculation.TaxLevel
	}

	return c.JSON(http.StatusOK, response)
}

func (h *CalculatorHandler) GetAllCalculations(c echo.Context) error {
	taxCalculations, err := h.taxCalculatorService.GetAllCalculations()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, taxCalculations)
}
