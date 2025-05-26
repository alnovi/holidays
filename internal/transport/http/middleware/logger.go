package middleware

import (
	"log/slog"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RequestLogger(logger *slog.Logger) func(next echo.HandlerFunc) echo.HandlerFunc {
	config := middleware.RequestLoggerConfig{
		HandleError:     false,
		LogRemoteIP:     true,
		LogUserAgent:    true,
		LogProtocol:     true,
		LogLatency:      true,
		LogMethod:       true,
		LogURI:          true,
		LogStatus:       true,
		LogError:        true,
		LogResponseSize: true,
		LogValuesFunc: func(c echo.Context, values middleware.RequestLoggerValues) error {
			isFavicon := strings.Contains(values.URI, "favicon")
			isAssets := strings.Contains(values.URI, "assets")
			isPublic := strings.Contains(values.URI, "public")

			if isFavicon || isAssets || isPublic {
				return nil
			}

			log := logger.With(
				slog.String("ip", values.RemoteIP),
				slog.String("agent", values.UserAgent),
				slog.String("protocol", values.Protocol),
				slog.Duration("latency", values.Latency),
				slog.String("method", values.Method),
				slog.String("uri", strings.TrimRight(values.URI, "/")),
				slog.Int("status", values.Status),
				slog.Int64("response_size", values.ResponseSize),
			)

			if values.Error != nil {
				log.Error(values.Error.Error())
			} else {
				log.Info(http.StatusText(values.Status))
			}

			return nil
		},
	}

	return middleware.RequestLoggerWithConfig(config)
}
