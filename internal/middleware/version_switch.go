package middleware

import (
	"github.com/labstack/echo/v4"
)

// VersionSwitchMiddleware detects the requested API version from headers
func VersionSwitchMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Check Accept-Version header
			requestedVersion := c.Request().Header.Get("Accept-Version")

			// If a valid version is provided, override the URL version
			if requestedVersion != "" {
				c.Set("version", requestedVersion)
			}
			return next(c)
		}
	}
}
