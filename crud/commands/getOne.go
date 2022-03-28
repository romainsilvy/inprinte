package crud

import (
	"encoding/json"
	"inprinteBackoffice/structures"
	"inprinteBackoffice/utils"
	"net/http"

	"github.com/gorilla/mux"
)

func GetOne(w http.ResponseWriter, r *http.Request) {
	//global vars
	var id int
	var status string

	//connect the database
	db := utils.DbConnect()

	//create cors header
	utils.SetCorsHeaders(&w)

	//get url values
	vars := mux.Vars(r)
	id_command := vars["id_command"]

	//create the sql query
	sqlQuery := ("SELECT  command.id AS id, command.state FROM command WHERE command.id = " + id_command + ";")

	//execute the sql query
	row := db.QueryRow(sqlQuery)

	//parse the query
	//retrieve the values and check errors
	err := row.Scan(&id, &status)
	utils.CheckErr(err)

	//add the values to the response
	oneCommand := structures.OneCommand{
		Id:     id,
		Status: status,
	}

	//create the json response
	json.NewEncoder(w).Encode(oneCommand)
}
