package crud

import (
	"database/sql"
	"inprinte/backend/utils"
	"log"
	"strconv"
)

func InsertIntoProductFile(db *sql.DB, id_product int, url string) {
	sql := "INSERT INTO product_file (id_product, url) VALUES (\"" + strconv.Itoa(id_product) + "\", \"" + url + "\");"
	_, err := db.Exec(sql)
	log.Println(sql)
	utils.CheckErr(err)
}
