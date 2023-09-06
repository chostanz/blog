package database

import (
	"log"

	sqlx "github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB = Koneksi()

func Koneksi() *sqlx.DB {
	db, err := sqlx.Connect("postgres", "user=postgres password=00000 dbname=db_blog sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}
	return db
}
