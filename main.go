package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
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
	//loading env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// protected route
	app.Get("/protected", jwt, handlers.Protected)
	//registration page
	// Register routes
	app.Get("/register", handlers.Register)
	app.Post("/register", handlers.RegisterPost)
	app.Get("/login", handlers.LoginPg)
	app.Post("/login", handlers.Login)
	app.Get("/dashboard", handlers.Dashboard)
	//logout
	app.Get("/logout", handlers.Logout)
	//product route
	app.Get("/products", handlers.ProductHandler)
	//cart route
	app.Post("/add-to-cart", handlers.AddtoCart)
	app.Get("/get-cart", handlers.GetfromCart)
	// Listen on port 3000
	app.Listen(":3000")
}
