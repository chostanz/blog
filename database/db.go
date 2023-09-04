package database

import (
	"log"

	"github.com/jinzhu/gorm"
	sqlx "github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB = Koneksi()
var dbGorm = DbGorm()

func Koneksi() *sqlx.DB {
	db, err := sqlx.Connect("postgres", "user=postgres password=12345 dbname=db_blog sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}
	return db
}

func DbGorm() *gorm.DB {
	dbgorm, err := gorm.Open("postgres", "user=postgres password=12345 dbname=db_blog sslmode=disable")
	if err != nil {
		panic("Failed to connect to database")
	}
	return dbgorm
}
