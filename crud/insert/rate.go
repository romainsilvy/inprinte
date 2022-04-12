package crud

import (
	"database/sql"
	"inprinte/backend/utils"
	"log"
	"strconv"
)

func InsertIntoRate(db *sql.DB, id_product, id_user, stars_number int) {
	sql := "INSERT INTO rate (id_product,id_user,stars_number) VALUES (\"" + strconv.Itoa(id_product) + "\", \"" + strconv.Itoa(id_user) + "\", \"" + strconv.Itoa(stars_number) + "\");"
	_, err := db.Exec(sql)
	log.Println(sql)
	utils.CheckErr(err)
}
