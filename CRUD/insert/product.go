package CRUD

import (
	"database/sql"
	"inprinte/backend/utils"
	"log"
	"strconv"
)

func InsertIntoProduct(db *sql.DB, name, description string, price, id_category, id_user int, pending_validation, is_alive bool) {
	sql := "INSERT INTO product (name,description,price,id_category,id_user,pending_validation,is_alive)VALUES (\"" + name + "\", \"" + description + "\", \"" + strconv.Itoa(price) + "\", \"" + strconv.Itoa(id_category) + "\", \"" + strconv.Itoa(id_user) + "\", " + strconv.FormatBool(pending_validation) + ", " + strconv.FormatBool(is_alive) + ");"
	_, err := db.Exec(sql)
	log.Println(sql)
	utils.CheckErr(err)
}