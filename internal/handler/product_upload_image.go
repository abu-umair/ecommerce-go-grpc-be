package handler

import (
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gofiber/fiber/v2"
)

func UploadProductImageHandler(c *fiber.Ctx) error {
	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "image data not found",
		})
	}

	// product_1623232.png (membuat format imagge name)
	timestamp := time.Now().UnixNano()
	fileName := fmt.Sprintf("product_%d%s", timestamp, filepath.Ext(file.Filename))
	uploadPath := "./storage/product/" + fileName
	// c.SaveFile(file, "./storage/product/product.jpeg")
	err = c.SaveFile(file, uploadPath)

	//return error jika ada
	if err != nil {
		fmt.Println(err)

		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "internal server error",
		})
	}

	return c.JSON(fiber.Map{
		"success":   true,
		"message":   "Upload success",
		"file_name": fileName,
	})
}
