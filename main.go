package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	jtoken "github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"github.com/raghavendrajha119/Ecommerce_website/config"
	"github.com/raghavendrajha119/Ecommerce_website/handlers"
	"github.com/raghavendrajha119/Ecommerce_website/middlewares"
)

func main() {
	//loading env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	app := fiber.New()
	jwt := middlewares.NewAuthMiddleware(config.GetSecret())
	// Default route
	app.Static("/", "./public", fiber.Static{Index: "Home.html"})
	app.Get("/home", handlers.Home)
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
	//admin routes
	admin := app.Group("/admin")
	admin.Use(func(c *fiber.Ctx) error {
		if tokenCookie := c.Cookies("jwt"); tokenCookie != "" {
			token, err := jtoken.Parse(tokenCookie, func(token *jtoken.Token) (interface{}, error) {
				return []byte(config.GetSecret()), nil
			})
			if err == nil && token.Valid {
				claims := token.Claims.(jtoken.MapClaims)
				if roleClaim, ok := claims["role"].(string); ok {
					if roleClaim != "admin" {
						return c.SendString("Unauthorized")
					}
				} else {
					return c.SendString("Invalid role claim")
				}
			}
		}
		return c.Next()
	})
	admin.Get("/", handlers.AdminDashboard)
	admin.Get("/products", handlers.AdminProducts)
	admin.Post("/add-products", handlers.AdminaddProducts)
	// Listen on port 3000
	app.Listen(":3000")
}
