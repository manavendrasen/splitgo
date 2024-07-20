package main

import (
	"auth-service/src/database"
	"auth-service/src/handler"
	"auth-service/src/middleware"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	// Initializing Database Connection
	database.ConnectDB()

	// Initializing Redis Connection
	database.ConnectCache()

	e.GET("/", handler.GetAppStatus)

	/*-------------------
					AUTH
	-------------------*/

	e.POST("/api/v1/auth/sign-up", handler.SignUp)
	e.POST("/api/v1/auth/login", handler.Login)
	e.POST("/api/v1/auth/refresh", handler.Refresh)
	e.POST("/api/v1/auth/logout", handler.Logout)

	/** -------------------
				API GATEWAY
	---------------------**/

	e.GET("/api/v1/payments", middleware.Auth(handler.GetPayments))
	e.POST("/api/v1/payments", middleware.Auth(handler.AddPayment))
	e.PATCH("/api/v1/payments", middleware.Auth(handler.UpdatePayment))
	e.DELETE("/api/v1/payments", middleware.Auth(handler.DeletePayment))

	e.Logger.Fatal(e.Start(":8080"))
}
