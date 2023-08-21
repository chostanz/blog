package utils

import (
	"encoding/json"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type JwtCustomClaims struct {
	IdUser string `json:"id_user"`
	jwt.RegisteredClaims
}

func GetIdUserJWT(c echo.Context) string {

	claims := JwtCustomClaims{}

	user := c.Get("users").(*jwt.Token)
	tmp, _ := json.Marshal(user.Claims)
	_ = json.Unmarshal(tmp, &claims)

	return claims.IdUser
}
