package handler

import (
	"errors"
	"net/http"
	"payment-service/middleware"
	"payment-service/repository"
	"strings"

	"github.com/labstack/echo/v4"
)

func GetPayments(c echo.Context) error {
	ac := c.(*middleware.AuthContext)
	ID, _, _ := ac.GetCurrentUser()
	payments, err := repository.ViewPayment(ID)

	if err != nil {
		return errors.New("DB_ERROR")
	}

	return c.JSON(http.StatusOK, payments)
}

func AddPayment(c echo.Context) error {
	ac := c.(*middleware.AuthContext)
	ID, _, _ := ac.GetCurrentUser()

	var body struct {
		// ToUserID    uint    `json:"to_user_id"`
		To          string  `json:"to" validate:"required"`
		Amount      float32 `json:"amount" validate:"required"`
		Description string  `json:"description"`
	}

	err := c.Bind(&body)

	if err != nil {
		status := &Status{
			Message: "INVALID_BODY",
		}
		return c.JSON(http.StatusBadGateway, status)
	}

	if len(strings.Trim(body.To, " ")) == 0 {
		status := &Status{
			Message: "TO_FIELD_REQUIRED",
		}
		return c.JSON(http.StatusBadRequest, status)
	}

	payment, err := repository.AddPayment(ID, body.To, body.Amount, body.Description)

	if err != nil {
		status := &Status{
			Message: "DB_ERROR",
		}
		return c.JSON(http.StatusBadRequest, status)
	}

	return c.JSON(http.StatusOK, payment)
}

func UpdatePayment(c echo.Context) error {
	ac := c.(*middleware.AuthContext)
	ID, _, _ := ac.GetCurrentUser()

	var body struct {
		PaymentID   uint    `json:"payment_id" validate:"required"`
		To          string  `json:"to" validate:"required"`
		Amount      float32 `json:"amount" validate:"required"`
		Description string  `json:"description"`
	}

	err := c.Bind(&body)

	if err != nil {
		status := &Status{
			Message: "INVALID_BODY",
		}
		return c.JSON(http.StatusBadGateway, status)
	}

	if len(strings.Trim(body.To, " ")) == 0 {
		status := &Status{
			Message: "TO_FIELD_REQUIRED",
		}
		return c.JSON(http.StatusBadRequest, status)
	}

	// add validation if all fields are present or not

	payment, err := repository.UpdatePayment(ID, body.PaymentID, body.Amount, body.To, body.Description)

	if err != nil {
		status := &Status{
			Message: err.Error(),
		}
		return c.JSON(http.StatusBadRequest, status)
	}

	return c.JSON(http.StatusOK, payment)
}

func DeletePayment(c echo.Context) error {
	ac := c.(*middleware.AuthContext)
	ID, _, _ := ac.GetCurrentUser()

	var body struct {
		PaymentID uint `query:"payment_id" validate:"required"`
	}

	err := c.Bind(&body)

	if err != nil {
		status := &Status{
			Message: "INVALID_PARAMS",
		}
		return c.JSON(http.StatusBadGateway, status)
	}

	payment, err := repository.DeletePayment(ID, body.PaymentID)

	if err != nil {
		status := &Status{
			Message: err.Error(),
		}
		return c.JSON(http.StatusBadRequest, status)
	}

	result := make(map[string]int)
	result["rowsAffected"] = int(payment)
	return c.JSON(http.StatusOK, result)
}
