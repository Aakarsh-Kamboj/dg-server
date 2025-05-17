package v1

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

// HealthResponse represents the health check response
type HealthResponse struct {
	Status string `json:"status"`
	Uptime string `json:"uptime"`
}

var startTime = time.Now()

// HealthCheckHandler returns API health status
func HealthCheckHandler(c echo.Context) error {
	uptime := time.Since(startTime).String()
	response := HealthResponse{
		Status: "healthy",
		Uptime: uptime,
	}
	return c.JSON(http.StatusOK, response)
}
