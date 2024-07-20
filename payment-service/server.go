package main

import (
	"payment-service/src/database"
	"payment-service/src/handler"
	"payment-service/src/middleware"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	// Initializing Database Connection
	database.ConnectDB()

	e.GET("/", handler.GetAppStatus)

	/*
		Payment Service
		- get all payments for user
		- add payment for user
		- edit payment details for user
		- delete payment for user
	*/

	e.GET("/api/v1/payment", middleware.Auth(handler.GetPayments))
	e.POST("/api/v1/payment", middleware.Auth(handler.AddPayment))
	e.PATCH("/api/v1/payment", middleware.Auth(handler.UpdatePayment))
	e.DELETE("/api/v1/payment", middleware.Auth(handler.DeletePayment))

	e.Logger.Fatal(e.Start(":8080"))

}
