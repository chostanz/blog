package controller

import (
	"blog/models"
	"blog/service"
	"database/sql"
	"log"
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
		// Handle error 500 (Internal Server Error)
		response := models.Response{
			Code:    500,
			Message: "Terjadi kesalahan internal pada server. Mohon coba beberapa saat lagi!",
			Status:  false,
		}
		return c.JSON(http.StatusInternalServerError, response)
	}
	return c.JSON(http.StatusOK, content)
}

func GetSpecContent(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var getContent models.Content

	getContent, err := service.Content(id)
	if err != nil {
		if err == sql.ErrNoRows {
			response := models.Response{
				Code:    404,
				Message: "Konten tidak ditemukan!",
				Status:  false,
			}
			return c.JSON(http.StatusNotFound, response)
		} else {
			return c.JSON(http.StatusInternalServerError, &models.Response{
				Code:    500,
				Message: "Terjadi kesalahan internal pada server. Mohon coba beberapa saat lagi!",
				Status:  false,
			})
		}
	}

	return c.JSON(http.StatusOK, getContent)
}

func GetMyContent(c echo.Context) error {
	tokenStr := c.Request().Header.Get("Authorization")
	if tokenStr == "" {
		return c.JSON(http.StatusUnauthorized, "Token tidak tersedia")
	}

	tokenSplit := strings.Split(tokenStr, " ")
	if len(tokenSplit) != 2 {
		return c.JSON(http.StatusUnauthorized, "Invalid token!")
	}

	tokenOnly := tokenSplit[1]

	// Gunakan pustaka JWT untuk memeriksa dan memecahkan token
	token, err := jwt.Parse(tokenOnly, func(token *jwt.Token) (interface{}, error) {
		return []byte("rahasia"), nil // Ganti dengan kunci rahasia Anda
	})
	if err != nil || !token.Valid {
		return c.JSON(http.StatusUnauthorized, &models.Response{
			Code:    401,
			Message: "Invalid token",
			Status:  false,
		})
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.JSON(http.StatusUnauthorized, &models.Response{
			Code:    401,
			Message: "Invalid token claims!",
			Status:  false,
		})
	}

	authorID := int(claims["id_user"].(float64)) // Pastikan "user_id" sesuai dengan yang disimpan dalam token

	myContent, err := service.MyContent(authorID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &models.Response{
			Code:    500,
			Message: "Terjadi kesalahan internal pada server. Mohon coba beberapa saat lagi!",
			Status:  false,
		})
	}

	return c.JSON(http.StatusOK, myContent)
}

func CreateContent(c echo.Context) error {
	tokenStr := c.Request().Header.Get("Authorization")
	tokenSplit := strings.Split(tokenStr, " ")
	tokenOnly := tokenSplit[1]

	if tokenStr == "" {
		return c.JSON(http.StatusUnauthorized, "Token tidak tersedia")
	}
	authorID := c.Get("author_id").(int)
	_, err := service.GetAuthorInfoFromToken(tokenOnly)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "Invalid token atau token tidak ditemukan!")
	}

	var createContent models.Content
	if err := c.Bind(&createContent); err != nil {
		return c.JSON(http.StatusBadRequest, &models.Response{
			Code:    400,
			Message: "Data invalid!",
			Status:  false,
		})
	}

	errValidate := c.Validate(&createContent)
	if errValidate != nil {
		return c.JSON(http.StatusInternalServerError, &models.Response{
			Code:    422,
			Message: "Data tidak boleh kosong!",
			Status:  false,
		})
	}

	err = service.CreateContent(createContent, authorID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &models.Response{
			Code:    500,
			Message: "Terjadi kesalahan internal pada server. Mohon coba beberapa saat lagi!",
			Status:  false,
		})
	}

	return c.JSON(http.StatusCreated, &models.Response{
		Code:    201,
		Message: "Konten berhasil dibuat!",
		Status:  true,
	})

}

func ContentUpdate(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	tokenStr := c.Request().Header.Get("Authorization")
	tokenSplit := strings.Split(tokenStr, " ")
	tokenOnly := tokenSplit[1]

	if tokenStr == "" {
		return c.JSON(http.StatusUnauthorized, "Token tidak tersedia")
	}

	userID, err := service.GetAuthorInfoFromToken(tokenOnly)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "Invalid token atau token tidak ada")
	}

	perviousContent, errGet := service.Content(id)
	if errGet != nil {
		return c.JSON(http.StatusInternalServerError, &models.Response{
			Code:    500,
			Message: "Gagal mengambil data content saat ini. Mohon coba beberapa saat lagi!",
			Status:  false,
		})
	}

	var editContent models.Content
	if err := c.Bind(&editContent); err != nil {
		return c.JSON(http.StatusBadRequest, &models.Response{
			Code:    400,
			Message: "Data invalid!",
			Status:  false,
		})
	}

	errValidate := c.Validate(&editContent)
	if errValidate != nil {
		return c.JSON(http.StatusInternalServerError, &models.Response{
			Code:    422,
			Message: "Data tidak boleh kosong!",
			Status:  false,
		})
	}
	// Mengambil author_id dari konten yang ingin diedit
	originalContent, err := service.Content(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &models.Response{
			Code:    500,
			Message: "Terjadi kesalahan internal server. Konten dengan id tersebut tidak ada",
			Status:  false,
		})
	}

	// Memeriksa apakah user_id dari token cocok dengan author_id dari konten
	if userID != originalContent.Author_id {
		return c.JSON(http.StatusForbidden, &models.Response{
			Code:    403,
			Message: "Kamu tidak memiliki akses untuk mengedit content ini",
			Status:  false,
		})
	}

	c.Set("perviousContent", perviousContent)

	_, err = service.EditContent(editContent, id, userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &models.Response{
			Code:    500,
			Message: "Terjadi kesalahan internal pada server. Mohon coba beberapa saat lagi!",
			Status:  false,
		})
	}

	log.Println(perviousContent)
	return c.JSON(http.StatusOK, &models.Response{
		Code:    200,
		Message: "Konten telah diperbarui",
		Status:  true,
	})
}

