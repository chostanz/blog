package models

type LoginParam struct {
	Username string `json:"username" db:"username" validate:"required"`
	Password string `json:"password" db:"password" validate:"required"`
	Role     int    `json:"id_role" db:"role_id"`
}

type RegisterParam struct {
	Id       int    `json:"id" db:"id"`
	Username string `json:"username" db:"username" validate:"required"`
	Email    string `json:"email" db:"email" validate:"required,email"`
	Password string `json:"password" db:"password" validate:"required"`
}
