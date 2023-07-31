package middlewares

// Middlewares is created basically to manage the data from users like authentication and other functionalities
import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/raghavendrajha119/Ecommerce_website/config"
	"github.com/raghavendrajha119/Ecommerce_website/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Middleware JWT function checks for valid jwt token
func NewAuthMiddleware(secret string) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(secret),
	})
}

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func ComparePasswords(hashedPwd, plainPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(plainPwd), []byte(hashedPwd))
	return err == nil
}

func CheckJWT(c *fiber.Ctx, cookieName string) error {
	return jwtware.New(jwtware.Config{
		SigningKey:  []byte(config.Secret),
		TokenLookup: fmt.Sprintf("cookie:%s", cookieName),
	})(c)
}

// gorm connection syntax
func OpenDBUser() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=Raghav@123 dbname=lib port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&models.User{})
	return db, nil
}
func OpenDB() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=Raghav@123 dbname=lib port=5432 sslmode=disable"
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
