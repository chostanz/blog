package models

type LoginParam struct {
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}

type RegisterParam struct {
	Id       int    `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}
