package crud

import (
	"database/sql"
	"inprinte/backend/utils"
	"log"
	"strconv"
)

func InsertIntoCommandLine(db *sql.DB, id_command, id_product, quantity int, state string) {
	sql := "INSERT INTO command_line (id_command,id_product,quantity,state) VALUES (\"" + strconv.Itoa(id_command) + "\", \"" + strconv.Itoa(id_product) + "\", \"" + strconv.Itoa(quantity) + "\", \"" + state + "\");"
	_, err := db.Exec(sql)
	log.Println(sql)
	utils.CheckErr(err)
}
