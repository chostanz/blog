package middleware

import (
	"blog/controller"
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
			fmt.Println("Error: Token not provided") // Mencetak error
			return c.JSON(http.StatusUnauthorized, "Missing token")
		}

		token, err := jwt.Parse(tokenOnly, func(token *jwt.Token) (interface{}, error) {
			return []byte("rahasia"), nil
		})
		if err != nil || !token.Valid {
			return c.JSON(http.StatusUnauthorized, "Invalid token")
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.JSON(http.StatusUnauthorized, "Invalid token claims")
		}

		roleID := int(claims["id_role"].(float64))
		fmt.Println("Role ID:", roleID) // Tambahkan log ini untuk memeriksa nilai roleID

		if roleID == 1 || roleID == 2 {
			fmt.Println("Access granted")
			c.Set("id_user", int(claims["id_user"].(float64))) // Menyimpan ID User
			return next(c)
		} else {
			fmt.Println("Access denied")
			return c.JSON(http.StatusForbidden, "Access denied")
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
			return c.JSON(http.StatusUnauthorized, "Missing token")
		}

		token, err := jwt.Parse(tokenOnly, func(token *jwt.Token) (interface{}, error) {
			return []byte("rahasia"), nil
		})

		if err != nil || !token.Valid {
			return c.JSON(http.StatusUnauthorized, "Invalid token")
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			fmt.Println("Tipe asli dari token.Claims:", reflect.TypeOf(token.Claims))
			return c.JSON(http.StatusUnauthorized, "Klaim token tidak valid")
		}

		fmt.Println(claims)

		// Check if the token is in the invalidTokens list
		if _, exists := controller.InvalidTokens[token.Raw]; exists {
			return c.JSON(http.StatusUnauthorized, "Sesi berakhir! Silahkan login kembali")
		}

		c.Set("users", token)

		roleID := int(claims["id_role"].(float64))
		if roleID != 1 {
			return c.JSON(http.StatusForbidden, "Access denied")
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
			return c.JSON(http.StatusUnauthorized, "Missing token")
		}

		token, err := jwt.Parse(tokenOnly, func(token *jwt.Token) (interface{}, error) {
			return []byte("rahasia"), nil
		})
		if err != nil || !token.Valid {
			return c.JSON(http.StatusUnauthorized, "Invalid token")
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.JSON(http.StatusUnauthorized, "Invalid token claims")
		}
		if _, exists := controller.InvalidTokens[token.Raw]; exists {
			return c.JSON(http.StatusUnauthorized, "Sesi berakhir! Silahkan login kembali")
		}

		c.Set("users", token)

		roleID := int(claims["id_role"].(float64))
		if roleID != 2 {
			return c.JSON(http.StatusForbidden, "Access denied")
		}
		return next(c)
	}
}
