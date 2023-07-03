package handlers

import (
	"database/sql"
	"fmt"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"github.com/raghavendrajha119/Ecommerce_website/models"
)

// AddToCart adds a product to the user's cart
func AddToCart(c *fiber.Ctx) error {
	request := new(models.AddToCartRequest)
	if err := c.BodyParser(request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err := storeProductInCart(request.UserID, request.ProductID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to add product to cart"})
	}

	return c.SendStatus(fiber.StatusOK)
}

// storeProductInCart stores the product ID in the cart
func storeProductInCart(userID, productID int) error {
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "Raghav@123"
		dbname   = "lib"
	)
	psqlconnect := fmt.Sprintf("host= %s port = %d user= %s password= %s dbname= %s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconnect)
	if err != nil {
		return err
	}
	defer db.Close()

	// Execute the SQL query to store the product ID
	_, err = db.Exec("INSERT INTO product (user_id, product_id) VALUES ($1, $2)", userID, productID)
	if err != nil {
		return err
	}

	return nil
}

// GetCartProducts fetches the products from the user's cart
func GetCartProducts(c *fiber.Ctx) error {
	// Fetch the product IDs from the database
	// You can modify this code based on your database structure and query
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "Raghav@123"
		dbname   = "lib"
	)
	psqlconnect := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconnect)
	if err != nil {
		return err
	}
	defer db.Close()

	rows, err := db.Query("SELECT product_id FROM product")
	if err != nil {
		return err
	}
	defer rows.Close()

	var productIDs []int
	for rows.Next() {
		var productID int
		err := rows.Scan(&productID)
		if err != nil {
			return err
		}
		productIDs = append(productIDs, productID)
	}

	// Return the product IDs as JSON response
	return c.JSON(fiber.Map{"productIDs": productIDs})
}
