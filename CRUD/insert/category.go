package CRUD

import (
	"database/sql"
	"inprinte/backend/utils"
	"log"
)


func InsertIntoCategory(db *sql.DB, name string) {
	sql := "INSERT INTO category (name) VALUES (\"" + name + "\");"
	_, err := db.Exec(sql)
	log.Println(sql)
	utils.CheckErr(err)
}