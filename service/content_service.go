package service

import (
	"blog/models"
	"strconv"

	"github.com/golang-jwt/jwt"
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
	if err != nil {
		return models.Content{}, err
	}
	return specContent, nil

}

func GetAuthorID(tokenStr string) (int, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte("rahasia"), nil
	})

	if err != nil {
		return 0, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if customClaims, ok := claims["custom_claims"].(map[string]interface{}); ok {
			if authorID, ok := customClaims["author_id"]; ok {
				if id, ok := authorID.(float64); ok {
					return int(id), nil
				}
			}
		}
	}

	// if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
	// 	authorID := int(claims["author_id"].(float64))
	// 	return authorID, nil
	// }

	return 0, err
}

func GetUsernameByID(userID int) (string, error) {
	var username string
	err := db.QueryRow("SELECT username FROM users WHERE id = $1", userID).Scan(&username)
	if err != nil {
		return "", err
	}
	return username, nil
}

func CreateContent(createContent models.Content, tokenStr string) error {
	authorID, err := GetAuthorID(tokenStr)
	if err != nil {
		return err
	}

	// Dapatkan username berdasarkan authorID
	username, err := GetUsernameByID(authorID)
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}
	_, errInsert := db.Exec("INSERT INTO contents (author_id, title, content, category_id, created_by) VALUES ($1, $2, $3, $4, $5)", authorID, createContent.Title, createContent.Content_post, createContent.Category_id, username)

	if errInsert != nil {
		return errInsert
	}
	return nil
}
func EditContent(editContent models.Content, id int) (models.Content, error) {
	idStr := strconv.Itoa(id)

	_, err := db.NamedExec("UPDATE contents SET title = :title, content = :content, category_id = :category_id, author_id = :author_id WHERE id = :id", map[string]interface{}{
		"title":       editContent.Title,
		"content":     editContent.Content_post,
		"category_id": editContent.Category_id,
		"author_id":   editContent.Author_id,
		"id":          idStr,
	})

	if err != nil {
		return models.Content{}, err
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
