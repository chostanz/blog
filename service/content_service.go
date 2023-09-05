package service

import (
	"blog/models"
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/jmoiron/sqlx"
)

func ContentAll() ([]models.Content, error) {
	contentGet := []models.Content{}

	rows, err := db.Queryx("select * from contents")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		place := models.Content{}
		rows.StructScan(&place)
		contentGet = append(contentGet, place) //menambah elemen baru menggunakan append ke users
	}
	return contentGet, nil
}

func Content(id int) (models.Content, error) {
	var specContent models.Content
	idStr := strconv.Itoa(id)

	err := db.Get(&specContent, "SELECT * FROM contents WHERE id = $1", idStr)
	// if err != nil {
	// 	return models.Content{}, err
	// }
	return specContent, err

}

func MyContent(authorID int) ([]models.Content, error) {
	idStr := strconv.Itoa(authorID)
	var myContent []models.Content

	err := db.Select(&myContent, "SELECT * FROM contents WHERE author_id = $1", idStr)
	if err != nil {
		fmt.Printf("Error fetching content: %s", err)
	}
	return myContent, err
}

func GetUsernameByID(authorID int) (string, error) {
	var username string
	err := db.QueryRow("SELECT username FROM users WHERE id = $1", authorID).Scan(&username)
	if err != nil {
		return "", err
	}
	return username, nil
}

func GetAuthorInfoFromToken(tokenStr string) (int, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte("rahasia"), nil
	})
	if err != nil || !token.Valid {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("Invalid token claims")
	}

	authorID := int(claims["id_user"].(float64))

	return authorID, nil
}

func CreateContent(createContent models.Content, authorID int) error {
	username, err := GetUsernameByID(authorID)
	if err != nil {
		return err
	}
	_, errInsert := db.Exec("INSERT INTO contents (author_id, title, content, category_id, created_by) VALUES ($1, $2, $3, $4, $5)", authorID, createContent.Title, createContent.Content_post, createContent.Category_id, username)

	if errInsert != nil {
		return errInsert
	}
	return nil
}

func EditContent(editContent models.Content, id int, authorID int) (models.Content, error) {
	idStr := strconv.Itoa(id)
	username, err := GetUsernameByID(authorID)
	if err != nil {
		return models.Content{}, err
	}

	_, errInsert := db.NamedExec("UPDATE contents SET title = :title, content = :content, category_id = :category_id, author_id = :author_id, modified_by = :modified_by WHERE id = :id", map[string]interface{}{
		"title":       editContent.Title,
		"content":     editContent.Content_post,
		"category_id": editContent.Category_id,
		"author_id":   authorID,
		"modified_by": username,
		"id":          idStr,
	})

	if errInsert != nil {
		return models.Content{}, errInsert
	}
	return editContent, nil
}

func DeleteContent(deleteContent models.Content, id int) (models.Content, error) {
	idStr := strconv.Itoa(id)

	_, err := db.Exec("Delete from contents where id = $1", idStr)
	if err != nil {
		return models.Content{}, err
	}
	return deleteContent, nil
}

type ContentService struct {
	db *sqlx.DB
}

func NewContentService(db *sqlx.DB) *ContentService {
	return &ContentService{db: db}
}

func (s *ContentService) UploadCoverImage(contentID int, coverURL string) error {
	_, err := s.db.Exec("UPDATE contents SET cover_image_url = $1 WHERE id = $2", coverURL, contentID)
	if err != nil {
		log.Println("Error updating cover image URL:", err)
		return err
	}
	return nil
}
