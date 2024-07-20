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


	e.Logger.Fatal(e.Start(":8080"))

}
