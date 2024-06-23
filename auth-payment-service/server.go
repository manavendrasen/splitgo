package main

import (
	"net/http"
	"payment-service/database"
	"payment-service/handler"
	"payment-service/middleware"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(echoMiddleware.LoggerWithConfig(echoMiddleware.LoggerConfig{
		Format:           `${time_rfc3339} >> [${status}][${method}] ${uri} ${error} (${latency_human})` + "\n",
		CustomTimeFormat: "2006-01-02 15:04:05.00000",
	}))

	e.GET("/", func(c echo.Context) error {
		status := make(map[string]string, 1)
		status["Status"] = "Up"
		return c.JSON(http.StatusOK, status)
	})

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

	e.POST("/sign-up", handler.SignUp)
	e.POST("/login", handler.Login)

	// Initializing Database Connection
	database.Connect()

	e.Logger.Fatal(e.Start(":8080"))

}
