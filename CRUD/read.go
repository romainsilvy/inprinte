package CRUD

import (
	"log"
	"database/sql"

	databaseTools "inprinte/backend/database"
	_ "github.com/go-sql-driver/mysql"
)

func Read() {
	db := databaseTools.DbConnect();
	table := "users"

	switch table {
	case "users":
		readUser(1, db)
	}
}

func readUser(id_user int, db *sql.DB) {
	_, err := db.Exec(`SELECT * FROM users WHERE id = ?`, id_user)
	if err != nil {
		log.Fatal(err)
	}
}