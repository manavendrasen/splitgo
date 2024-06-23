package handler

import (
	"log"
	"net/http"
	"payment-service/middleware"

	"github.com/labstack/echo/v4"
)

func GetPayments(c echo.Context) error {
	ac := c.(*middleware.AuthContext)
	ID, Email, DisplayName := ac.GetCurrentUser()
	log.Print(ID)
	log.Print(Email)
	log.Print(DisplayName)
	return c.String(http.StatusOK, Email)
}

func AddPayment(c echo.Context) error {
	return c.String(http.StatusOK, "Add")
}

func UpdatePayment(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "Update "+id)
}

func DeletePayment(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "Delete "+id)
}
