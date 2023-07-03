package handlers

import (
	"database/sql"
	"fmt"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"github.com/raghavendrajha119/Ecommerce_website/models"
)

// GetCustomers retrieves the details of all customers from the database
func GetCustomers(c *fiber.Ctx) error {
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "Raghav@123"
		dbname   = "lib"
	)

	// Establish a connection to the database
	psqlconnect := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconnect)
	if err != nil {
		return err
	}
	defer db.Close()

	// Query the customers table to fetch all customer details
	rows, err := db.Query("SELECT id, name, email FROM users")
	if err != nil {
		return err
	}
	defer rows.Close()

	var customers []models.Customer

	// Iterate over the query results and populate the customers slice
	for rows.Next() {
		var customer models.Customer
		err := rows.Scan(&customer.ID, &customer.Name, &customer.Email)
		if err != nil {
			return err
		}
		customers = append(customers, customer)
	}

	// Return the customers as a JSON response
	return c.JSON(customers)
}
