package models

type User struct {
	ID       int    `json:"id,omitempty"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
