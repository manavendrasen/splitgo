package main

import (
	"auth-service/src/database"
	"auth-service/src/handler"
	"auth-service/src/middleware"
	"log"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	// Initializing Database Connection
	database.ConnectDB()

	// Initializing Redis Connection
	database.ConnectCache()

	/*-------------------
					AUTH
	-------------------*/
	e.GET("/", handler.GetAppStatus)
	e.POST("/api/v1/auth/sign-up", handler.SignUp)
	e.POST("/api/v1/auth/login", handler.Login)
	e.POST("/api/v1/auth/refresh", handler.Refresh)
	e.POST("/api/v1/auth/logout", handler.Logout)

	/** -------------------
				API GATEWAY
	---------------------**/

	conn, err := grpc.NewClient(":9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Cannot Connect to Payment Service")
	}
	defer conn.Close()
	log.Print("API Gateway CONNECTED to Payment Service")

	paymentServiceHandler := handler.NewPaymentServiceHandler(conn)
	e.GET("/api/v1/payments", middleware.Auth(paymentServiceHandler.GetPayments))
	e.POST("/api/v1/payments", middleware.Auth(paymentServiceHandler.AddPayment))
	e.PATCH("/api/v1/payments", middleware.Auth(paymentServiceHandler.UpdatePayment))
	e.DELETE("/api/v1/payments/:paymentId", middleware.Auth(paymentServiceHandler.DeletePayment))

	e.Logger.Fatal(e.Start(":8080"))
}
