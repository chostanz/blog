package controller

import (
	"blog/models"
	"blog/service"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt"
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

// func CreateContent(c echo.Context) error {
// 	tokenStr := c.Request().Header.Get("Authorization")
// 	if tokenStr == "" {
// 		return c.JSON(http.StatusUnauthorized, "Token not provided")
// 	}

// 	e := echo.New()
// 	e.Validator = &utils.CustomValidator{Validator: validator.New()}
// 	var createContent models.Content

// 	c.Bind(&createContent)

// 	err := c.Validate(&createContent)

// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, "Data yang dimasukkan tidak valid")

// 	}
// 	_, errService := service.GetAuthorID(tokenStr)

// 	if errService != nil {
// 		return c.JSON(http.StatusBadRequest, "Gagal menambahkan konten")
// 	}

// 	return c.JSON(http.StatusOK, "Berhasil menambahkan konten")
// }

func CreateContent(c echo.Context) error {
	tokenStr := c.Request().Header.Get("Authorization")
	if tokenStr == "" {
		return c.JSON(http.StatusUnauthorized, "Token not provided")
	}

	// Memvalidasi token
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "Invalid token")
	}

	// Mengambil payload dari token
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Menggunakan informasi dari payload, seperti user ID
		IdUser := int(claims["user_id"].(float64))
		contentData := models.Content{
			Author_id: IdUser,
		}

		errService := service.CreateContent(contentData, tokenStr)

		if errService != nil {
			return c.JSON(http.StatusBadRequest, "Gagal menambahkan konten")
		}
		return c.JSON(http.StatusOK, "Content created successfully")
	}

	return c.JSON(http.StatusUnauthorized, "Invalid token")
}
