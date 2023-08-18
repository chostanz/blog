package service

import (
	"blog/models"
	"strconv"
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
