package models

type Users struct {
	Id       int    `json:"id" db:"id"`
	Username string `json:"username" db:"username" validate:"required"`
	Email    string `json:"email" db:"email" validate:"required, email"`
	Password string `json:"password" db:"password" validate:"required"`
}

type Password struct {
	Password string `json:"password" db:"password" validate:"required"`
}
