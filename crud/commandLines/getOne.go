package crud

import (
	"encoding/json"
	structures "inprinteBackoffice/structures"
	utils "inprinteBackoffice/utils"

	"net/http"

	"github.com/gorilla/mux"
)

func GetCommandLine(w http.ResponseWriter, r *http.Request) {
	//set cors header
	utils.SetCorsHeaders(&w)

	//global vars
	var oneCommandLines []structures.GetCommandLine
	var state string
	var id int

	//get the id from the url
	vars := mux.Vars(r)
	id_commandLine := vars["id_commandLine"]

	//connect the database
	db := utils.DbConnect()

	//create the sql query
	sqlQuery := "SELECT command_line.id, command_line.state FROM command_line WHERE command_line.id = " + id_commandLine + ";"
	rows, err := db.Query(sqlQuery)
	utils.CheckErr(err)

	//parse the query
	for rows.Next() {
		//retrieve the values and check errors
		err = rows.Scan(&id, &state)
		utils.CheckErr(err)

		//add the values to the response
		oneCommandLines = append(oneCommandLines, structures.GetCommandLine{
			Id:     id,
			Status: state,
		})
	}

	//create the json response
	utils.SetXTotalCountHeader(&w, len(oneCommandLines))
	json.NewEncoder(w).Encode(oneCommandLines)
}
