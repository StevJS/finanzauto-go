package middleware

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

const API_KEY = "sk_test_12345"

func APIKeyMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		apiKey := c.Request().Header.Get("X-API-Key")

		if apiKey == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "missing API key")
		}

		if apiKey != API_KEY {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid API key")
		}

		return next(c)
	}
}
