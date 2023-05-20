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
	Username string
	Password string
	Name     string
}
type Register struct {
	Name     string
	Email    string
	Password string
}
