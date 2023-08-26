package repository

// a repository helps to interact with database
import (
	"errors"

	_ "github.com/lib/pq"
	"github.com/raghavendrajha119/Ecommerce_website/middlewares"
	"github.com/raghavendrajha119/Ecommerce_website/models"
	"gorm.io/gorm"
)

func FindByCredentials(email, password string) (*models.User, error) {
	//connecting with database
	db, err := middlewares.OpenDB()
	if err != nil {
		panic(err)
	}
	var user models.User
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	passwordmatch := middlewares.ComparePasswords(password, user.Password)
	if passwordmatch {
		return &user, nil
	} else {
		return nil, errors.New("invalid password")
	}
	// here i am using a ComparePasswords function called in middleware which when called returns the true/false based on password
}
func FindByGoogleAcc(email, password string) (*models.User, error) {
	db, err := middlewares.OpenDB()
	if err != nil {
		return nil, err
	}
	user := models.User{}
	db.Where("email = ?", email).First(&user)

	if user.ID == 0 {
		return nil, errors.New("user not found")

	}

	if password != user.Password {
		return nil, errors.New("invalid password")
	}

	return &user, nil
}
