package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/abu-umair/ecommerce-go-grpc-be/internal/handler"
	"github.com/abu-umair/ecommerce-go-grpc-be/pb/service"
	"github.com/abu-umair/ecommerce-go-grpc-be/pkg/database"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() { //?Sebagai gRpc server
	ctx := context.Background()
	godotenv.Load()

	lis, err := net.Listen("tcp", ":50051") //? mengembalikan 2 value yaitu listener dan error
	if err != nil {
		log.Panicf("Error when listening %v", err)
	}

	database.ConnectDB(ctx, os.Getenv("DB_URI"))
	log.Println("Database is connected")

	serviceHandler := handler.NewServiceHandler()

	serv := grpc.NewServer()

	service.RegisterHelloWorldServiceServer(serv, serviceHandler)

	if os.Getenv("ENVIRONMENT") == "dev" {
		reflection.Register(serv)
		log.Println("Reflection is registered")
	}

	log.Println("Server is running on :50051 port")
	if err := serv.Serve(lis); err != nil {
		log.Panicf("Server is error %v", err)
	}
}
