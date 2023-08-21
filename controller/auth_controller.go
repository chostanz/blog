package controller

import (
	"blog/models"
	"blog/service"
	"blog/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type TokenCheck struct {
	Token string `json:"token"`
}

type JwtCustomClaims struct {
	IdUser string `json:"id_user"`
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

	idUser, isAuthentication := service.CheckCredential(loginbody)

	if !isAuthentication {
		return c.JSON(http.StatusUnauthorized, &models.LoginResp{
			Message: "Akun tidak ada atau password salah",
			Status:  false,
		})
	}
	claims := &JwtCustomClaims{
		strconv.Itoa(idUser),
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, _ := token.SignedString([]byte("secret"))

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
