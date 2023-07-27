package models

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type User struct {
	ID       int
	Email    string
	Password string
	Name     string
}
type Register struct {
	Name     string
	Email    string
	Password string
}
type AddToCartRequest struct {
	ProductID int `json:"productId"`
	UserID    int `json:"userId"`
}
type Customer struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
type AdminLoginRequest struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}
