package crud

import (
	"encoding/json"
	"inprinteBackoffice/structures"
	utils "inprinteBackoffice/utils"
	"net/http"

	"github.com/gorilla/mux"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	//create cors header
	utils.SetCorsHeaders(&w)

	if r.Method == "DELETE" {

		if utils.Securized(w, r) {
			//connect the database
			db := utils.DbConnect()

			//get url values
			vars := mux.Vars(r)
			id_role := vars["id_role"]

			//create the sql query
			sqlQuery := ("DELETE FROM role WHERE id = " + id_role + ";")

			//execute the sql query
			_, err := db.Exec(sqlQuery)
			utils.CheckErr(err)

			//close the database connection
			db.Close()

			//create the json response
			json.NewEncoder(w).Encode(structures.JsonResponseRole{
				Type:    "success",
				Message: "Role deleted",
			})
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
	}
}
