package controller

import (
	"blog/models"
	"blog/service"
	"blog/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func GetAllCategory(c echo.Context) error {
	category, err := service.AllCategory()
	if err != nil {
		response := models.Response{
			Message: "Halaman tidak ditemukan atau url salah",
			Status:  false,
		}
		return c.JSON(http.StatusNotFound, response)
	}

	return c.JSON(http.StatusOK, category)
}

func GetSpecCategory(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Println("Fetching category with ID:", id)

	getCategory, err := service.SpecCategory(id)
	if err != nil {
		fmt.Println("Error:", err)
		return echo.NewHTTPError(http.StatusNotFound, &models.Response{
			Message: "Kategori tidak ditemukan",
		})
	}

	fmt.Println("Category fetched successfully - ID:", getCategory.Id)
	return c.JSON(http.StatusOK, getCategory)
}

func GetContentCategory(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Println("Fetching category with ID:", id)

	getCategory, err := service.ContentCategory(id)
	if err != nil {
		fmt.Println("Error:", err)
		return echo.NewHTTPError(http.StatusNotFound, &models.Response{
			Message: "Kategori tidak ditemukan",
		})
	}

	fmt.Println("Category fetched successfully - ID:", getCategory.Id)
	return c.JSON(http.StatusOK, getCategory)
}

func CategoryAdd(c echo.Context) error {
	e := echo.New()
	e.Validator = &utils.CustomValidator{Validator: validator.New()}
	var createCategory models.Kategori
	c.Bind(&createCategory)
	err := c.Validate(&createCategory)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Data yang dimasukkan tidak valid")

	}
	err = service.CreateCategory(createCategory) // Memanggil fungsi CreateCategory dari service

	if err != nil {
		return c.JSON(http.StatusBadRequest, "Gagal menambahkan kategori")
	}

	return c.JSON(http.StatusOK, "Berhasil menambahkan kategori")

}

func CategoryUpdate(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var editCategory models.Kategori

	c.Bind(&editCategory)
	err := c.Validate(&editCategory)

	if err == nil {
		_, updateErr := service.EditCategory(editCategory, id)
		if updateErr != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "data belum dimasukkan")
		}
	}
	return c.JSON(http.StatusOK, &models.Response{
		Message: "Berhasil update",
		Status:  true,
	})
}

func CategoryDelete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var deleteCategory models.Kategori
	_, err := service.DeleteCategory(deleteCategory, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "okee")
}
