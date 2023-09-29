package models

import (
	"database/sql"
	"time"
)

type Content struct {
	Id            int            `json:"id" db:"id"`
	Author_id     int            `json:"author_id" db:"author_id"`
	Category_id   int            `json:"category_id" db:"category_id"`
	Title         string         `json:"title" db:"title" validate:"required"`
	Content_post  string         `json:"content" db:"content" validate:"required"`
	CoverImageURL string         `json:"cover_image_url" db:"cover_image_url"`
	Created_at    time.Time      `json:"created_at" db:"created_at"`
	Modified_at   sql.NullTime   `json:"modified_at" db:"modified_at"`
	Created_by    string         `json:"created_by" db:"created_by"`
	Modified_by   sql.NullString `json:"modified_by" db:"modified_by"`
}
