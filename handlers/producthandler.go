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
	var products []models.Product
	if err := db.Find(&products).Error; err != nil {
		return err
	}
	return c.JSON(products)
}
