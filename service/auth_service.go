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

func RegisterUser(userRegister models.RegisterParam) error {
	_, err := db.NamedExec("INSERT INTO users (username, email, password) VALUES (:username, :email, :password)", userRegister)
	if err != nil {
		return err
	}
	return nil
}
