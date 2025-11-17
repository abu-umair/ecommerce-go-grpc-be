package main

import (
	"context"
	"log"
	"net"
	"os"
	"time"

	"github.com/abu-umair/ecommerce-go-grpc-be/internal/grpcmiddleware"
	"github.com/abu-umair/ecommerce-go-grpc-be/internal/handler"
	"github.com/abu-umair/ecommerce-go-grpc-be/internal/repository"
	"github.com/abu-umair/ecommerce-go-grpc-be/internal/service"
	"github.com/abu-umair/ecommerce-go-grpc-be/pb/auth"
	"github.com/abu-umair/ecommerce-go-grpc-be/pb/cart"
	"github.com/abu-umair/ecommerce-go-grpc-be/pb/order"
	"github.com/abu-umair/ecommerce-go-grpc-be/pb/product"
	"github.com/abu-umair/ecommerce-go-grpc-be/pkg/database"
	"github.com/joho/godotenv"
	gocache "github.com/patrickmn/go-cache"
	"github.com/xendit/xendit-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() { //?Sebagai gRpc server
	ctx := context.Background()
	godotenv.Load()

	xendit.Opt.SecretKey = os.Getenv("XENDIT_SECRET")

	lis, err := net.Listen("tcp", ":50052") //? mengembalikan 2 value yaitu listener dan error
	if err != nil {
		log.Panicf("Error when listening %v", err)
	}

	db := database.ConnectDB(ctx, os.Getenv("DB_URI"))
	log.Println("Database is connected")

	cacheService := gocache.New(time.Hour*24, time.Hour*24)

	authMiddleware := grpcmiddleware.NewAuthMiddleware(cacheService)

	authRepository := repository.NewAuthRepository(db)
	authService := service.NewAuthService(authRepository, cacheService)
	authHandler := handler.NewAuthHandler(authService)

	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepository)
	productHandler := handler.NewProductHandler(productService)

	cartRepository := repository.NewCartRepository(db)
	cartService := service.NewCartService(productRepository, cartRepository)
	cartHandler := handler.NewCartHandler(cartService)

	orderRepository := repository.NewOrderRepository(db)
	orderService := service.NewOrderService(db, orderRepository, productRepository)
	orderHandler := handler.NewOrderHandler(orderService)

	serv := grpc.NewServer(

		grpc.ChainUnaryInterceptor(
			grpcmiddleware.ErrorMiddleware, //?memasukkan func error middleware
			authMiddleware.Middleware,
		),
	)

	auth.RegisterAuthServiceServer(serv, authHandler)
	product.RegisterProductServiceServer(serv, productHandler)
	cart.RegisterCartServiceServer(serv, cartHandler)
	order.RegisterOrderServiceServer(serv, orderHandler)

	if os.Getenv("ENVIRONMENT") == "dev" {
		reflection.Register(serv)
		log.Println("Reflection is registered")
	}

	log.Println("Server is running on :50052 port")
	if err := serv.Serve(lis); err != nil {
		log.Panicf("Server is error %v", err)
	}
}
