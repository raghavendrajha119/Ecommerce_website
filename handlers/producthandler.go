package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/raghavendrajha119/Ecommerce_website/middlewares"
	"github.com/raghavendrajha119/Ecommerce_website/models"
)

func ProductHandler(c *fiber.Ctx) error {
	db, err := middlewares.OpenDB()
	if err != nil {
		panic(err)
	}
	productId := c.Query("id")
	var product models.Product
	if err := db.Find(&product, productId).Error; err != nil {
		return err
	}
	return c.JSON(product)
}

// Fetch similar products based on title, description, and category
func SimilarProducts(c *fiber.Ctx) error {
	db, err := middlewares.OpenDB()
	if err != nil {
		panic(err)
	}
	title := c.Query("title")
	description := c.Query("description")
	category := c.Query("category")

	var similarProducts []models.Product
	if err := db.Where("Title LIKE ? OR Description LIKE ? OR Category = ?", "%"+title+"%", "%"+description+"%", category).Find(&similarProducts).Error; err != nil {
		return err
	}
	fmt.Println(similarProducts)
	return c.JSON(similarProducts)
}
