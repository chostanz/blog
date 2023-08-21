package controller

import (
	"blog/models"
	"blog/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func ProfileUpdate(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var UserProfile models.Profile
	c.Bind(&UserProfile)
	err := c.Validate(&UserProfile)

	if err == nil {
		_, registerErr := service.EditProfile(UserProfile, id)
		if registerErr != nil {

			return echo.NewHTTPError(http.StatusBadRequest, "raiso")
		}
		return c.JSON(http.StatusCreated, &models.RegisterResp{
			Message: "Berhasil update",
			Status:  true,
		})
	}

	return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	// if errValidate := c.Validate(&UserProfile); errValidate != nil {
	// 	return echo.NewHTTPError(http.StatusBadRequest, "Data tidak boleh kosong")
	// }

	// _, errorUpdate := service.Profile(UserProfile, id)
	// if errorUpdate != nil {
	// 	return echo.NewHTTPError(http.StatusBadRequest, "gabisa bang")
	// }

	// return c.JSON(http.StatusOK, &models.Response{
	// 	Message: "Data sukses diupdate",
	// 	Status:  true,
	// })

}

func GetSpecProfile(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var userProfile models.Profile

	userProfile, err := service.GetProfile(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "tidak ada")
	}
	return c.JSON(http.StatusOK, userProfile)
}

func PasswordUpdate(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var passUpdate models.Password
	c.Bind(&passUpdate)
	err := c.Validate(&passUpdate)

	if err == nil {
		_, registerErr := service.EditPassword(passUpdate, id)
		if registerErr != nil {

			return echo.NewHTTPError(http.StatusBadRequest, "raiso")
		}
		return c.JSON(http.StatusCreated, &models.Response{
			Message: "Berhasil update",
			Status:  true,
		})
	}

	return echo.NewHTTPError(http.StatusBadRequest, err.Error())

}
