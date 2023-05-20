package handlers

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	jtoken "github.com/golang-jwt/jwt/v4"
	_ "github.com/lib/pq"
	"github.com/raghavendrajha119/Ecommerce_website/config"
	"github.com/raghavendrajha119/Ecommerce_website/models"
	"github.com/raghavendrajha119/Ecommerce_website/repository"
)

// Login route
func Login(c *fiber.Ctx) error {
	// Extract the credentials from the request body
	loginRequest := new(models.LoginRequest)
	if err := c.BodyParser(loginRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	// Find the user by credentials
	user, err := repository.FindByCredentials(loginRequest.Email, loginRequest.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	day := time.Hour * 24
	// Create the JWT claims, which includes the user ID and expiry time
	claims := jtoken.MapClaims{
		"ID":    user.ID,
		"email": user.Email,
		"name":  user.Name,
		"exp":   time.Now().Add(day * 1).Unix(),
	}
	// Create token
	token := jtoken.NewWithClaims(jtoken.SigningMethodHS256, claims)
	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(config.Secret))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	// Return the token
	return c.JSON(models.LoginResponse{
		Token: t,
	})
}

// Protected route
func Protected(c *fiber.Ctx) error {
	// Get the user from the context and return it
	user := c.Locals("user").(*jtoken.Token)
	claims := user.Claims.(jtoken.MapClaims)
	email := claims["email"].(string)
	Name := claims["fav"].(string)
	return c.SendString("Welcome 👋" + email + " " + Name)
}

// Initial Login page
func LoginPg(c *fiber.Ctx) error {
	return c.Render("public/login.html", map[string]interface{}{"h1": "Log in here...."})
}
func LoginPack(c *fiber.Ctx) error {
	//collecting data
	loginRequest := new(models.LoginRequest)
	if err := c.BodyParser(loginRequest); err != nil {
		return err
	}
	row, err := repository.FindByCredentials(loginRequest.Email, loginRequest.Password)
	if err != nil {
		panic(err)
	}

	return c.Render("public/afterlogin.html", map[string]interface{}{"msg": "Welcome to Go Shopping " + row.Name + " Continue Shopping......"})
}

// Initial user_register
func Register(c *fiber.Ctx) error {
	return c.Render("public/register.html", map[string]interface{}{
		"p": "Register Here"})
}

// would add the data of user inside the database
func RegisterPost(c *fiber.Ctx) error {
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
		panic(err)
	}
	result := new(models.Register)
	c.BodyParser(result)
	data, err := db.Prepare("INSERT INTO users (email,password,name) VALUES ($1,$2,$3)")
	if err != nil {
		panic(err)
	}
	defer data.Close()
	_, err = data.Exec(result.Email, result.Password, result.Name)
	if err != nil {
		panic(err)
	}
	c.Redirect("/registered")
	return nil
}

// After successful Registration
func RegisterSuccessful(c *fiber.Ctx) error {

	return c.Render("public/registrationsuccessful.html", map[string]interface{}{
		"msg": "Welcome to Go Shopping Continue Shopping......"})
}