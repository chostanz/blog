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
			Code:    404,
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
		return c.JSON(http.StatusNotFound, &models.Response{
			Code:    404,
			Message: "Kategori tidak ditemukan",
			Status:  false,
		})
	}

	fmt.Println("Category fetched successfully - ID:", getCategory.Id)
	return c.JSON(http.StatusOK, getCategory)
}

func GetContentCategory(c echo.Context) error {
	idStr, _ := strconv.Atoi(c.Param("id"))
	fmt.Println("Fetching category with ID:", idStr)

	getCategory, err := service.ContentCategory(idStr)
	if err != nil {
		fmt.Println("Error:", err)
		return c.JSON(http.StatusNotFound, &models.Response{
			Code:    404,
			Message: "Kategori tidak ditemukan",
			Status:  false,
		})
	}

	//fmt.Println("Category fetched successfully - ID:", getCategory[0].Id)
	return c.JSON(http.StatusOK, getCategory)
}

func CategoryAdd(c echo.Context) error {
	e := echo.New()
	e.Validator = &utils.CustomValidator{Validator: validator.New()}
	var createCategory models.Kategori
	c.Bind(&createCategory)
	err := c.Validate(&createCategory)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &models.Response{
			Code:    400,
			Message: "Data yang dimasukkan tidak valid",
			Status:  false,
		})

	}

	if err := service.CreateCategory(createCategory); err != nil {
		return c.JSON(http.StatusBadRequest, &models.Response{
			Code:    400,
			Message: "Gagal menambahkan kategori! Kategori sudah ada",
			Status:  false,
		})
	}
	return c.JSON(http.StatusOK, &models.Response{
		Code:    200,
		Message: "Berhasil menambahkan kategori!",
		Status:  true,
	})

}

func CategoryUpdate(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var editCategory models.Kategori

	c.Bind(&editCategory)
	err := c.Validate(&editCategory)

	if err == nil {
		_, updateErr := service.EditCategory(editCategory, id)
		if updateErr != nil {
			return c.JSON(http.StatusBadRequest, &models.Response{
				Code:    400,
				Message: "Kategori tidak boleh sama!",
				Status:  false,
			})
		}
	}
	return c.JSON(http.StatusOK, &models.Response{
		Code:    200,
		Message: "Kategori berhasil diubah!",
		Status:  true,
	})
}

func CategoryDelete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var deleteCategory models.Kategori
	_, err := service.DeleteCategory(deleteCategory, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, &models.Response{
			Code:    500,
			Message: "Terjadi kesalahan internal pada server. Mohon coba beberapa saat lagi!",
			Status:  false,
		})
	}
	return c.JSON(http.StatusOK, &models.Response{
		Code:    200,
		Message: "Berhasil dihapus!",
		Status:  true,
	})
}
