package controller

import (
	"blog/models"
	"blog/service"
	"blog/utils"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	e := echo.New()
	e.Validator = &utils.CustomValidator{Validator: validator.New()}

	var loginbody models.LoginParam
	c.Bind(&loginbody)

	err := c.Validate(&loginbody)

	if err != nil {
		return c.JSON(http.StatusBadRequest, &models.LoginResp{
			Message: "Data login invalid",
			Status:  false,
		})
	}
	_, isAuthentication := service.CheckCredential(loginbody)

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

func Register(c echo.Context) error {
	e := echo.New()
	e.Validator = &utils.CustomValidator{Validator: validator.New()}

	var userRegister models.RegisterParam

	c.Bind(&userRegister)
	err := c.Validate(&userRegister)

	if err == nil {
		registerErr := service.RegisterUser(userRegister)
		if registerErr != nil {

			return echo.NewHTTPError(http.StatusBadRequest, "Username telah digunakan!")
		}
		return c.JSON(http.StatusCreated, &models.RegisterResp{
			Message: "Berhasil register",
			Status:  true,
		})
	}

	return echo.NewHTTPError(http.StatusBadRequest, err.Error())

}

// func Register(c echo.Context) error {
// 	e := echo.New()
// 	e.Validator = &utils.CustomValidator{Validator: validator.New()}

// 	var userRegister models.RegisterParam

// 	if err := c.Bind(&userRegister); err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
// 	}

// 	if err := c.Validate(&userRegister); err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
// 	}

// 	service.RegisterUser(userRegister)

// 	return c.JSON(http.StatusCreated, &models.RegisterResp{
// 		Message: "Berhasil register",
// 		Status:  true,
// 	})
// }
