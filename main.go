package main

import (
	"github.com/LGROW101/assessment-tax/handler"
	"github.com/LGROW101/assessment-tax/service"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	taxCalculatorService := &service.TaxCalculatorService{}
	taxHandler := &handler.TaxHandler{TaxCalculatorService: taxCalculatorService}

	e.POST("/tax/calculations", taxHandler.CalculateTax)

	e.Logger.Fatal(e.Start(":8080"))
}
