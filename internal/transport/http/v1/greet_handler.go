package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// GreetHandler returns a greeting message for v1
func GreetHandler(c echo.Context) error {
	name := c.QueryParam("name")
	if name == "" {
		name = "Guest"
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Hello, " + name + "! Welcome to API Gateway v1!",
	})
}
