package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		tokenSplit := strings.Split(tokenString, " ")
		fmt.Println("Token", tokenSplit)
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

		c.Set("id_user", int(claims["id_user"].(float64))) // Menyimpan ID User dalam konteks
		return next(c)
	}
}

func AdminMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		tokenSplit := strings.Split(tokenString, " ")
		fmt.Println("Token", tokenSplit)
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
		fmt.Println("Token", tokenSplit)
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

		roleID := int(claims["id_role"].(float64))
		if roleID != 2 {
			return c.JSON(http.StatusForbidden, "Access denied")
		}

		return next(c)
	}
}
