package handler

import (
	"auth-service/src/middleware"
	"auth-service/src/util"
	"net/http"
	"strconv"
	"strings"

	pb "common"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
)

type handler struct {
	client pb.PaymentServiceClient
}

func NewPaymentServiceHandler(conn *grpc.ClientConn) *handler {
	return &handler{
		client: pb.NewPaymentServiceClient(conn),
	}
}

func (h *handler) GetPayments(c echo.Context) error {
	ac := c.(*middleware.AuthContext)
	ID, _, _ := ac.GetCurrentUser()

	payments, err := h.client.GetPayment(c.Request().Context(), &pb.GetPaymentRequest{
		From: uint64(ID),
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, util.SendMessage(err.Error()))
	}

	return c.JSON(http.StatusOK, payments)
}

func (h *handler) AddPayment(c echo.Context) error {
	ac := c.(*middleware.AuthContext)
	ID, _, _ := ac.GetCurrentUser()

	var body struct {
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

	payment, err := h.client.AddPayment(c.Request().Context(), &pb.AddPaymentRequest{
		From:        uint64(ID),
		To:          body.To,
		Description: body.Description,
		Amount:      body.Amount,
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, util.SendMessage(err.Error()))
	}

	return c.JSON(http.StatusCreated, payment)
}

func (h *handler) UpdatePayment(c echo.Context) error {
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

	payment, err := h.client.UpdatePayment(c.Request().Context(), &pb.UpdatePaymentRequest{
		From:        uint64(ID),
		To:          body.To,
		PaymentId:   uint64(body.PaymentID),
		Description: body.Description,
		Amount:      body.Amount,
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, util.SendMessage(err.Error()))
	}

	return c.JSON(http.StatusCreated, payment)
}

func (h *handler) DeletePayment(c echo.Context) error {
	ac := c.(*middleware.AuthContext)
	ID, _, _ := ac.GetCurrentUser()

	paymentId, err := strconv.Atoi(c.Param("paymentId"))

	if err != nil {
		return c.JSON(http.StatusBadGateway, util.SendMessage("INVALID_ID"))
	}

	payment, err := h.client.DeletePayment(c.Request().Context(), &pb.DeletePaymentRequest{
		UserId:    uint64(ID),
		PaymentId: uint64(paymentId),
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, util.SendMessage(err.Error()))
	}

	return c.JSON(http.StatusCreated, payment)
}
