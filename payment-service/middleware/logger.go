package middleware

import (
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

func Logger() echo.MiddlewareFunc {
	return echoMiddleware.LoggerWithConfig(echoMiddleware.LoggerConfig{
		Format:           `${time_rfc3339} >> [${status}][${method}] ${uri} ${error} (${latency_human})` + "\n",
		CustomTimeFormat: "2006-01-02 15:04:05.00000",
	})
}
