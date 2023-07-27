package repository

// a repository helps to interact with database
import (
	"errors"
	"fmt"

	"database/sql"

	_ "github.com/lib/pq"
	"github.com/raghavendrajha119/Ecommerce_website/middlewares"
	"github.com/raghavendrajha119/Ecommerce_website/models"
)

// Simulating a database call
func FindByCredentials(email, password string) (*models.User, error) {

	// Here you would query your database for the user with the given email

	//connecting with database
	const (
		host      = "localhost"
		port      = 5432
		user1     = "postgres"
		password1 = "Raghav@123"
		dbname    = "lib"
	)
	psqlconnect := fmt.Sprintf("host= %s port = %d user= %s password= %s dbname= %s sslmode=disable", host, port, user1, password1, dbname) //sslmode is disabled to send the data in plain text form
	db, err := sql.Open("postgres", psqlconnect)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	query := `SELECT id, email, password, name From users Where email = $1;` //here $1 i will replace with the actual parameter for a particular email
	row, err := db.Query(query, email)
	if err != nil {
		panic(err)
	}
	if !row.Next() { //checks if no rows of that query or email found
		return nil, errors.New("user not found")
	}
	var user models.User

	err = row.Scan(&user.ID, &user.Email, &user.Password, &user.Name)
	// & operator addresses the whole variable to initialize the user with datails
	if err != nil {
		return nil, err
	} //now row.scan() is used to read the id email and name from models to a local variable user

	passwordmatch := middlewares.ComparePasswords(password, user.Password)
	if passwordmatch {
		return &user, nil
	} else {
		return nil, errors.New("invalid password")
	}
	// here i am using a ComparePasswords function called in middleware which when called returns the true/false based on password
}
