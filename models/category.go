package models

type Kategori struct {
	Id       int    `json:"id" db:"id"`
	Category string `json:"category" db:"category" validate:"required"`
}

type SpecCategory struct {
	Id         int    `json:"id" db:"id"`
	Category   string `json:"category" db:"category" validate:"required"`
	Title      string `json:"title" db:"title"`
	Content    string `json:"content" db:"content"`
	CreatedAt  string `json:"created_at" db:"created_at"`
	ModifiedAt string `json:"modified_at" db:"modified_at"`
	CreatedBy  string `json:"created_by" db:"created_by"`
	ModifiedBy string `json:"modified_by" db:"modified_by"`
}
