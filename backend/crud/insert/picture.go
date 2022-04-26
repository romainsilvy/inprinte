package crud

import (
	"database/sql"
	"inprinte/backend/utils"
	"log"
	"strconv"
)

func InsertIntoPicture(db *sql.DB, url string, default_picture bool) {
	sql := "INSERT INTO picture (url, `default`) VALUES (\"" + url + "\", " + strconv.FormatBool(default_picture) + ");"
	_, err := db.Exec(sql)
	log.Println(sql)
	utils.CheckErr(err)
}
