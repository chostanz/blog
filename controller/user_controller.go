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
			Code:    404,
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
		return echo.NewHTTPError(http.StatusBadRequest, &models.Response{
			Code:    400,
			Message: "Data invalid",
			Status:  false,
		})
	}

	_, editErr := service.EditUser(editUser, id)
	if editErr != nil {
		fmt.Println("Error editing user:", editErr)
		return c.JSON(http.StatusBadRequest, &models.Response{
			Code:    400,
			Message: "Gagal memperbarui pengguna",
			Status:  false,
		})
	}

	return c.JSON(http.StatusOK, &models.Response{
		Code:    200,
		Message: "Berhasil update",
		Status:  true,
	})
}

func UserRoleUpdate(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, &models.Response{
			Code:    400,
			Message: "ID pengguna tidak valid",
			Status:  false,
		})
	}

	var editRole models.Role
	c.Bind(&editRole)
	_, err = service.EditUserRole(editRole, userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &models.Response{
			Code:    400,
			Message: "Gagal memperbarui peran pengguna",
			Status:  false,
		})
	}

	return c.JSON(http.StatusOK, &models.Response{
		Code:    200,
		Message: "Berhasil update role",
		Status:  true,
	})
}

func UserDelete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var deleteUser models.Users
	_, err := service.DeleteUser(deleteUser, id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, &models.Response{
			Code:    500,
			Message: "Terjadi kesalahan internal",
			Status:  false,
		})

	}
	return c.JSON(http.StatusOK, &models.Response{
		Code:    200,
		Message: "Berhasil menghapus user",
		Status:  true,
	})
}
