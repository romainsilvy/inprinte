package CRUD

import (
	"database/sql"
	"inprinte/backend/utils"
	"log"
	"strconv"
)

func InsertIntoPicture(db *sql.DB, url string) {
	sql := "INSERT INTO picture (url, `default`) VALUES (\"" + url + "\", " + strconv.FormatBool(false) + ");"
	_, err := db.Exec(sql)
	log.Println(sql)
	utils.CheckErr(err)
}