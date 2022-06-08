package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:1234@tcp(localhost:3306)/todo_db")

	if err != nil {
		log.Fatal(err)

		defer db.Close()
	}

	return db
}
