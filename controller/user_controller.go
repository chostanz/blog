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
			Code:    500,
			Message: "Terjadi kesalahan internal server. Mohon coba beberapa saat lagi",
			Status:  false,
		}
		return c.JSON(http.StatusInternalServerError, response)
	}
	return c.JSON(http.StatusOK, users)
}

func GetSpecUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var getUser models.Role

	getUser, err := service.GetSpecUser(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &models.Response{
			Code:    500,
			Message: "Terjadi kesalahan internal pada server. Mohon coba beberapa saat lagi!",
			Status:  false,
		})
	}

	return c.JSON(http.StatusOK, getUser)
}

func UserUpdate(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var editUser models.UserEdit
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
			Code:    422,
			Message: "ID pengguna tidak valid",
			Status:  false,
		})
	}

	var editRole models.Role
	c.Bind(&editRole)

	errVal := c.Validate(&editRole)
	if errVal != nil {
		return c.JSON(http.StatusUnprocessableEntity, &models.Response{
			Code:    422,
			Message: "Data yang dimasukkan tidak valid",
			Status:  false,
		})

	}
	_, err = service.EditUserRole(editRole, userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &models.Response{
			Code:    500,
			Message: "Terjadi kesalahan internal pada server. Mohon coba beberapa saat lagi!",
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
			Message: "Terjadi kesalahan internal server. Mohon coba beberapa saat lagi",
			Status:  false,
		})

	}
	return c.JSON(http.StatusOK, &models.Response{
		Code:    200,
		Message: "Berhasil menghapus user",
		Status:  true,
	})
}

func GetAllRole(c echo.Context) error {
	roles, err := service.GetRole()

	if err != nil {
		response := models.Response{
			Code:    500,
			Message: "Terjadi kesalahan internal server. Mohon coba beberapa saat lagi",
			Status:  false,
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	return c.JSON(http.StatusOK, roles)
}