func ContentDelete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	tokenStr := c.Request().Header.Get("Authorization")
	tokenSplit := strings.Split(tokenStr, " ")
	tokenOnly := tokenSplit[1]

	if tokenStr == "" {
		return c.JSON(http.StatusUnauthorized, "Token tidak tersedia")
	}

	userID, err := service.GetAuthorInfoFromToken(tokenOnly)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, &models.Response{
			Code:    401,
			Message: "Invalid token atau token tidak ada",
			Status:  false,
		})
	}

	var deleteContent models.Content
	if err := c.Bind(&deleteContent); err != nil {
		return c.JSON(http.StatusBadRequest, &models.Response{
			Code:    400,
			Message: "Data invalid",
			Status:  false,
		})
	}
	// Mengambil author_id dari konten yang ingin diedit
	originalContent, err := service.Content(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, &models.Response{
			Code:    404,
			Message: "Konten tidak tersedia!",
			Status:  false,
		})
	}

	// Memeriksa apakah user_id dari token cocok dengan author_id dari konten
	if userID != originalContent.Author_id {
		return c.JSON(http.StatusForbidden, &models.Response{
			Code:    403,
			Message: "Kamu tidak memiliki akses untuk menghapus content ini",
			Status:  false,
		})
	}

	_, err = service.DeleteContent(deleteContent, id, userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &models.Response{
			Code:    500,
			Message: "Terjadi kesalahan internal server. Gagal menghapus content",
			Status:  false,
		})
	}
	return c.JSON(http.StatusOK, &models.Response{
		Code:    200,
		Message: "Berhasil dihapus!",
		Status:  true,
	})
}

func DeleteContent(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	tokenStr := c.Request().Header.Get("Authorization")
	tokenSplit := strings.Split(tokenStr, " ")
	tokenOnly := tokenSplit[1]

	if tokenStr == "" {
		return c.JSON(http.StatusUnauthorized, &models.Response{
			Code:    401,
			Message: "Token tidak tersedia",
			Status:  false,
		})
	}

	userID, err := service.GetAuthorInfoFromToken(tokenOnly)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, &models.Response{
			Code:    401,
			Message: "Invalid token atau token tidak ada!",
			Status:  false,
		})
	}

	var deleteContent models.Content
	if err := c.Bind(&deleteContent); err != nil {
		return c.JSON(http.StatusBadRequest, &models.Response{
			Code:    400,
			Message: "Data invalid!",
			Status:  false,
		})
	}

	deleteContent, err = service.Content(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, &models.Response{
			Code:    404,
			Message: "Konten tidak tersedia!",
			Status:  false,
		})
	}

	_, err = service.DeleteContent(deleteContent, id, userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &models.Response{
			Code:    500,
			Message: "Terjadi kesalahan internal server. Gagal menghapus content",
			Status:  false,
		})
	}
	return c.JSON(http.StatusOK, &models.Response{
		Code:    200,
		Message: "Berhasil dihapus!",
		Status:  true,
	})
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
		return e.JSON(http.StatusBadRequest, &models.Response{
			Code:    400,
			Message: "Invalid content ID",
			Status:  false,
		})
	}

	// Menerima berkas yang diunggah
	file, err := e.FormFile("cover_image")
	if err != nil {
		return e.JSON(http.StatusBadRequest, &models.Response{
			Code:    400,
			Message: "Error mengupload gambar",
			Status:  false,
		})
	}

	// Simpan gambar dengan nama unik di server (ganti dengan path yang sesuai)
	pathGambar := "E:/golang/blog/cover/" + file.Filename
	if err := saveUploadedFile(file, pathGambar); err != nil {
		return e.JSON(http.StatusInternalServerError, &models.Response{
			Code:    500,
			Message: "Terjadi kesalahan internal server. Mohon coba beberapa saat lagi!",
			Status:  false,
		})
	}

	// Dapatkan URL gambar yang baru diunggah
	baseURL := "http://localhost:1234"
	coverURL := baseURL + "/cover/" + file.Filename

	// Upload URL gambar sampul ke layanan ContentService
	if err := c.contentService.UploadCoverImage(int(contentID), coverURL); err != nil {
		return e.JSON(http.StatusInternalServerError, "Failed to upload cover image URL")
	}

	return e.JSON(http.StatusOK, &models.Response{
		Code:    200,
		Message: "Berhasil mengupload gambar!",
		Status:  true,
	})
}
