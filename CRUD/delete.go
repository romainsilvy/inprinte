package CRUD

import (
	"log"
	"database/sql"

	databaseTools "inprinte/backend/database"
	_ "github.com/go-sql-driver/mysql"
)

func Delete() {
	db := databaseTools.DbConnect();
	table := "users"

	switch table {
	case "users":
		deleteUser(1, db)
	}
}

func deleteUser(id_user int, db *sql.DB) {
	_, err := db.Exec(`DELETE FROM users WHERE id = ?`, id_user)
	if err != nil {
		log.Fatal(err)
	}
}