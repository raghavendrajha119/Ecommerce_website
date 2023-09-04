package handlers

// It handles the functionalities various render,redirect and response operations
import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	jtoken "github.com/golang-jwt/jwt/v4"
	_ "github.com/lib/pq"
	"github.com/raghavendrajha119/Ecommerce_website/config"
	"github.com/raghavendrajha119/Ecommerce_website/middlewares"
	"github.com/raghavendrajha119/Ecommerce_website/models"
	"github.com/raghavendrajha119/Ecommerce_website/repository"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
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
			return []byte(config.GetSecret()), nil
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
func GoogleAuthStart(c *fiber.Ctx) error {
	oauthConf := &oauth2.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		RedirectURL:  os.Getenv("REDIRECT_URL"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}
	url := oauthConf.AuthCodeURL("state", oauth2.AccessTypeOffline)
	return c.Redirect(url)
}
func GoogleAuthCallback(c *fiber.Ctx) error {
	oauthConf := &oauth2.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		RedirectURL:  os.Getenv("REDIRECT_URL"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}

	code := c.Query("code")

	token, err := oauthConf.Exchange(c.Context(), code)
	if err != nil {
		return err
	}
	fmt.Println("Refesh :", token.RefreshToken, "Access :", token.AccessToken)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	profile, err := models.ConvertToken(token.AccessToken)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	_, err = repository.FindByGoogleAcc(profile.Email, profile.SUB)
	if err != nil {
		if err.Error() == "user not found" {

			db, err := middlewares.OpenDB()
			if err != nil {
				return c.JSON(err)
			}

			user := models.User{
				Email:    profile.Email,
				Password: profile.SUB,
				Name:     profile.GivenName,
				Role:     "user",
			}
			err = db.Create(&user).Error
			if err != nil {
				return c.JSON(err)
			}
		}
	}
	day := time.Hour * 24
	FindUserCredReg, err := repository.FindByGoogleAcc(profile.Email, profile.SUB)

	if err != nil {

		return c.JSON(err)
	}

	claims := jtoken.MapClaims{
		"ID":       FindUserCredReg.ID,
		"email":    FindUserCredReg.Email,
		"name":     FindUserCredReg.Name,
		"usertype": FindUserCredReg.Role,
		"exp":      time.Now().Add(day * 1).Unix(),
	}

	jwttoken := jtoken.NewWithClaims(jtoken.SigningMethodHS256, claims)

	t, err := jwttoken.SignedString([]byte(config.GetSecret()))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	c.Cookie(&fiber.Cookie{
		Name:    "jwt",
		Value:   t,
		Expires: time.Now().Add(24 * time.Hour),
	})
	if FindUserCredReg.Role == "user" {
		return c.Redirect("/dashboard.html")
	} else if FindUserCredReg.Role == "admin" {
		return c.Redirect("/admin/Dashboard.html")
	} else {
		return c.JSON(fiber.Map{
			"msg": "Invalid",
		})
	}

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
		"role":  user.Role,
		"exp":   time.Now().Add(day * 1).Unix(),
	}
	// Creating token
	token := jtoken.NewWithClaims(jtoken.SigningMethodHS256, claims)
	// Generating encoded token and send it as response.
	t, err := token.SignedString([]byte(config.GetSecret()))
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
	if user.Role == "user" {
		return c.Redirect("/dashboard")
	} else if user.Role == "admin" {
		return c.Redirect("/admin/Dashboard.html")
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
	return c.Redirect("Home.html")
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
	db, err := middlewares.OpenDB()
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
		Role:     "user",
	}
	if err := db.Create(&users).Error; err != nil {
		panic(err)
	}
	// Create the JWT token for the newly registered user (similar to the Login function)
	day := time.Hour * 24
	claims := jtoken.MapClaims{
		"ID":    users.ID,
		"email": users.Email,
		"name":  users.Name,
		"role":  users.Role,
		"exp":   time.Now().Add(day * 1).Unix(),
	}

	token := jtoken.NewWithClaims(jtoken.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(config.GetSecret()))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Storing the token in cookies
	c.Cookie(&fiber.Cookie{
		Name:    "jwt",
		Value:   t,
		Expires: time.Now().Add(time.Hour * 24),
	})

	// Redirect the user to the dashboard page after successful registration
	if users.Role == "user" {
		return c.Redirect("/dashboard")
	} else if users.Role == "admin" {
		return c.Redirect("/admin/Dashboard.html")
	} else {
		return c.JSON(fiber.Map{
			"msg": "Invalid",
		})
	}
}
func Categories(c *fiber.Ctx) error {
	db, err := middlewares.OpenDB()
	if err != nil {
		panic(err)
	}

	var categories []string
	if err := db.Model(&models.Product{}).Pluck("DISTINCT category", &categories).Error; err != nil {
		return err
	}

	return c.JSON(categories)
}
