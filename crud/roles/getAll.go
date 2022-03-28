package crud

import (
	"encoding/json"
	structures "inprinteBackoffice/structures"
	utils "inprinteBackoffice/utils"

	"net/http"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	//global vars
	var Roles []structures.Roles

	//connect the database
	db := utils.DbConnect()

	//create cors header
	utils.SetCorsHeaders(&w)

	sqlQuery := "SELECT id, role FROM role"
	//execute the sql query and check errors
	rows, err := db.Query(sqlQuery)
	utils.CheckErr(err)

	//parse the query
	for rows.Next() {
		//global vars
		var role string
		var id int

		//retrieve the values and check errors
		err = rows.Scan(&id, &role)
		utils.CheckErr(err)

		//add the values to the response
		Roles = append(Roles, structures.Roles{
			Id:   id,
			Role: role,
		})
	}
	//close the rows

	//create the json response
	utils.SetXTotalCountHeader(&w, len(Roles))
	json.NewEncoder(w).Encode(Roles)
}
