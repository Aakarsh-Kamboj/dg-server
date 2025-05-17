package v2

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// GreetHandler returns a greeting message for v2 (if different)
func GreetHandler(c echo.Context) error {
	name := c.QueryParam("name")
	if name == "" {
		name = "Guest"
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Hello, " + name + "! This is the upgraded API v2!",
	})
}
