package CRUD

import (
	databaseTools "inprinte/backend/database"
	"inprinte/backend/utils"
	"net/http"

	"github.com/gorilla/mux"
)

func DeleteFavorite(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id_product := params["id_product"]
	db := databaseTools.DbConnect()
	id_user := "L'ID USER DU COOKIE QUI FAIT LA REQUEST"

	_, err := db.Exec("DELETE FROM favorite WHERE favorite.id_product = $1 AND favorite.id_user = $2", id_product, id_user)

	utils.CheckErr(err)
}
