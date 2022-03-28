package CRUD

import (
	"database/sql"
	"inprinte/backend/utils"
	"log"
	"strconv"
)

func InsertIntoCommand(db *sql.DB, id_user int, date, state string) {
	sql := "INSERT INTO `command` (id_user,date,state) VALUES(\"" + strconv.Itoa(id_user) + "\", \"" + date + "\", \"" + state + "\");"
	_, err := db.Exec(sql)
	log.Println(sql)
	utils.CheckErr(err)
}