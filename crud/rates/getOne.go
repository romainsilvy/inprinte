package crud

import (
	"encoding/json"
	structures "inprinteBackoffice/structures"
	utils "inprinteBackoffice/utils"
	"net/http"

	"github.com/gorilla/mux"
)

func GetOne(w http.ResponseWriter, r *http.Request) {
	//global vars
	var oneRate structures.AllRates
	var firstname, lastname, name string
	var id, id_product, id_user, stars_number int
	vars := mux.Vars(r)
	id_rate := vars["id_rate"]

	//connect the database
	db := utils.DbConnect()

	//create cors header
	utils.SetCorsHeaders(&w)

	sqlQuery := "select rate.id, product.id, product.name, user.id AS id_user, user.first_name, user.last_name, rate.stars_number FROM rate INNER JOIN product ON product.id = rate.id_product INNER JOIN user ON user.id = rate.id_user WHERE rate.id = " + id_rate + ";"

	//execute the sql query and check errors
	row := db.QueryRow(sqlQuery)

	err := row.Scan(&id, &id_product, &name, &id_user, &firstname, &lastname, &stars_number)
	utils.CheckErr(err)

	oneRate = structures.AllRates{
		Id:           id,
		Id_product:   id_product,
		Name:         name,
		Id_user:      id_user,
		Firstname:    firstname,
		Lastname:     lastname,
		Stars_number: stars_number,
	}

	//create the json response
	utils.SetXTotalCountHeader(&w, 1)
	json.NewEncoder(w).Encode(oneRate)
}
