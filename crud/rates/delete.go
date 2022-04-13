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
			id_rate := vars["id_rate"]

			//create the sqlQuery
			sqlQuery := ("DELETE FROM rate WHERE id = " + id_rate + ";")

			//execute the sql query
			_, err := db.Exec(sqlQuery)
			utils.CheckErr(err)

			//close the database connection
			db.Close()

			//create the json response
			json.NewEncoder(w).Encode(structures.JsonResponseRate{
				Type:    "success",
				Message: "Rate deleted",
			})
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
	}
}
