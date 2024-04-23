package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	// Create a new Echo instance
	e := echo.New()

	// Set up the routes
	setupRoutes(e)

	// Create a new HTTP request for the root endpoint
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	// Perform the request
	e.ServeHTTP(rec, req)

	// Assert the response
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, `{"message": "Hello, world!"}`, rec.Body.String())
}

func setupRoutes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		responseData := map[string]interface{}{
			"message": "Hello, world!",
		}
		return c.JSON(http.StatusOK, responseData)
	})
}
