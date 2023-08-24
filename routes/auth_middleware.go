package routes

import (
	"blog/models"
	"blog/service"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func AuthorMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			return c.JSON(http.StatusUnauthorized, "Missing token")
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Ganti "rahasia" dengan kunci yang sesuai
			return []byte("rahasia"), nil
		})
		if err != nil || !token.Valid {
			fmt.Println("Error verifying token:", err)
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

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			return c.JSON(http.StatusUnauthorized, "Missing token")
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte("rahasia"), nil // Ganti dengan kunci Anda
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

func AdminMiddleware(admin echo.HandlerFunc) echo.HandlerFunc {
	return middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		_, isAuthentication, _, err := service.CheckCredential(models.LoginParam{Username: username, Password: password, Role: 1})
		if err != nil {
			return false, err
		}
		return isAuthentication, nil
	})(admin)
}
