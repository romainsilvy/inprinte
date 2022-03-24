package CRUD

import (
	"database/sql"
	"inprinte/backend/utils"
	"log"
)

func InsertIntoRole(db *sql.DB, roleuser string) {
	sql := "INSERT INTO role (role) VALUES (\"" + roleuser + "\");"
	_, err := db.Exec(sql)
	log.Println(sql)
	utils.CheckErr(err)
}