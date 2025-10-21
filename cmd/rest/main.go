package main

import (
	"log"
	"mime"
	"net/http"
	"os"
	"path"

	"github.com/abu-umair/ecommerce-go-grpc-be/internal/handler"
	"github.com/gofiber/fiber/v2"
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
	app := fiber.New()

	app.Get("/storage/product/:filename", handleGetFileName)

	app.Post("/product/upload", handler.UploadProductImageHandler)

	app.Listen(":4000")

}
