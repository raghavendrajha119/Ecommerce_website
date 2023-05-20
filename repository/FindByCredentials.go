package repository

import (
	"errors"
	"fmt"

	"database/sql"

	_ "github.com/lib/pq"
	"github.com/raghavendrajha119/Ecommerce_website/models"
)

// Simulate a database call
func FindByCredentials(email, password string) (*models.User, error) {

	// Here you would query your database for the user with the given email

	//connecting with database
	const (
		host      = "localhost"
		port      = 5432
		user      = "postgres"
		password1 = "Raghav@123"
		dbname    = "lib"
	)
	psqlconnect := fmt.Sprintf("host= %s port = %d user= %s password= %s dbname= %s sslmode=disable", host, port, user, password1, dbname)
	db, err := sql.Open("postgres", psqlconnect)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	query := `SELECT id, email, password, name From users Where email = $1 AND password = $2;`
	row, err := db.Query(query, email, password)
	if err != nil {
		panic(err)
	}
	if row.Next() {
		var user models.User
		err = row.Scan(&user.ID, &user.Email, &user.Password, &user.Name)
		if err != nil {
			return nil, err
		}
		return &user, nil
	} else {
		return nil, errors.New("user not found")
	}
}
