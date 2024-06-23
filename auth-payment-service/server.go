package main

import (
	"payment-service/database"
	"payment-service/handler"
	"payment-service/middleware"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	
	// Initializing Database Connection
	database.Connect()

	e.GET("/", handler.GetAppStatus)

	/*
		Auth  Service
		- sign up
		- login
	*/

	e.POST("/sign-up", handler.SignUp)
	e.POST("/login", handler.Login)

	/*
		Payment Service
		- get all payments for user
		- add payment for user
		- edit payment details for user
		- delete payment for user
	*/

	e.GET("/payment", middleware.Auth(handler.GetPayments))
	e.POST("/payment", middleware.Auth(handler.AddPayment))
	e.PATCH("/payment", middleware.Auth(handler.UpdatePayment))
	e.DELETE("/payment", middleware.Auth(handler.DeletePayment))

	

	e.Logger.Fatal(e.Start(":8080"))

}
