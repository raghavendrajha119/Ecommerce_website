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
	app.Listen(":3000")
}
