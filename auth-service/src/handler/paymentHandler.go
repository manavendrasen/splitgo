package handler

import (
	"auth-service/src/middleware"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetPayments(c echo.Context) error {
	ac := c.(*middleware.AuthContext)
	ID, _, _ := ac.GetCurrentUser()


	return c.JSON(http.StatusOK, ID)	
}

func AddPayment(c echo.Context) error {
	ac := c.(*middleware.AuthContext)
	ID, _, _ := ac.GetCurrentUser()


	return c.JSON(http.StatusOK, ID)
}


func UpdatePayment(c echo.Context) error {
	ac := c.(*middleware.AuthContext)
	ID, _, _ := ac.GetCurrentUser()


	return c.JSON(http.StatusOK, ID)
}

func DeletePayment(c echo.Context) error {
	ac := c.(*middleware.AuthContext)
	ID, _, _ := ac.GetCurrentUser()


	return c.JSON(http.StatusOK, ID)
}
