package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"gorm.io/gorm"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"unique"`
	Password string
	Role     string
}
type Register struct {
	Name     string
	Email    string
	Password string
}
type Product struct {
	gorm.Model
	ID          uint
	Title       string
	Price       float64
	Description string
	Category    string
	Image       string
	Quantity    int
}
type Cart struct {
	gorm.Model
	UserID    uint
	ProductID uint
	Quantity  int
}
type BoughtProduct struct {
	gorm.Model
	UserID      uint
	ProductID   uint
	Quantity    int
	TotalAmount float64
}
type GooglePayload struct {
	SUB           string `json:"sub"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Locale        string `json:"locale"`
}

func ConvertToken(accessToken string) (*GooglePayload, error) {

	resp, httpErr := http.Get(fmt.Sprintf("https://www.googleapis.com/oauth2/v3/userinfo?access_token=%s", accessToken))
	if httpErr != nil {
		return nil, httpErr
	}

	defer resp.Body.Close()

	respBody, bodyErr := ioutil.ReadAll(resp.Body)
	if bodyErr != nil {
		return nil, bodyErr
	}

	var body map[string]interface{}
	if err := json.Unmarshal(respBody, &body); err != nil {
		return nil, err
	}

	if body["error"] != nil {
		return nil, errors.New("invalid token")
	}

	var data GooglePayload
	err := json.Unmarshal(respBody, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
