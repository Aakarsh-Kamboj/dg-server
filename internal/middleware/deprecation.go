package middleware

import (
	"github.com/labstack/echo/v4"
)

// DeprecatedAPIVersions defines versions that are deprecated
var DeprecatedAPIVersions = map[string]bool{
	"v1": true, // Mark v1 as deprecated
}

// DeprecationMiddleware adds a deprecation warning header
func DeprecationMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Extract an API version from URL
			version := c.Param("version")
			if DeprecatedAPIVersions[version] {
				c.Response().Header().Set("X-API-Deprecated", "true")
				c.Response().Header().Set("Warning", "299 - 'API version "+version+" is deprecated. Please upgrade to the latest version.'")
			}
			return next(c)
		}
	}
}
