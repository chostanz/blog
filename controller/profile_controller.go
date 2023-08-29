package controller

import (
	"blog/models"
	"blog/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func ProfileUpdate(c echo.Context) error {
	idUser := c.Get("id_user").(int) // Mengambil ID User

	var UserProfile models.Profile
	c.Bind(&UserProfile)
	err := c.Validate(&UserProfile)

	if err == nil {
		_, registerErr := service.EditProfile(UserProfile, idUser)
		if registerErr != nil {

			return echo.NewHTTPError(http.StatusBadRequest, "raiso")
		}
		return c.JSON(http.StatusCreated, &models.RegisterResp{
			Message: "Berhasil update",
			Status:  true,
		})
	}

	return echo.NewHTTPError(http.StatusBadRequest, err.Error())
}

func GetSpecProfile(c echo.Context) error {
	idUser := c.Get("id_user").(int) // Mengambil ID User dari konteks

	var userProfile models.Profile

	userProfile, err := service.GetProfile(idUser)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "tidak ada")
	}
	return c.JSON(http.StatusOK, userProfile)
}

func PasswordUpdate(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request data")
	}

	var passUpdate models.ChangePasswordRequest
	if err := c.Bind(&passUpdate); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request data")
	}

	if err := c.Validate(&passUpdate); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Validation error")
	}

	errS := service.EditPassword(passUpdate, id)
	if errS != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "An error occurred")
	}

	return c.JSON(http.StatusOK, &models.Response{
		Message: "Password updated successfully",
		Status:  true,
	})

}
