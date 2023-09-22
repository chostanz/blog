package middleware

import (
	"blog/controller"
	"blog/models"
	"blog/service"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		tokenSplit := strings.Split(tokenString, " ")
		tokenOnly := tokenSplit[1]
		fmt.Println("Token", tokenOnly)
		if tokenString == "" {
			return c.JSON(http.StatusUnauthorized, &models.Response{
				Code:    401,
				Message: "Token tidak ditemukan!",
				Status:  false,
			})
		}

		token, err := jwt.Parse(tokenOnly, func(token *jwt.Token) (interface{}, error) {
			return []byte("rahasia"), nil
		})
		if err != nil || !token.Valid {
			return c.JSON(http.StatusUnauthorized, &models.Response{
				Code:    401,
				Message: "Token invalid!",
				Status:  false,
			})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.JSON(http.StatusUnauthorized, &models.Response{
				Code:    401,
				Message: "Token claims invalid!",
				Status:  false,
			})
		}
		// Check apakah token ada di invalidTokens
		if _, exists := controller.InvalidTokens[token.Raw]; exists {
			return c.JSON(http.StatusUnauthorized, &models.Response{
				Code:    401,
				Message: "Sesi berakhir! Silahkan login kembali",
				Status:  false,
			})
		}

		roleID := int(claims["id_role"].(float64))
		fmt.Println("Role ID:", roleID) // Tambahkan log ini untuk memeriksa nilai roleID

		if roleID == 1 || roleID == 2 {
			fmt.Println("Access granted")
			c.Set("id_user", int(claims["id_user"].(float64))) // Menyimpan ID User
			return next(c)
		} else {
			fmt.Println("Access denied")
			return c.JSON(http.StatusForbidden, &models.Response{
				Code:    403,
				Message: "Akses ditolak!",
				Status:  false,
			})
		}

	}
}

func AdminMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		tokenSplit := strings.Split(tokenString, " ")
		fmt.Println("Token:", tokenSplit)
		tokenOnly := tokenSplit[1]
		if tokenString == "" {
			return c.JSON(http.StatusUnauthorized, &models.Response{
				Code:    401,
				Message: "Token tidak ditemukan!",
				Status:  false,
			})
		}

		token, err := jwt.Parse(tokenOnly, func(token *jwt.Token) (interface{}, error) {
			return []byte("rahasia"), nil
		})

		if err != nil || !token.Valid {
			return c.JSON(http.StatusUnauthorized, &models.Response{
				Code:    401,
				Message: "Token invalid!",
				Status:  false,
			})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			fmt.Println("Tipe asli dari token.Claims:", reflect.TypeOf(token.Claims))
			return c.JSON(http.StatusUnauthorized, &models.Response{
				Code:    401,
				Message: "Token claims tidak valid!",
				Status:  false,
			})
		}

		fmt.Println(claims)

		// Check apakah token ada di invalidTokens
		if _, exists := controller.InvalidTokens[token.Raw]; exists {
			return c.JSON(http.StatusUnauthorized, &models.Response{
				Code:    401,
				Message: "Sesi berakhir! Silahkan login kembali",
				Status:  false,
			})
		}

		c.Set("users", token)

		roleID := int(claims["id_role"].(float64))
		if roleID != 1 {
			return c.JSON(http.StatusForbidden, &models.Response{
				Code:    403,
				Message: "Akses ditolak!",
				Status:  false,
			})
		}

		return next(c)
	}
}

func AuthorMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		tokenSplit := strings.Split(tokenString, " ")
		//fmt.Println("Token", tokenSplit)
		tokenOnly := tokenSplit[1]
		if tokenString == "" {
			return c.JSON(http.StatusUnauthorized, &models.Response{
				Code:    401,
				Message: "Token tidak ditemukan!",
				Status:  false,
			})
		}

		token, err := jwt.Parse(tokenOnly, func(token *jwt.Token) (interface{}, error) {
			return []byte("rahasia"), nil
		})
		if err != nil || !token.Valid {
			return c.JSON(http.StatusUnauthorized, &models.Response{
				Code:    401,
				Message: "Token invalid!",
				Status:  false,
			})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.JSON(http.StatusUnauthorized, &models.Response{
				Code:    401,
				Message: "Token claims invalid!",
				Status:  false,
			})
		}
		if _, exists := controller.InvalidTokens[token.Raw]; exists {
			return c.JSON(http.StatusUnauthorized, &models.Response{
				Code:    401,
				Message: "Sesi berakhir! Silahkan login kembali",
				Status:  false,
			})
		}
		authorID := int(claims["id_user"].(float64))
		c.Set("author_id", authorID)

		c.Set("users", token)

		roleID := int(claims["id_role"].(float64))
		if roleID != 2 {
			return c.JSON(http.StatusForbidden, &models.Response{
				Code:    403,
				Message: "Akses ditolak!",
				Status:  false,
			})
		}
		return next(c)
	}
}

func LogoutMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		tokenSplit := strings.Split(tokenString, " ")
		//fmt.Println("Token", tokenSplit)
		tokenOnly := tokenSplit[1]
		if tokenString == "" {
			return c.JSON(http.StatusUnauthorized, &models.Response{
				Code:    401,
				Message: "Token tidak ditemukan!",
				Status:  false,
			})
		}

		token, err := jwt.Parse(tokenOnly, func(token *jwt.Token) (interface{}, error) {
			return []byte("rahasia"), nil
		})
		if err != nil || !token.Valid {
			return c.JSON(http.StatusUnauthorized, &models.Response{
				Code:    401,
				Message: "Token invalid!",
				Status:  false,
			})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.JSON(http.StatusUnauthorized, &models.Response{
				Code:    401,
				Message: "Token claims invalid!",
				Status:  false,
			})
		}
		if _, exists := controller.InvalidTokens[token.Raw]; exists {
			return c.JSON(http.StatusUnauthorized, &models.Response{
				Code:    401,
				Message: "Sesi berakhir! Silahkan login kembali",
				Status:  false,
			})
		}
		authorID := int(claims["id_user"].(float64))
		c.Set("author_id", authorID)

		c.Set("users", token)

		roleID := int(claims["id_role"].(float64))
		if roleID == 1 || roleID == 2 {
			fmt.Println("Access granted")
		} else {
			return c.JSON(http.StatusForbidden, &models.Response{
				Code:    403,
				Message: "Akses ditolak!",
				Status:  false,
			})
		}
		return next(c)
	}
}
func SetContentServiceMiddleware(contentService *service.ContentService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("contentService", contentService)
			return next(c)
		}
	}
}
