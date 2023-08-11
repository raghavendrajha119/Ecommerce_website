package config

// The secret key used to sign the JWT, this is to be a secure key and should not be stored in the code
import (
	"os"
)

func GetSecret() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		panic("JWT_SECRET environment variable not set")
	}
	return secret
}
