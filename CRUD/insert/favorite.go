package CRUD

import (
	"database/sql"
	"inprinte/backend/utils"
	"log"
	"strconv"
)

func InsertIntoFavorite(db *sql.DB, id_product, id_user int) {
	sql := "INSERT INTO favorite (id_product,id_user) VALUES (\"" + strconv.Itoa(id_product) + "\", \"" + strconv.Itoa(id_user) + "\");"
	_, err := db.Exec(sql)
	log.Println(sql)
	utils.CheckErr(err)
}