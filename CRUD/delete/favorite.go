package CRUD

import (
	databaseTools "inprinte/backend/database"
	"inprinte/backend/utils"
	"log"
	"net/http"
)

func DeleteFavorite(w http.ResponseWriter, r *http.Request) {
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

	db := databaseTools.DbConnect()

	_, err := db.Exec(`DELETE FROM favorite WHERE favorite.id_product = ? AND favorite.id_user = ?`, id_product[0], id_user[0])

	utils.CheckErr(err)
}
