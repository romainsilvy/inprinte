package crud

import (
	"encoding/json"
	"inprinteBackoffice/structures"
	"inprinteBackoffice/utils"
	"net/http"

	"github.com/gorilla/mux"
)

func GetCategory(w http.ResponseWriter, r *http.Request) {
	//create cors header
	utils.SetCorsHeaders(&w)

	//global vars
	var id int
	var name string
	var is_alive bool

	if r.Method == "GET" {
		//connect the database
		db := utils.DbConnect()

		//get url values
		vars := mux.Vars(r)
		id_category := vars["id_category"]

		//create the sql query
		sqlQuery := ("SELECT id, name, is_alive FROM category WHERE category.id = " + id_category + " ;")

		//execute the sql query
		row := db.QueryRow(sqlQuery)

		//parse the query
		err := row.Scan(&id, &name, &is_alive)
		utils.CheckErr(err)

		//add the values to the response
		oneCategory := structures.GetCategory{
			Id:      id,
			Name:    name,
			IsAlive: is_alive,
		}

		//close the database connection
		db.Close()

		//create the json response
		json.NewEncoder(w).Encode(oneCategory)
	}
}
