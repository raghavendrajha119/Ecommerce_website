package repository

import (
	"errors"

	"database/sql"

	"github.com/raghavendrajha119/Ecommerce_website/models"
)

// Simulate a database call
func FindByCredentials(db *sql.DB, email, password string) (*models.User, error) {
	// Here you would query your database for the user with the given email
	query := "SELECT * From users Where email = ? AND password = ?"
	row, err := db.QueryRow(query, email, password)
	if err != nil {
		panic(err)
	}
	if row.Next() {
		var user models.User
		err = row.Scan(&user.ID, &user.Email, &user.Password, &user.FavouritePhrase)
		if err != nil {
			return nil, err
		}
		return &user, nil
	} else {
		return nil, errors.New("user not found")
	}

}
