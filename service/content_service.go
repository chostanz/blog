package service

import (
	"blog/models"
	"strconv"

	"github.com/golang-jwt/jwt"
)

func ContentAll() ([]models.Content, error) {
	contentGet := []models.Content{}

	rows, err := db.Queryx("select * from content")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		//membuat var bru dari struct Karyawan yg nilai awalnya kosong
		place := models.Content{}
		rows.StructScan(&place)
		contentGet = append(contentGet, place) //menambah elemen baru menggunakan append ke users
	}
	return contentGet, nil
}

func Content(id int) (models.Content, error) {
	var specContent models.Content
	idStr := strconv.Itoa(id)

	err := db.Get(&specContent, "SELECT * FROM content WHERE id = $1", idStr)
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
		authorID := int(claims["author_id"].(float64))
		return authorID, nil
	}

	return 0, err
}

func CreateContent(createContent models.Content, tokenStr string) error {
	authorID, err := GetAuthorID(tokenStr)
	if err != nil {
		return err
	}
	_, errInsert := db.Exec("INSERT into content (author_id, title, content) VALUES ($1, $2, $3)", authorID, createContent.Title, createContent.Content_post)
	if errInsert != nil {
		return err
	}
	return nil
}
