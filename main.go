package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/raghavendrajha119/Ecommerce_website/config"
	"github.com/raghavendrajha119/Ecommerce_website/handlers"
	"github.com/raghavendrajha119/Ecommerce_website/middlewares"
)

func main() {
	app := fiber.New()
	jwt := middlewares.NewAuthMiddleware(config.Secret)
	// Default route
	app.Static("/", "./public")
	app.Get("/", handlers.Home)
	// protected route
	app.Get("/protected", jwt, handlers.Protected)
	//registration page
	// Register routes
	app.Get("/register", handlers.Register)
	app.Post("/register", handlers.RegisterPost)
	app.Get("/registered", handlers.RegisterSuccessful)
	app.Get("/login", handlers.LoginPg)
	app.Post("/login", handlers.Login)
	app.Get("/dashboard", handlers.Dashboard)
	//logout
	app.Get("/logout", handlers.Logout)
	//admin routes
	app.Get("/admin/login", handlers.AdminLoginPage)
	app.Post("/admin/login", handlers.AdminLogin)
	app.Get("/admin", handlers.AdminPage)
	app.Get("/admin/orders", handlers.OrdersPage)
	app.Get("/admin/products", handlers.ProductsPage)
	app.Get("/admin/customers", handlers.CustomersPage)
	//cart routes
	app.Post("/add-to-cart", handlers.AddToCart)
	app.Get("/cart-products", handlers.GetCartProducts)
	// Listen on port 3000
	app.Listen(":3000")
}
