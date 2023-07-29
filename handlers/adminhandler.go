package handlers

import (
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func AdminPage(c *fiber.Ctx) error {
	return c.Render("public/admin.html", map[string]interface{}{})
}
func AdminLoginPage(c *fiber.Ctx) error {
	return c.Render("public/adminlogin.html", map[string]interface{}{})
}

func AdminLogin(c *fiber.Ctx) error {

	return c.Redirect("/admin")
}

func AdminProtected(c *fiber.Ctx) error {
	return nil
}

func OrdersPage(c *fiber.Ctx) error {
	return c.Render("public/orders.html", map[string]interface{}{})
}
func ProductsPage(c *fiber.Ctx) error {
	return c.Render("public/products.html", map[string]interface{}{})
}

func CustomersPage(c *fiber.Ctx) error {
	return c.Render("public/customers.html", map[string]interface{}{})
}
