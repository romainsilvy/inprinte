package crud

import (
	"encoding/json"
	"inprinteBackoffice/structures"
	"inprinteBackoffice/utils"
	"net/http"
	"strconv"
)

func Update(w http.ResponseWriter, r *http.Request) {
	//set cors header
	utils.SetCorsHeaders(&w)

	if r.Method == "PUT" {
		// parse json from put Request
		var oneCommand structures.GetCommandLine
		err := json.NewDecoder(r.Body).Decode(&oneCommand)
		utils.CheckErr(err)

		// connect the database
		db := utils.DbConnect()

		// create the sql query
		sqlQuery := ("UPDATE command_line SET state = '" + oneCommand.Status + "' WHERE id = " + strconv.Itoa(oneCommand.Id) + ";")

		// execute the sql query
		_, err = db.Exec(sqlQuery)
		utils.CheckErr(err)

		// close the database connection
		db.Close()

		// create the json response
		json.NewEncoder(w).Encode(oneCommand)
	}
}
