package crud

import (
	"encoding/json"
	"inprinteBackoffice/structures"
	"inprinteBackoffice/utils"
	"net/http"
	"strconv"
)

func Update(w http.ResponseWriter, r *http.Request) {
	//create cors header
	utils.SetCorsHeaders(&w)

	if r.Method == "PUT" {

		if utils.Securized(w, r) {
			// parse json from put Request
			var oneCommand structures.GetCommand

			//parse the body
			err := json.NewDecoder(r.Body).Decode(&oneCommand)
			utils.CheckErr(err)

			// connect the database
			db := utils.DbConnect()

			// create the sql query
			sqlQuery := ("UPDATE command SET state = '" + oneCommand.Status + "' WHERE id = " + strconv.Itoa(oneCommand.Id) + ";")

			// execute the sql query
			_, err = db.Exec(sqlQuery)
			utils.CheckErr(err)

			// close the database connection
			db.Close()

			// create the json response
			json.NewEncoder(w).Encode(oneCommand)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
	}
}
