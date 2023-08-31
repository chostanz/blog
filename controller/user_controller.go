package controller

import (
	"blog/models"
	"blog/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetAllUser(c echo.Context) error {
	users, err := service.UsersAll()

	if err != nil {
		response := models.Response{
			Message: "Halaman tidak ditemukan atau url salah",
			Status:  false,
		}
		return c.JSON(http.StatusNotFound, response)
	}
	return c.JSON(http.StatusOK, users)
}

func UserUpdate(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var editUser models.Users
	c.Bind(&editUser)
	err := c.Validate(&editUser)

	if err != nil {
		fmt.Println("Validation error:", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	_, editErr := service.EditUser(editUser, id)
	if editErr != nil {
		fmt.Println("Error editing user:", editErr)
		return c.JSON(http.StatusBadRequest, editErr.Error())
	}

	return c.JSON(http.StatusOK, &models.Response{
		Message: "Berhasil update",
		Status:  true,
	})
}

func UserRoleUpdate(c echo.Context) error {
	userID, _ := strconv.Atoi(c.Param("id"))
	roleID, _ := strconv.Atoi(c.Param("role_id"))

	err := service.EditUserRole(userID, roleID)
	if err != nil {
		fmt.Println("Error editing user role:", err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, &models.Response{
		Message: "Berhasil update role",
		Status:  true,
	})
}

func UserDelete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var deleteUser models.Users
	_, err := service.DeleteUser(deleteUser, id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Terjadi kesalahan internal")

	}
	return c.JSON(http.StatusOK, &models.Response{
		Message: "Berhasil meghapus user",
		Status:  true,
	})
}
