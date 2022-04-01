package crud

import (
	"database/sql"
	"inprinte/backend/utils"
	"log"
	"strconv"
)

func InsertIntoPicture(db *sql.DB, url string, default_picture, non_product_picture bool) {
	sql := "INSERT INTO picture (url, `default`, non_product_picture) VALUES (\"" + url + "\", " + strconv.FormatBool(default_picture) + ", " + strconv.FormatBool(non_product_picture) + ");"
	_, err := db.Exec(sql)
	log.Println(sql)
	utils.CheckErr(err)
}
