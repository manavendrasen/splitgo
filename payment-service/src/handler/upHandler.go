package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAppStatus(c echo.Context) error {
	status := make(map[string]string, 1)
	status["Status"] = "Up"
	return c.JSON(http.StatusOK, status)
}
