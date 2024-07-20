package handler

import (
	"net/http"
	"payment-service/middleware"
	"payment-service/repository"
	"payment-service/util"
	"strings"

	"github.com/labstack/echo/v4"
)

// GetPayments retrieves user payments and returns them as JSON.
func GetPayments(c echo.Context) error {
	ac := c.(*middleware.AuthContext)
	ID, _, _ := ac.GetCurrentUser()
	payments, err := repository.ViewPayment(ID)

	if err != nil {
		return c.JSON(http.StatusBadGateway, util.SendMessage(err.Error()))
	}

	return c.JSON(http.StatusOK, payments)
}

// AddPayment processes the request to add a new payment.
// It extracts the current user's ID from the authentication context,
// then reads and validates the request body to create a new payment entry.
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
		return c.JSON(http.StatusBadGateway, util.SendMessage("INVALID_BODY"))
	}

	if len(strings.Trim(body.To, " ")) == 0 {
		return c.JSON(http.StatusBadRequest, util.SendMessage("TO_FIELD_REQUIRED"))
	}

	payment, err := repository.AddPayment(ID, body.To, body.Amount, body.Description)

	if err != nil {

		return c.JSON(http.StatusBadRequest, util.SendMessage(err.Error()))
	}

	return c.JSON(http.StatusOK, payment)
}

// UpdatePayment updates an existing payment's details.
// It first verifies the user's identity from the authentication context,
// then validates the request body for the required fields. If validation passes,
// it proceeds to update the payment details in the repository.
func UpdatePayment(c echo.Context) error {
	// Extract the current user's ID from the authentication context
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
		return c.JSON(http.StatusBadGateway, util.SendMessage("INVALID_BODY"))
	}

	if len(strings.Trim(body.To, " ")) == 0 {
		return c.JSON(http.StatusBadRequest, util.SendMessage("TO_FIELD_REQUIRED"))
	}

	// TODO: add validation if all fields are present or not

	payment, err := repository.UpdatePayment(ID, body.PaymentID, body.Amount, body.To, body.Description)

	if err != nil {
		return c.JSON(http.StatusBadRequest, util.SendMessage(err.Error()))
	}

	return c.JSON(http.StatusOK, payment)
}

// DeletePayment handles the request to delete a specific payment.
// It extracts the current user's ID from the authentication context,
// then reads the payment ID from the request query. After validating the input,
// it calls the repository to delete the specified payment. 
func DeletePayment(c echo.Context) error {
	ac := c.(*middleware.AuthContext)
	ID, _, _ := ac.GetCurrentUser()

	var body struct {
		PaymentID uint `query:"payment_id" validate:"required"`
	}

	err := c.Bind(&body)

	if err != nil {
		return c.JSON(http.StatusBadGateway, util.SendMessage("INVALID_PARAMS"))
	}

	payment, err := repository.DeletePayment(ID, body.PaymentID)

	if err != nil {
		return c.JSON(http.StatusBadRequest, util.SendMessage(err.Error()))
	}

	result := make(map[string]int)
	result["rowsAffected"] = int(payment)
	return c.JSON(http.StatusOK, result)
}
