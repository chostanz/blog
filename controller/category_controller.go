package controller

import (
	"blog/models"
	"blog/service"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetAllCategory(c echo.Context) error {
	category, err := service.AllCategory()
	if err != nil {
		response := models.Response{
			Code:    500,
			Message: "Terjadi kesalahan internal pada server. Mohon coba beberapa saat lagi!",
			Status:  false,
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	return c.JSON(http.StatusOK, category)
}

func GetSpecCategory(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Println("Fetching category with ID:", id)

	getCategory, err := service.SpecCategory(id)
	if err != nil {
		if err == sql.ErrNoRows {
			response := models.Response{
				Code:    404,
				Message: "Kategori tidak ditemukan!",
				Status:  false,
			}
			return c.JSON(http.StatusNotFound, response)
		} else {
			return c.JSON(http.StatusInternalServerError, &models.Response{
				Code:    500,
				Message: "Terjadi kesalahan internal server. Mohon coba beberapa saat lagi!",
				Status:  false,
			})
		}
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
		return c.JSON(http.StatusInternalServerError, &models.Response{
			Code:    500,
			Message: "Terjadi kesalahan internal server. Mohon coba beberapa saat lagi",
			Status:  false,
		})
	}

	if len(getCategory) == 0 {
		fmt.Println("No content found in category - ID:", idStr)
		return c.JSON(http.StatusNotFound, &models.Response{
			Code:    http.StatusNotFound,
			Message: "Tidak ada konten dalam kategori ini",
			Status:  false,
		})
	}

	return c.JSON(http.StatusOK, getCategory)
}

func CategoryAdd(c echo.Context) error {
	var createCategory models.Kategori
	// c.Bind(&createCategory)
	if err := c.Bind(&createCategory); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, &models.Response{
			Code:    422,
			Message: "Permintaan tidak valid. data yang dimasukkan kosong atau tidak sesuai",
			Status:  false,
		})
	}
	err := c.Validate(&createCategory)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, &models.Response{
			Code:    422,
			Message: "Data tidak boleh kosong!",
			Status:  false,
		})
	}
	if err := service.CreateCategory(createCategory); err != nil {
		return c.JSON(http.StatusInternalServerError, &models.Response{
			Code:    500,
			Message: "Terjadi kesalahan internal server. Gagal menambahkan kategori!",
			Status:  false,
		})
	}
	return c.JSON(http.StatusCreated, &models.Response{
		Code:    201,
		Message: "Berhasil menambahkan kategori!",
		Status:  true,
	})

}

func CategoryUpdate(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var editCategory models.Kategori

	c.Bind(&editCategory)
	err := c.Validate(&editCategory)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, &models.Response{
			Code:    422,
			Message: "Data yang dimasukkan tidak valid",
			Status:  false,
		})

	}

	if err == nil {
		_, updateErr := service.EditCategory(editCategory, id)
		if updateErr != nil {
			return c.JSON(http.StatusInternalServerError, &models.Response{
				Code:    500,
				Message: "Tejadi kesalahan internal pada server. Mohon coba beberapa saat lagi!",
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

	var deleteCateg models.Kategori
	c.Bind(&deleteCateg)

	_, errID := service.SpecCategory(id)
	if errID != nil {
		return c.JSON(http.StatusNotFound, &models.Response{
			Code:    404,
			Message: "Kategori tidak tersedia!",
			Status:  false,
		})
	}

	var deleteCategory models.Kategori
	_, err := service.DeleteCategory(deleteCategory, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &models.Response{
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
