package crud

import (
	"encoding/json"
	"inprinteBackoffice/structures"
	utils "inprinteBackoffice/utils"
	"net/http"

	"github.com/gorilla/mux"
)

func DeleteOne(w http.ResponseWriter, r *http.Request) {
	//create cors header
	utils.SetCorsHeaders(&w)

	if r.Method == "DELETE" {

		//connect the database
		db := utils.DbConnect()

		//get url values
		vars := mux.Vars(r)
		id_user := vars["id_user"]

		//create the sqlQuery
		sqlQuery := ("UPDATE user SET is_alive = false WHERE id = " + id_user + ";")
		_, err := db.Exec(sqlQuery)
		utils.CheckErr(err)

		//close the database connection
		db.Close()

		//create the json response
		json.NewEncoder(w).Encode(structures.InsertOneUser{
			Type: "success",
		})
	}
}
