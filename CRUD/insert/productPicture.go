package CRUD

import (
	"database/sql"
	"inprinte/backend/utils"
	"log"
	"strconv"
)

func InsertIntoProductPicture(db *sql.DB, id_product, id_picture int) {
	sql := "INSERT INTO product_picture (id_product, id_picture) VALUES (\"" + strconv.Itoa(id_product) + "\", \"" + strconv.Itoa(id_picture) + "\");"
	_, err := db.Exec(sql)
	log.Println(sql)
	utils.CheckErr(err)
}