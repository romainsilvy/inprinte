package crud

import (
	"encoding/json"
	"inprinteBackoffice/structures"
	utils "inprinteBackoffice/utils"
	"net/http"

	"github.com/gorilla/mux"
)

func DeleteOne(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		//create cors header
		utils.SetCorsHeaders(&w)

		//connect the database
		db := utils.DbConnect()

		//get url values
		vars := mux.Vars(r)
		id_commandLine := vars["id_commandLine"]

		//create the sql query
		sqlQuery := ("DELETE FROM command_line WHERE id = " + id_commandLine + ";")

		_, err := db.Exec(sqlQuery)
		utils.CheckErr(err)

		//close the database connection
		db.Close()

		//create the json response
		json.NewEncoder(w).Encode(structures.InsertOneCommandLines{
			Type: "success",
		})
	}
}
