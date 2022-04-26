package crud

import (
	"database/sql"
	"inprinte/backend/utils"
	"log"
)

func InsertIntoAddress(db *sql.DB, street, city, state, country, zipCode string) {
	sql := "INSERT INTO address (street,city,state,country,zip_code) VALUES (\"" + street + "\", \"" + city + "\", \"" + state + "\", \"" + country + "\", \"" + zipCode + "\");"
	_, err := db.Exec(sql)
	log.Println(sql)
	utils.CheckErr(err)
}
