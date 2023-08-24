package controller

import (
	"blog/models"
	"blog/service"
	"blog/utils"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type TokenCheck struct {
	Token string `json:"token"`
}

type JwtCustomClaims struct {
	IdUser int `json:"id_user"`
	IdRole int `json:"id_role"`
	jwt.RegisteredClaims
}

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

	userId, isAuthentication, roleID, _ := service.CheckCredential(loginbody)

	if !isAuthentication {
		return c.JSON(http.StatusUnauthorized, &models.LoginResp{
			Message: "Akun tidak ada atau password salah",
			Status:  false,
		})
	}
	claims := &JwtCustomClaims{
		IdUser: userId,
		IdRole: roleID,
	}

	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// t, _ := token.SignedString([]byte("rahasia"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte("rahasia"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &models.AuthResp{
			Message: "Failed to create token",
			Status:  false,
		})
	}

	return c.JSON(http.StatusOK, &models.AuthResp{
		Message: "Berhasil login",
		Status:  true,
		Token:   t,
	})

}

func RegisterReader(c echo.Context) error {
	e := echo.New()
	e.Validator = &utils.CustomValidator{Validator: validator.New()}

	var userRegister models.RegisterParam

	c.Bind(&userRegister)
	err := c.Validate(&userRegister)

	if err == nil {
		registerErr := service.RegisterReader(userRegister)
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

func RegisterAuthor(c echo.Context) error {
	e := echo.New()
	e.Validator = &utils.CustomValidator{Validator: validator.New()}

	var userRegister models.RegisterParam

	c.Bind(&userRegister)
	err := c.Validate(&userRegister)

	if err == nil {
		registerErr := service.RegisterAuthor(userRegister)
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
