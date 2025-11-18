package main

import (
	"context"
	"log"
	"mime"
	"net/http"
	"os"
	"path"

	"github.com/abu-umair/ecommerce-go-grpc-be/internal/handler"
	"github.com/abu-umair/ecommerce-go-grpc-be/internal/repository"
	"github.com/abu-umair/ecommerce-go-grpc-be/internal/service"
	"github.com/abu-umair/ecommerce-go-grpc-be/pkg/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func handleGetFileName(c *fiber.Ctx) error {
	fileNameParam := c.Params("filename")
	filePath := path.Join("storage", "product", fileNameParam)
	if _, err := os.Stat(filePath); err != nil {
		if os.IsNotExist(err) {
			return c.Status(http.StatusNotFound).SendString("Not Found")
		}
		log.Println(err)
		return c.Status(http.StatusInternalServerError).SendString("Internal Server Error")
	}

	//? membuka file
	file, err := os.Open(filePath)
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusInternalServerError).SendString("Internal Server Error")
	}

	//? mengirim file sbg response
	ext := path.Ext(filePath)
	mimeType := mime.TypeByExtension(ext)

	c.Set("Content-Type", mimeType) //?konversi agar tampilan gambar sesuai (dinamis)
	// c.Set("Content-Type", "image/png") //?konversi agar tampilan gambar sesuai (blm dinamis)
	return c.SendStream(file)
}

func main() {
	godotenv.Load()
	ctx := context.Background()
	app := fiber.New()

	db := database.ConnectDB(ctx, os.Getenv("DB_URI"))
	orderRepository := repository.NewOrderRepository(db)
	webhookService := service.NewWebhookService(orderRepository)
	webhookHandler := handler.NewWebhookHandler(webhookService)

	app.Use(cors.New())

	app.Get("/storage/product/:filename", handleGetFileName)

	app.Post("/product/upload", handler.UploadProductImageHandler)

	app.Post("/webhook/xendit/invoice", webhookHandler.ReceiveInvoice) //? endpoint webhook

	app.Listen(":4000")

}
