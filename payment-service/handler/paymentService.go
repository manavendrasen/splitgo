package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetPayments(c echo.Context) error {
	return c.String(http.StatusOK, "Get")
}

func AddPayment(c echo.Context) error {
	return c.String(http.StatusOK, "Add")
}

func UpdatePayment(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "Update " + id)
}

func DeletePayment(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "Delete " + id)
}