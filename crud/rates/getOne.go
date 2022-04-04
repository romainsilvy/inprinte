package crud

import (
	"encoding/json"
	structures "inprinteBackoffice/structures"
	utils "inprinteBackoffice/utils"
	"net/http"

	"github.com/gorilla/mux"
)

func GetRate(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		//create cors header
		utils.SetCorsHeaders(&w)

		//global vars
		var rate structures.GetRates
		var firstname, lastname, name string
		var id, id_product, id_user, stars_number int

		vars := mux.Vars(r)
		id_rate := vars["id_rate"]

		//connect the database
		db := utils.DbConnect()
		sqlQuery := "select rate.id, product.id, product.name, user.id AS id_user, user.first_name, user.last_name, rate.stars_number FROM rate INNER JOIN product ON product.id = rate.id_product INNER JOIN user ON user.id = rate.id_user WHERE rate.id = " + id_rate + ";"

		//execute the sql query and check errors
		row := db.QueryRow(sqlQuery)

		//parse the query
		err := row.Scan(&id, &id_product, &name, &id_user, &firstname, &lastname, &stars_number)
		utils.CheckErr(err)

		//add the values to the response
		rate = structures.GetRates{
			Id:           id,
			Id_product:   id_product,
			Name:         name,
			Id_user:      id_user,
			Firstname:    firstname,
			Lastname:     lastname,
			Stars_number: stars_number,
		}

		//close the database connection
		db.Close()

		//create the json response
		utils.SetXTotalCountHeader(&w, 1)
		json.NewEncoder(w).Encode(rate)
	}
}
