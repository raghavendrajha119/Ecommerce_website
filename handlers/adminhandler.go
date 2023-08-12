package handlers

import (
	"path/filepath"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/raghavendrajha119/Ecommerce_website/middlewares"
	"github.com/raghavendrajha119/Ecommerce_website/models"
)

func AdminDashboard(c *fiber.Ctx) error {
	db, err := middlewares.OpenDB()
	if err != nil {
		return err
	}
	totalUsers, err := middlewares.GetTotalUsers(db)
	if err != nil {
		return err
	}
	totalAdminUsers, err := middlewares.GetTotalAdminUsers(db)
	if err != nil {
		return err
	}
	totalProducts, err := middlewares.GetTotalProducts(db)
	if err != nil {
		return err
	}
	data := fiber.Map{
		"totalUsers":      totalUsers,
		"totalAdminUsers": totalAdminUsers,
		"totalProducts":   totalProducts,
	}
	return c.JSON(data)
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

// user list
func AdminGetUsers(c *fiber.Ctx) error {
	db, err := middlewares.OpenDB()
	if err != nil {
		return err
	}

	var users []models.User
	if err := db.Find(&users).Error; err != nil {
		return err
	}

	return c.JSON(users)
}

//handling admin-create and remove

func AdminMakeAdmin(c *fiber.Ctx) error {
	userID, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return err
	}

	db, err := middlewares.OpenDB()
	if err != nil {
		return err
	}

	if err := middlewares.UpdateUserRole(db, uint(userID), "admin"); err != nil {
		return err
	}

	return c.SendString("User role updated to admin")
}

func AdminRemoveAdmin(c *fiber.Ctx) error {
	userID, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return err
	}

	db, err := middlewares.OpenDB()
	if err != nil {
		return err
	}

	if err := middlewares.UpdateUserRole(db, uint(userID), "user"); err != nil {
		return err
	}

	return c.SendString("User role updated to user")
}
