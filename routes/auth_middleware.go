package routes

import (
	"blog/models"
	"blog/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func AuthorMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		_, isAuthentication, _, err := service.CheckCredential(models.LoginParam{Username: username, Password: password, Role: 2})
		if err != nil {
			return false, err
		}
		return isAuthentication, nil
	})(next)
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
