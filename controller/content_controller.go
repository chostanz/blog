package controller

import (
	"blog/models"
	"blog/service"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type JwtCustomClaim struct {
	IdUser   int `json:"id_user"`
	IdAuthor int `json:"author_id"`
}

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

func CreateContent(c echo.Context) error {
	tokenStr := c.Request().Header.Get("Authorization")
	tokenSplit := strings.Split(tokenStr, " ")
	tokenOnly := tokenSplit[1]

	if tokenStr == "" {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{"error": "Token not provided"})
	}
	authorID := c.Get("author_id").(int)
	_, err := service.GetAuthorInfoFromToken(tokenOnly)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{"error": "Invalid or missing token"})
	}

	var createContent models.Content
	if err := c.Bind(&createContent); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid data provided"})
	}

	err = service.CreateContent(createContent, authorID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": "Failed to create content"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"message": "Content created successfully"})

}

func ContentUpdate(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	tokenStr := c.Request().Header.Get("Authorization")
	tokenSplit := strings.Split(tokenStr, " ")
	tokenOnly := tokenSplit[1]

	if tokenStr == "" {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{"error": "Token not provided"})
	}

	userID, err := service.GetAuthorInfoFromToken(tokenOnly)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{"error": "Invalid or missing token"})
	}

	var editContent models.Content
	if err := c.Bind(&editContent); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid data provided"})
	}
	// Mengambil author_id dari konten yang ingin diedit
	originalContent, err := service.Content(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": "Failed to fetch content"})
	}

	// Memeriksa apakah user_id dari token cocok dengan author_id dari konten
	if userID != originalContent.Author_id {
		return c.JSON(http.StatusForbidden, map[string]interface{}{"error": "You are not authorized to edit this content"})
	}

	_, err = service.EditContent(editContent, id, userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": "Failed to edit content"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"message": "Content edited successfully"})
}

func ContentDelete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var deleteContent models.Content
	_, err := service.DeleteContent(deleteContent, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "okee")
}
