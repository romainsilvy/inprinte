package crud

import (
	"encoding/json"
	"inprinteBackoffice/structures"
	utils "inprinteBackoffice/utils"
	"net/http"

	"github.com/gorilla/mux"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		//create cors header
		utils.SetCorsHeaders(&w)

		//connect the database
		db := utils.DbConnect()

		//get url values
		vars := mux.Vars(r)
		id_user := vars["id_user"]

		//create the sqlQuery
		sqlQuery := ("UPDATE user SET is_alive = false WHERE id = " + id_user + ";")

		//execute the sql query
		_, err := db.Exec(sqlQuery)
		utils.CheckErr(err)

		//close the database connection
		db.Close()

		//create the json response
		json.NewEncoder(w).Encode(structures.JsonResponseUser{
			Type:    "success",
			Message: "User is_alive has been set to false",
		})
	}
}
