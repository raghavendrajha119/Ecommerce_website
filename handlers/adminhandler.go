package handlers

import (
	"path/filepath"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/raghavendrajha119/Ecommerce_website/middlewares"
	"github.com/raghavendrajha119/Ecommerce_website/models"
)

func AdminDashboard(c *fiber.Ctx) error {
	return c.Render("public/admin/Dashboard.html", nil)
}
func AdminProducts(c *fiber.Ctx) error {
	db, err := middlewares.OpenDB()
	if err != nil {
		panic(err)
	}
	var product []models.Product
	if err := db.Find(&product).Error; err != nil {
		return err
	}
	c.Set("Content-Type", "application/json")
	return c.JSON(product)
}

// Admin adds the products into the db
func AdminaddProducts(c *fiber.Ctx) error {
	title := c.FormValue("title")
	price, _ := strconv.ParseFloat(c.FormValue("price"), 64)
	description := c.FormValue("description")
	category := c.FormValue("category")

	// Handle image upload
	file, err := c.FormFile("image")
	if err != nil {
		return err
	}
	// Generate image filename based on the product title
	imageFilename := generateImageFilename(title, filepath.Ext(file.Filename))
	imagePath := "public/img/" + imageFilename

	if err := c.SaveFile(file, imagePath); err != nil {
		return err
	}

	// Create a new product in the database
	product := models.Product{
		Title:       title,
		Price:       price,
		Description: description,
		Category:    category,
		Image:       imagePath,
	}
	db, err := middlewares.OpenDB()
	if err != nil {
		return err
	}
	if err := db.Create(&product).Error; err != nil {
		return err
	}

	return c.Redirect("/admin/products")
}
func generateImageFilename(productTitle, ext string) string {
	return productTitle + ext
}
