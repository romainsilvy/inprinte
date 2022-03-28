package CRUD

import (
	"database/sql"
	"inprinte/backend/utils"
	"log"
	"strconv"
)

func InsertIntoUser(db *sql.DB, first_name, last_name, email, password, phone string, is_alive bool, id_address, id_role int) {
	sql := "INSERT INTO user (first_name,last_name,email,password,phone,is_alive, id_address, id_role) VALUES (\"" + first_name + "\", \"" + last_name + "\", \"" + email + "\", \"" + password + "\", \"" + phone + "\", " + strconv.FormatBool(is_alive) + ", \"" + strconv.Itoa(id_address) + "\", \"" + strconv.Itoa(id_role) + "\");"

	_, err := db.Exec(sql)
	log.Println(sql)
	utils.CheckErr(err)
}
