package crud

import (
	"encoding/json"
	"fmt"
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
	var role string

	//connect the database
	db := utils.DbConnect()

	//get url values
	vars := mux.Vars(r)
	id_role := vars["id_role"]

	//create the sql query
	sqlQuery := ("SELECT id, role FROM role WHERE role.id = " + id_role + " ;")
	fmt.Println("SELECT id, role FROM role WHERE role.id = " + id_role + " ;")
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
