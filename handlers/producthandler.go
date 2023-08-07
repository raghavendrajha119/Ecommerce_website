package handlers

import (
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
