package controller

import (
	"blog/models"
	"blog/service"
	"fmt"
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

//		return c.JSON(http.StatusOK, "Berhasil menambahkan konten")
//	}

func CreateContent(c echo.Context) error {
	tokenStr := c.Request().Header.Get("Authorization")
	//fmt.Println(tokenStr)
	tokenSplit := strings.Split(tokenStr, " ")
	fmt.Println("Token", tokenSplit)
	tokenOnly := tokenSplit[1]
	if tokenStr == "" {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{"error": "Token not provided"})
	}

	var createContent models.Content
	if err := c.Bind(&createContent); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid data provided"})
	}

	authorID, err := service.GetAuthorID(tokenOnly)
	fmt.Println(authorID)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{"error": "Invalid or missing token"})
	}

	username, err := service.GetUsernameByID(authorID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}

	createContent.Author_id = authorID
	createContent.Created_by = username

	if err := service.CreateContent(createContent, tokenStr); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": "Failed to create content"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"message": "Content created successfully"})
}

func ContentUpdate(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var editContent models.Content

	c.Bind(&editContent)
	err := c.Validate(&editContent)

	if err == nil {
		_, updateErr := service.EditContent(editContent, id)
		if updateErr != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "data belum dimasukkan")
		}
	}
	return c.JSON(http.StatusOK, &models.Response{
		Message: "Berhasil update",
		Status:  true,
	})
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

// func CreateContent(c echo.Context) error {
// 	tokenStr := c.Request().Header.Get("Authorization")
// 	if tokenStr == "" {
// 		return c.JSON(http.StatusUnauthorized, "Token not provided")
// 	}

// 	// Memvalidasi token
// 	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
// 		return []byte("rahasia"), nil
// 	})
// 	if err == nil {
// 		return c.JSON(http.StatusUnauthorized, "Invalid token")
// 	}

// 	// Mengambil payload dari token
// 	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
// 		// Menggunakan informasi dari payload, seperti user ID
// 		IdUser := int(claims["id_user"].(float64))
// 		contentData := models.Content{
// 			Author_id: IdUser,
// 		}

// 		errService := service.CreateContent(contentData, tokenStr)

// 		if errService != nil {
// 			return c.JSON(http.StatusBadRequest, "Gagal menambahkan konten")
// 		}
// 		return c.JSON(http.StatusOK, "Content created successfully")
// 	}

// 	return c.JSON(http.StatusUnauthorized, "Invalid token")
// }
