package CRUD

import (
	"database/sql"
	databaseTools "inprinte/backend/database"
	"inprinte/backend/utils"
	"log"
	"net/http"
	"strconv"
)

func InsertIntoFavoriteRequest(db *sql.DB, id_product, id_user int) {
	sql := "INSERT INTO favorite (id_product,id_user) VALUES (\"" + strconv.Itoa(id_product) + "\", \"" + strconv.Itoa(id_user) + "\");"
	_, err := db.Exec(sql)
	log.Println(sql)
	utils.CheckErr(err)
}

func InsertIntoFavorite(w http.ResponseWriter, r *http.Request) {
	db := databaseTools.DbConnect()
	id_user, ok := r.URL.Query()["id_user"]
	if !ok || len(id_user[0]) < 1 {
		log.Println("Url Param 'id_user' is missing")
		return
	}

	id_product, ok := r.URL.Query()["id_product"]
	if !ok || len(id_product[0]) < 1 {
		log.Println("Url Param 'id_product' is missing")
		return
	}
	intIdProduct, _ := strconv.Atoi(id_product[0])
	intIdUser, _ := strconv.Atoi(id_user[0])

	InsertIntoFavoriteRequest(db, intIdProduct, intIdUser)
}
