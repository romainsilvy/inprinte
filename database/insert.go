package databaseTools

import (
	"database/sql"
	utils "inprinte/backend/utils"
	"log"
	"strconv"
)

func InsertIntoRole(db *sql.DB, roleuser string) {
	sql := "INSERT INTO role (role) VALUES (\"" + roleuser + "\");"
	_, err := db.Exec(sql)
	log.Println(sql)
	utils.CheckErr(err)
}

func InsertIntoCategory(db *sql.DB, name string) {
	sql := "INSERT INTO category (name) VALUES (\"" + name + "\");"
	_, err := db.Exec(sql)
	log.Println(sql)
	utils.CheckErr(err)
}

func InsertIntoAddress(db *sql.DB, street, city, state, country, zipCode string) {
	sql := "INSERT INTO address (street,city,state,country,zip_code) VALUES (\"" + street + "\", \"" + city + "\", \"" + state + "\", \"" + country + "\", \"" + zipCode + "\");"
	_, err := db.Exec(sql)
	log.Println(sql)
	utils.CheckErr(err)
}

func InsertIntoUser(db *sql.DB, first_name, last_name, email, password, phone string, is_alive bool, id_address, id_role int) {
	sql := "INSERT INTO user (first_name,last_name,email,password,phone,is_alive, id_address, id_role) VALUES (\"" + first_name + "\", \"" + last_name + "\", \"" + email + "\", \"" + password + "\", \"" + phone + "\", " + strconv.FormatBool(is_alive) + ", \"" + strconv.Itoa(id_address) + "\", \"" + strconv.Itoa(id_role) + "\");"

	_, err := db.Exec(sql)
	log.Println(sql)
	utils.CheckErr(err)
}

func InsertIntoPicture(db *sql.DB, url string) {
	sql := "INSERT INTO picture (url, `default`) VALUES (\"" + url + "\", " + strconv.FormatBool(false) + ");"
	_, err := db.Exec(sql)
	log.Println(sql)
	utils.CheckErr(err)
}

func InsertIntoProduct(db *sql.DB, name, description string, price, id_category, id_user int, pending_validation, is_alive bool) {
	sql := "INSERT INTO product (name,description,price,id_category,id_user,pending_validation,is_alive)VALUES (\"" + name + "\", \"" + description + "\", \"" + strconv.Itoa(price) + "\", \"" + strconv.Itoa(id_category) + "\", \"" + strconv.Itoa(id_user) + "\", " + strconv.FormatBool(pending_validation) + ", " + strconv.FormatBool(is_alive) + ");"
	_, err := db.Exec(sql)
	log.Println(sql)
	utils.CheckErr(err)
}

func InsertIntoProductFile(db *sql.DB, id_product int, url string) {
	sql := "INSERT INTO product_file (id_product, url) VALUES (\"" + strconv.Itoa(id_product) + "\", \"" + url + "\");"
	_, err := db.Exec(sql)
	log.Println(sql)
	utils.CheckErr(err)
}

func InsertIntoProductPicture(db *sql.DB, id_product, id_picture int) {
	sql := "INSERT INTO product_picture (id_product, id_picture) VALUES (\"" + strconv.Itoa(id_product) + "\", \"" + strconv.Itoa(id_picture) + "\");"
	_, err := db.Exec(sql)
	log.Println(sql)
	utils.CheckErr(err)
}

func InsertIntoRate(db *sql.DB, id_product, id_user, stars_number int) {
	sql := "INSERT INTO rate (id_product,id_user,stars_number) VALUES (\"" + strconv.Itoa(id_product) + "\", \"" + strconv.Itoa(id_user) + "\", \"" + strconv.Itoa(stars_number) + "\");"
	_, err := db.Exec(sql)
	log.Println(sql)
	utils.CheckErr(err)
}

func InsertIntoFavorite(db *sql.DB, id_product, id_user int) {
	sql := "INSERT INTO favorite (id_product,id_user) VALUES (\"" + strconv.Itoa(id_product) + "\", \"" + strconv.Itoa(id_user) + "\");"
	_, err := db.Exec(sql)
	log.Println(sql)
	utils.CheckErr(err)
}

func InsertIntoCommandLine(db *sql.DB, id_command, id_product, quantity int, state string) {
	sql := "INSERT INTO command_line (id_command,id_product,quantity,state) VALUES (\"" + strconv.Itoa(id_command) + "\", \"" + strconv.Itoa(id_product) + "\", \"" + strconv.Itoa(quantity) + "\", \"" + state + "\");"
	_, err := db.Exec(sql)
	log.Println(sql)
	utils.CheckErr(err)
}

func InsertIntoCommand(db *sql.DB, id_user int, date, state string) {
	sql := "INSERT INTO `command` (id_user,date,state) VALUES(\"" + strconv.Itoa(id_user) + "\", \"" + date + "\", \"" + state + "\");"
	_, err := db.Exec(sql)
	log.Println(sql)
	utils.CheckErr(err)
}
