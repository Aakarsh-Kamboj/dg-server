package middleware

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// RequestIDMiddleware assigns a unique X-Request-ID to each request
func RequestIDMiddleware(logger *zap.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			requestID := c.Request().Header.Get("X-Request-ID")

			// If no request ID exists, generate a new one
			if requestID == "" {
				requestID = uuid.New().String()
				c.Request().Header.Set("X-Request-ID", requestID)
			}

			// Add Request ID to response header
			c.Response().Header().Set("X-Request-ID", requestID)

			// Log the request with request ID
			logger.Info("Incoming API Request",
				zap.String("method", c.Request().Method),
				zap.String("path", c.Request().URL.Path),
				zap.String("request_id", requestID),
				zap.String("user-agent", c.Request().UserAgent()),
				zap.String("remote-ip", c.RealIP()),
			)

			return next(c)
		}
	}
}
