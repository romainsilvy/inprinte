package crud

import (
	"encoding/json"
	structures "inprinteBackoffice/structures"
	utils "inprinteBackoffice/utils"

	"net/http"

	"github.com/gorilla/mux"
)

func GetOne(w http.ResponseWriter, r *http.Request) {
	utils.SetCorsHeaders(&w)

	//get the id from the url
	vars := mux.Vars(r)
	id_commandLine := vars["id_commandLine"]

	toReturn := ReturnData(id_commandLine)

	var response = structures.JsonReponseCommandLines{
		Data: toReturn,
	}

	//create the json response
	utils.SetXTotalCountHeader(&w, 1)
	json.NewEncoder(w).Encode(response)
}

func ReturnData(id_commandLine string) []structures.OneCommandLines {
	//global vars
	var oneCommandLines []structures.OneCommandLines
	var state string
	var id int

	//connect the database
	db := utils.DbConnect()

	//create the sql query
	sqlQuery := "SELECT command_line.id, command_line.state FROM command_line WHERE command_line.id = " + id_commandLine + ";"

	//execute the sql query and check errors
	rows, err := db.Query(sqlQuery)
	utils.CheckErr(err)

	//parse the query
	for rows.Next() {

		//retrieve the values and check errors
		err = rows.Scan(&id, &state)
		utils.CheckErr(err)

		//add the values to the response
		oneCommandLines = append(oneCommandLines, structures.OneCommandLines{
			Id:     id,
			Status: state,
		})
	}
	return oneCommandLines
}
