package main

import (
	"log"
	"net"
	"payment-service/src/database"
	"payment-service/src/handler"

	"google.golang.org/grpc"
)

func main() {
	grpcServer := grpc.NewServer()
	handler.NewHandler(grpcServer)

	// Initializing Database Connection
	database.ConnectDB()

	l, err := net.Listen("tcp", ":9000")

	if err != nil {
		log.Fatal(err.Error())
	}

	defer l.Close()

	if err := grpcServer.Serve(l); err != nil {
		log.Fatal(err.Error())
	}
}
