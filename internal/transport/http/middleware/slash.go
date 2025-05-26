package middleware

import (
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func TrailingSlash() func(next echo.HandlerFunc) echo.HandlerFunc {
	return middleware.AddTrailingSlashWithConfig(middleware.TrailingSlashConfig{
		Skipper: func(c echo.Context) bool {
			return strings.Contains(c.Request().RequestURI, "swagger")
		},
	})
}
