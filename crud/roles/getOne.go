package crud

import (
	"encoding/json"
	"inprinteBackoffice/structures"
	"inprinteBackoffice/utils"
	"net/http"
)

func GetOne(w http.ResponseWriter, r *http.Request) {
	//global vars
	var id int
	var role string

	//connect the database
	db := utils.DbConnect()

	//create cors header
	utils.SetCorsHeaders(&w)

	//create the sql query
	sqlQuery := ("SELECT id, role FROM role ;")

	//execute the sql query
	row := db.QueryRow(sqlQuery)

	//parse the query
	//retrieve the values and check errors
	err := row.Scan(&id, &role)
	utils.CheckErr(err)

	//add the values to the response
	oneRole := structures.OneRole{
		Id:   id,
		Role: role,
	}

	//create the json response
	json.NewEncoder(w).Encode(oneRole)
}
