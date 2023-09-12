package service

import (
	"blog/models"
	"fmt"
	"strconv"
)

func UsersAll() ([]models.Users, error) {
	usersGet := []models.Users{}

	rows, err := db.Queryx("SELECT u.id, u.username, u.password, u.email, r.role FROM users u JOIN users_roles ur ON u.id = ur.user_id JOIN roles r ON ur.role_id = r.id")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		place := models.Users{}
		rows.StructScan(&place)
		usersGet = append(usersGet, place)
	}
	return usersGet, nil
}

func EditUser(editUser models.Users, id int) (models.Users, error) {
	idStr := strconv.Itoa(id)

	_, err := db.NamedExec("UPDATE users SET username = :username, email = :email, password = :password WHERE id = :id", map[string]interface{}{
		"username": editUser.Username,
		"password": editUser.Password,
		"email":    editUser.Email,
		"id":       idStr,
	})
	if err != nil {
		fmt.Println("Error updating user:", err)
		return models.Users{}, err
	}

	return editUser, nil
}

func EditUserRole(editRole models.Role, userID int) (models.Role, error) {
	idStr := strconv.Itoa(userID)

	_, err := db.NamedExec("UPDATE users_roles SET role_id = :role_id WHERE user_id = :user_id", map[string]interface{}{
		"role_id": editRole.RoleID,
		"user_id": idStr,
	})
	if err != nil {
		fmt.Println("Error updating user role:", err)
		return models.Role{}, err
	}
	return editRole, nil
}

func DeleteUser(deleteUser models.Users, id int) (models.Users, error) {
	idStr := strconv.Itoa(id)

	_, err := db.Exec("DELETE FROM users WHERE id = $1", idStr)
	if err != nil {
		return deleteUser, err
	}
	return deleteUser, nil
}
