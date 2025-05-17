package v2

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// InfoHandler provides API version information (only in v2)
func InfoHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"version": "v2",
		"message": "Welcome to API v2! New features available!",
	})
}
