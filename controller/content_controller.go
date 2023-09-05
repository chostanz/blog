package controller

import (
	"blog/models"
	"blog/service"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt"
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

func GetMyContent(c echo.Context) error {
	tokenStr := c.Request().Header.Get("Authorization")
	if tokenStr == "" {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{"error": "Token not provided"})
	}

	tokenSplit := strings.Split(tokenStr, " ")
	if len(tokenSplit) != 2 {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{"error": "Invalid token format"})
	}

	tokenOnly := tokenSplit[1]

	// Gunakan pustaka JWT untuk memeriksa dan memecahkan token
	token, err := jwt.Parse(tokenOnly, func(token *jwt.Token) (interface{}, error) {
		return []byte("rahasia"), nil // Ganti dengan kunci rahasia Anda
	})
	if err != nil || !token.Valid {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{"error": "Invalid token"})
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{"error": "Invalid token claims"})
	}

	authorID := int(claims["id_user"].(float64)) // Pastikan "user_id" sesuai dengan yang disimpan dalam token

	myContent, err := service.MyContent(authorID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": "Failed to fetch content"})
	}

	return c.JSON(http.StatusOK, myContent)
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

	var deleteContent models.Content
	if err := c.Bind(&deleteContent); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid data provided"})
	}
	// Mengambil author_id dari konten yang ingin diedit
	originalContent, err := service.Content(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": "Failed to fetch content"})
	}

	// Memeriksa apakah user_id dari token cocok dengan author_id dari konten
	if userID != originalContent.Author_id {
		return c.JSON(http.StatusForbidden, map[string]interface{}{"error": "You are not authorized to delete this content"})
	}

	_, err = service.EditContent(deleteContent, id, userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": "Failed to delete content"})
	}

	//	return echo.NewHTTPError(http.StatusInternalServerError, err.Error())

	return c.JSON(http.StatusOK, "okee")
}

type ContentController struct {
	contentService *service.ContentService
}

func NewContentController(contentService *service.ContentService) *ContentController {
	return &ContentController{contentService: contentService}
}

func (c *ContentController) UploadCoverImage(e echo.Context) error {
	contentID, err := strconv.Atoi(e.Param("contentID"))
	if err != nil {
		return e.String(http.StatusBadRequest, "Invalid content ID")
	}

	// Menerima berkas yang diunggah
	file, err := e.FormFile("cover_image")
	if err != nil {
		return e.String(http.StatusBadRequest, "Error uploading image")
	}

	// Simpan gambar dengan nama unik di server (ganti dengan path yang sesuai)
	pathGambar := "E:/golang/blog/cover/" + file.Filename
	if err := saveUploadedFile(file, pathGambar); err != nil {
		return e.String(http.StatusInternalServerError, "Error saving image")
	}

	// Dapatkan URL gambar yang baru diunggah
	baseURL := "http://localhost:8080"
	coverURL := baseURL + "/cover/" + file.Filename

	// Upload URL gambar sampul ke layanan ContentService
	if err := c.contentService.UploadCoverImage(int(contentID), coverURL); err != nil {
		return e.String(http.StatusInternalServerError, "Failed to upload cover image URL")
	}

	return e.String(http.StatusOK, "Cover image uploaded and URL updated")
}
