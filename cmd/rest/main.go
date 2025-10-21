package main

import (
	"github.com/abu-umair/ecommerce-go-grpc-be/internal/handler"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Post("/product/upload", handler.UploadProductImageHandler)

	app.Listen(":3000")

}
