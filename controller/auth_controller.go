package controller

import (
	"blog/models"
	"blog/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	loginbody := new(models.LoginParam)
	c.Bind(loginbody)

	err := c.Validate(loginbody)

	if err == nil {
		return c.JSON(http.StatusBadRequest, &models.LoginResp{
			Message: "Data login invalid",
			Status:  false,
		})
	}
	_, isAuthentication := service.CheckCredential(*loginbody)

	if !isAuthentication {
		return c.JSON(http.StatusUnauthorized, &models.LoginResp{
			Message: "Akun tidak ada atau password salah",
			Status:  false,
		})
	}

	return c.JSON(http.StatusOK, &models.LoginResp{
		Message: "Berhasil login",
		Status:  true,
	})

}
