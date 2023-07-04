package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	jtoken "github.com/golang-jwt/jwt/v4"
	_ "github.com/lib/pq"
	"github.com/raghavendrajha119/Ecommerce_website/config"
	"github.com/raghavendrajha119/Ecommerce_website/middlewares"
	"github.com/raghavendrajha119/Ecommerce_website/models"
)

func AdminLoginPage(c *fiber.Ctx) error {
	return c.Render("public/adminlogin.html", map[string]interface{}{})
}

// AdminLogin handles the admin login route
func AdminLogin(c *fiber.Ctx) error {
	// Extract the credentials from the request body
	loginRequest := new(models.AdminLoginRequest)
	if err := c.BodyParser(loginRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Check if the admin credentials are valid
	if loginRequest.Username != "admin" || loginRequest.Password != "adminpassword" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid admin credentials"})
	}

	// Create the JWT claims, which includes the admin ID and expiry time
	claims := jtoken.MapClaims{
		"ID":       1, // You can set the actual admin ID here
		"username": loginRequest.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	// Create token
	token := jtoken.NewWithClaims(jtoken.SigningMethodHS256, claims)

	// Generate encoded token and send it as response
	t, err := token.SignedString([]byte(config.Secret))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Store the token in cookies
	c.Cookie(&fiber.Cookie{
		Name:     "admin_jwt",
		Value:    t,
		Expires:  time.Now().Add(time.Hour * 24),
		Secure:   true,
		HTTPOnly: true,
	})

	// Redirect the admin to the admin page after successful login
	return c.Redirect("/admin")
}

// AdminProtected is a middleware that ensures the route can only be accessed by authenticated admin users
func AdminProtected(c *fiber.Ctx) error {
	// Verify the admin JWT token
	err := middlewares.CheckJWT(c, "admin_jwt")
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	// Continue to the next handler if the token is valid
	return c.Next()
}

// AdminPage handles the admin page route
func AdminPage(c *fiber.Ctx) error {
	return c.Render("public/admin.html", map[string]interface{}{})
}

// OrdersPage handles the admin orders page route
func OrdersPage(c *fiber.Ctx) error {
	// Implement logic for fetching and rendering the admin orders page
	// You can retrieve order data from the database and pass it to the template

	return c.Render("public/orders.html", map[string]interface{}{})
}

// ProductsPage handles the admin products page route
func ProductsPage(c *fiber.Ctx) error {
	// Implement logic for fetching and rendering the admin products page
	// You can retrieve product data from the database and pass it to the template

	return c.Render("public/products.html", map[string]interface{}{})
}

// CustomersPage handles the
func CustomersPage(c *fiber.Ctx) error {
	// Implement logic for fetching and rendering the admin customers page
	// You can retrieve customer data from the database and pass it to the template

	return c.Render("public/customers.html", map[string]interface{}{})
}
