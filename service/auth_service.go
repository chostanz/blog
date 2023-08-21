package service

import (
	"blog/database"
	"blog/models"

	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB = database.Koneksi()

func CheckCredential(userLogin models.LoginParam) (int, bool) {
	var isAuthentication bool
	var id int

	rows, err := db.Query("SELECT CASE WHEN COUNT(*) > 0 THEN 'true' ELSE 'false' END FROM users WHERE username = $1 AND password = $2", userLogin.Username, userLogin.Password)
	if err != nil {
		return 0, false
	}

	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&isAuthentication)
		if err != nil {
			return 0, false
		}
	}

	rows, err = db.Query("SELECT id from users where username = $1 AND password = $2", userLogin.Username, userLogin.Password)
	if err != nil {
		return 0, false
	}

	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			return 0, false
		}
	}
	return id, isAuthentication
}

func RegisterReader(userRegister models.RegisterParam) error {
	_, err := db.NamedExec("INSERT INTO users (username, email, password) VALUES (:username, :email, :password)", userRegister)
	if err != nil {
		return err
	}

	var userID int
	err = db.Get(&userID, "SELECT id FROM users WHERE username = $1", userRegister.Username)
	if err != nil {
		return err
	}

	// Insert into users_roles table with default role_id
	_, err = db.Exec("INSERT INTO users_roles (role_id, user_id) VALUES ($1, $2)", 3, userID) // 3 is the default role_id for 'reader'
	if err != nil {
		return err
	}

	return nil
}

func RegisterAuthor(userRegister models.RegisterParam) error {
	_, err := db.NamedExec("INSERT INTO users (username, email, password) VALUES (:username, :email, :password)", userRegister)
	if err != nil {
		return err
	}

	var userID int
	err = db.Get(&userID, "SELECT id FROM users WHERE username = $1", userRegister.Username)
	if err != nil {
		return err
	}

	// Insert into users_roles table with default role_id
	_, err = db.Exec("INSERT INTO users_roles (role_id, user_id) VALUES ($1, $2)", 2, userID) // 3 is the default role_id for 'reader'
	if err != nil {
		return err
	}

	return nil
}
