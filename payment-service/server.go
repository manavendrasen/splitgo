package main

import (
	"net/http"
	"payment-service/database"
	"payment-service/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	
	// Initializing Database Connection
  database.Init()

	// Middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `${time_rfc3339} >> [${status}][${method}] ${uri} ${error} (${latency_human})` + "\n",
		CustomTimeFormat: "2006-01-02 15:04:05.00000",
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.BodyLimit("2M"))

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

	e.GET("/payment", handler.GetPayments)
	e.POST("/payment", handler.AddPayment)
	e.PATCH("/payment/:id", handler.UpdatePayment)
	e.DELETE("/payment/:id", handler.DeletePayment)
	

	e.Logger.Fatal(e.Start(":8080"))
}

