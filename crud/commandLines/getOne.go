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
	var state string
	var id int

	if r.Method == "GET" {
		//get the id from the url
		vars := mux.Vars(r)
		id_commandLine := vars["id_commandLine"]

		//connect the database
		db := utils.DbConnect()

		//create the sql query
		sqlQuery := ("SELECT command_line.id, command_line.state FROM command_line WHERE command_line.id = " + id_commandLine + ";")

		//execute the sql query
		row := db.QueryRow(sqlQuery)

		//parse the query
		err := row.Scan(&id, &state)
		utils.CheckErr(err)

		commandLine := structures.GetCommandLine{
			Id:     id,
			Status: state,
		}

		//close the database connection
		db.Close()

		//create the json response
		utils.SetXTotalCountHeader(&w, 1)
		json.NewEncoder(w).Encode(commandLine)
	}
}
