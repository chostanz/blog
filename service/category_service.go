package service

import (
	"blog/models"
	"strconv"
)

func AllCategory() ([]models.Kategori, error) {
	categoryGet := []models.Kategori{}

	rows, err := db.Queryx("select * from categories")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		place := models.Kategori{}
		rows.StructScan(&place)
		categoryGet = append(categoryGet, place)
	}
	return categoryGet, nil

}

func Category(id int) (models.Kategori, error) {
	var specCategory models.Kategori
	idStr := strconv.Itoa(id)

	err := db.Get(&specCategory, "SELECT * FROM categories WHERE id = $1", idStr)
	if err != nil {
		return models.Kategori{}, err
	}
	return specCategory, nil
}

func CreateCategory(createCategory models.Kategori) error {
	_, err := db.NamedExec("INSERT INTO categories (category) VALUES (:category)", createCategory)
	if err != nil {
		return err
	}
	return nil
}
