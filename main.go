package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/raghavendrajha119/Ecommerce_website/config"
	"github.com/raghavendrajha119/Ecommerce_website/handlers"
	"github.com/raghavendrajha119/Ecommerce_website/middlewares"
)

func main() {
	// Create a new Fiber instance
	app := fiber.New()
	// Create a new JWT middleware
	// Note: This is just an example, please use a secure secret key
	jwt := middlewares.NewAuthMiddleware(config.Secret)
	// Create a Login route
	app.Static("/", "./public")
	app.Get("/", handlers.Home)
	app.Post("/login", handlers.Login)
	// Create a protected route
	app.Get("/protected", jwt, handlers.Protected)
	//registration page
	app.Get("/register", handlers.Register)
	app.Post("/register", handlers.RegisterPost)
	app.Get("/registered", handlers.RegisterSuccessful)
	app.Get("/loginpg", handlers.LoginPg)
	app.Post("/loginpg", handlers.LoginPack)
	//logout
	app.Get("/logout", handlers.Logout)
	// Listen on port 3000
	// Register routes
	//admin routes
	app.Get("/admin/login", handlers.AdminLoginPage)
	app.Post("/admin/login", handlers.AdminLogin) // Admin login route
	// Middleware to protect admin routes
	app.Get("/admin", handlers.AdminPage)               // Admin page route
	app.Get("/admin/orders", handlers.OrdersPage)       // Admin orders page route
	app.Get("/admin/products", handlers.ProductsPage)   // Admin products page route
	app.Get("/admin/customers", handlers.CustomersPage) // Admin customers page route
	//cart routes
	app.Post("/add-to-cart", handlers.AddToCart)
	app.Get("/cart-products", handlers.GetCartProducts)
	app.Listen(":3000")
}
