package models

type Kategori struct {
	Id       int    `json:"id" db:"id"`
	Category string `json:"category" db:"category" validate:"required"`
}
