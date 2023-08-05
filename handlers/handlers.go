package handlers

// It handles the functionalities various render,redirect and response operations
import (
	"time"

	"github.com/gofiber/fiber/v2"
	jtoken "github.com/golang-jwt/jwt/v4"
	_ "github.com/lib/pq"
	"github.com/raghavendrajha119/Ecommerce_website/config"
	"github.com/raghavendrajha119/Ecommerce_website/middlewares"
	"github.com/raghavendrajha119/Ecommerce_website/models"
	"github.com/raghavendrajha119/Ecommerce_website/repository"
)

// Home page
func Home(c *fiber.Ctx) error {
	db, err := middlewares.OpenDB()
	if err != nil {
		panic(err)
	}
	var products []models.Product
	if err := db.Find(&products).Error; err != nil {
		return err
	}
	return c.JSON(products)
}

// Initial Login page
func LoginPg(c *fiber.Ctx) error {
	return c.Render("public/login.html", map[string]interface{}{"h1": "Log in here...."})
}

// Dashboard page
func Dashboard(c *fiber.Ctx) error {
	if tokenCookie := c.Cookies("jwt"); tokenCookie != "" {
		token, err := jtoken.Parse(tokenCookie, func(token *jtoken.Token) (interface{}, error) {
			return []byte(config.Secret), nil
		})
		if err == nil && token.Valid {
			claims := token.Claims.(jtoken.MapClaims)
			userName, okName := claims["name"].(string)
			userEmail, okEmail := claims["email"].(string)
			if okName && okEmail {
				return c.Render("public/dashboard.html", map[string]interface{}{
					"msg":   userName,
					"email": userEmail,
				})
			}
		}
	}
	return c.Redirect("/login")
}

// Login route
func Login(c *fiber.Ctx) error {
	// Extracting the credentials from the request body
	loginRequest := new(models.LoginRequest)
	if err := c.BodyParser(loginRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	// Getting the user by credentials
	user, err := repository.FindByCredentials(loginRequest.Email, loginRequest.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	day := time.Hour * 24
	claims := jtoken.MapClaims{
		"ID":    user.ID,
		"email": user.Email,
		"name":  user.Name,
		"exp":   time.Now().Add(day * 1).Unix(),
	}
	// Creating token
	token := jtoken.NewWithClaims(jtoken.SigningMethodHS256, claims)
	// Generating encoded token and send it as response.
	t, err := token.SignedString([]byte(config.Secret))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	// storing the token in cookies
	c.Cookie(&fiber.Cookie{
		Name:    "jwt",
		Value:   t,
		Expires: time.Now().Add(time.Hour * 24),
	})
	// Redirect the user to the desired page after successful login
	if user.Name != "" {
		return c.Redirect("/dashboard")
	} else {
		return c.JSON(fiber.Map{
			"msg": "Invalid",
		})
	}
}

// Logout route
func Logout(c *fiber.Ctx) error {
	// Clearing the authentication token cookie
	c.ClearCookie("jwt")
	// Redirect the user to the home page
	return c.Redirect("/")
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

// Initial user_register
func Register(c *fiber.Ctx) error {
	return c.Render("public/register.html", map[string]interface{}{
		"p": "Register Here"})
}

// would add the data of user inside the database
func RegisterPost(c *fiber.Ctx) error {
	db, err := middlewares.OpenDBUser()
	if err != nil {
		panic(err)
	}
	result := new(models.Register)
	c.BodyParser(result)

	hashpassword, err := middlewares.HashPassword(result.Password)
	if err != nil {
		return nil
	}

	users := models.User{
		Name:     result.Name,
		Email:    result.Email,
		Password: hashpassword,
	}
	if err := db.Create(&users).Error; err != nil {
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
func getUserIDFromJWT(c *fiber.Ctx) uint {
	tokenCookie := c.Cookies("jwt")
	if tokenCookie != "" {
		token, err := jtoken.Parse(tokenCookie, func(token *jtoken.Token) (interface{}, error) {
			return []byte(config.Secret), nil
		})
		if err == nil && token.Valid {
			claims := token.Claims.(jtoken.MapClaims)
			if userID, ok := claims["ID"].(float64); ok {
				return uint(userID)
			}
		}
	}
	return 0
}

func AddToCart(c *fiber.Ctx) error {
	userID := getUserIDFromJWT(c)

	// Parse the product ID from the request body
	request := struct {
		ProductID uint `json:"productId"`
	}{}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	// Call the AddToCart function to save the product in the cart
	if err := middlewares.AddToCart(userID, request.ProductID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to add product to cart"})
	}

	return c.SendStatus(fiber.StatusOK)
}
