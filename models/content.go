package models

type Content struct {
	Id           int    `json:"id" db:"id"`
	Author_id    int    `json:"author_id" db:"author_id"`
	Category_id  int    `json:"category_id" db:"category_id"`
	Title        string `json:"title" db:"title"`
	Content_post string `json:"content" db:"content"`
	Created_at   string `json:"created_at" db:"created_at"`
	Modified_at  string `json:"modified_at" db:"modified_at"`
}
