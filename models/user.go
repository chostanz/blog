package models

type Users struct {
	Id       int    `json:"id" db:"id"`
	Username string `json:"username" db:"username" validate:"required"`
	Email    string `json:"email" db:"email" validate:"required,email"`
	Role     string `json:"role" db:"role"`
}

type UserEdit struct {
	Id       int    `json:"id" db:"id"`
	Username string `json:"username" db:"username" validate:"required"`
	Email    string `json:"email" db:"email" validate:"required,email"`
	Password string `json:"password" db:"password" validate:"required"`
	Role     string `json:"role" db:"role"`
}

type Password struct {
	Password string `json:"password" db:"password" validate:"required"`
}

type User struct {
	PictureURL string `json:"picture_url" db:"picture_url"`
}

type Role struct {
	UserID int `json:"id" db:"user_id"`
	RoleID int `json:"role_id" db:"role_id" validate:"required"`
}

type Roles struct {
	Id   int    `json:"role_id" db:"id"`
	Role string `json:"role" db:"role"`
}
