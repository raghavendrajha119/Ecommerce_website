package middlewares

// Middlewares is created basically to manage the data from users like authentication and other functionalities
import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/raghavendrajha119/Ecommerce_website/config"
	"golang.org/x/crypto/bcrypt"
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
