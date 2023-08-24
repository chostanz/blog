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

func Category(id int) (models.SpecCategory, error) {
	var specCategory models.SpecCategory
	idStr := strconv.Itoa(id)

	err := db.Get(&specCategory, "SELECT c.category, co.id, co.title, co.content, co.created_at, co.modified_at, co.created_by, co.modified_by FROM categories c JOIN contents co ON c.id = co.category_id WHERE c.id = $1", idStr)
	if err != nil {
		return models.SpecCategory{}, err
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

func EditCategory(editCategory models.Kategori, id int) (models.Kategori, error) {
	idStr := strconv.Itoa(id)

	_, err := db.NamedExec("UPDATE categories SET category = :category WHERE id = :id", map[string]interface{}{
		"category": editCategory.Category,
		"id":       idStr,
	})

	if err != nil {
		return models.Kategori{}, err
	}
	return editCategory, nil
}

func DeleteCategory(deleteCategory models.Kategori, id int) (models.Kategori, error) {
	idStr := strconv.Itoa(id)

	_, err := db.Exec("Delete from categories where id = $1", idStr)
	if err != nil {
		return models.Kategori{}, err
	}
	return deleteCategory, nil
}
