package middleware

import (
	"time"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// RequestLoggerMiddleware logs API requests in a structured format
func RequestLoggerMiddleware(logger *zap.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()

			// Process request
			err := next(c)

			// Log request details
			logger.Info("API Request",
				zap.String("method", c.Request().Method),
				zap.String("path", c.Request().URL.Path),
				zap.Int("status", c.Response().Status),
				zap.String("latency", time.Since(start).String()),
				zap.String("user-agent", c.Request().UserAgent()),
				zap.String("remote-ip", c.RealIP()),
			)

			return err
		}
	}
}
