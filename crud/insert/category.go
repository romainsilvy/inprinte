package crud

import (
	"database/sql"
	"inprinte/backend/utils"
	"log"
	"strconv"
)

func InsertIntoCategory(db *sql.DB, name string, is_alive bool) {
	sql := "INSERT INTO category (name, is_alive) VALUES (\"" + name + "\", " + strconv.FormatBool(is_alive) + ");"
	_, err := db.Exec(sql)
	log.Println(sql)
	utils.CheckErr(err)
}
