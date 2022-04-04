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
		var oneCommand structures.OneCommand
		err := json.NewDecoder(r.Body).Decode(&oneCommand)
		utils.CheckErr(err)

		// connect the database
		db := utils.DbConnect()

		// create the sql query
		sqlQuery := ("UPDATE command_line SET state = '" + oneCommand.Status + "' WHERE id = " + strconv.Itoa(oneCommand.Id) + ";")
		_, err = db.Exec(sqlQuery)
		utils.CheckErr(err)

		// create the json response
		json.NewEncoder(w).Encode(oneCommand)
	}
}
