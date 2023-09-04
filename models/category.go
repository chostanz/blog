package models

import (
	"database/sql"
	"time"
)

type Kategori struct {
	Id       int    `json:"id" db:"id"`
	Category string `json:"category" db:"name" validate:"required"`
}

type SpecCategory struct {
	Id         int            `json:"id" db:"id"`
	Category   string         `json:"category" db:"name" validate:"required"`
	Title      string         `json:"title" db:"title"`
	Content    string         `json:"content" db:"content"`
	CreatedAt  time.Time      `json:"created_at" db:"created_at"`
	ModifiedAt sql.NullTime   `json:"modified_at" db:"modified_at"`
	CreatedBy  string         `json:"created_by" db:"created_by"`
	ModifiedBy sql.NullString `json:"modified_by" db:"modified_by"`
}
