package controller

import (
	"blog/models"
	"blog/service"
	"blog/utils"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func GetAllCategory(c echo.Context) error {
	category, err := service.AllCategory()
	if err != nil {
		response := models.Response{
			Message: "Raono woy",
			Status:  false,
		}
		return c.JSON(http.StatusNotFound, response)
	}

	return c.JSON(http.StatusOK, category)
}

func GetSpecCategory(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var getCategory models.Kategori

	getCategory, err := service.Category(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, getCategory)
}

func CategoryAdd(c echo.Context) error {
	e := echo.New()
	e.Validator = &utils.CustomValidator{Validator: validator.New()}
	var createCategory models.Kategori

	c.Bind(&createCategory)

	err := c.Validate(&createCategory)

	if err != nil {
		return c.JSON(http.StatusBadRequest, "Tidak ditemukan")

	}
	err = service.CreateCategory(createCategory) // Memanggil fungsi CreateCategory dari service

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Gagal menambahkan kategori")
	}

	return c.JSON(http.StatusOK, "ok")

}
