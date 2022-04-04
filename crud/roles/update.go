package crud

import (
	"encoding/json"
	"inprinteBackoffice/structures"
	"inprinteBackoffice/utils"
	"net/http"
	"strconv"
)

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		//create cors header
		utils.SetCorsHeaders(&w)

		// parse json from put Request
		var oneRole structures.GetRole
		err := json.NewDecoder(r.Body).Decode(&oneRole)
		utils.CheckErr(err)

		// connect the database
		db := utils.DbConnect()

		// create the sql query
		sqlQuery := ("UPDATE role SET role = '" + oneRole.Role + "' WHERE id = " + strconv.Itoa(oneRole.Id) + ";")

		// execute the sql query
		_, err = db.Exec(sqlQuery)
		utils.CheckErr(err)

		// close the database connection
		db.Close()

		// create the json response
		json.NewEncoder(w).Encode(oneRole)
	}
}
