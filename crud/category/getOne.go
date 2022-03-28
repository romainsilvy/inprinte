package crud

import (
	"encoding/json"
	"inprinteBackoffice/structures"
	"inprinteBackoffice/utils"
	"net/http"

	"github.com/gorilla/mux"
)

func GetOne(w http.ResponseWriter, r *http.Request) {
	//create cors header
	utils.SetCorsHeaders(&w)

	//global vars
	var id int
	var name string

	//connect the database
	db := utils.DbConnect()

	//get url values
	vars := mux.Vars(r)
	id_category := vars["id_category"]

	//create the sql query
	sqlQuery := ("SELECT id, name FROM category WHERE category.id = " + id_category + " ;")

	//execute the sql query
	row := db.QueryRow(sqlQuery)

	//parse the query
	//retrieve the values and check errors
	err := row.Scan(&id, &name)
	utils.CheckErr(err)

	//add the values to the response
	oneCategory := structures.OneCategory{
		Id:   id,
		Name: name,
	}

	//create the json response
	json.NewEncoder(w).Encode(oneCategory)
}
