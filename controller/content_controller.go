package controller

import (
	"blog/models"
	"blog/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetAllContent(c echo.Context) error {
	content, err := service.ContentAll()

	if err != nil {
		response := models.Response{
			Message: "Halaman tidak ada atau url salah",
			Status:  false,
		}
		return c.JSON(http.StatusNotFound, response)
	}

	return c.JSON(http.StatusOK, content)
}

func GetSpecContent(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var getContent models.Content

	getContent, err := service.Content(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, getContent)

}
