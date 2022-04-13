package crud

import (
	"encoding/json"
	"inprinteBackoffice/structures"
	"inprinteBackoffice/utils"
	"net/http"

	"github.com/gorilla/mux"
)

func GetRole(w http.ResponseWriter, r *http.Request) {
	//create cors header
	utils.SetCorsHeaders(&w)

	if r.Method == "GET" {

		if utils.Securized(w, r) {
			//global vars
			var id int
			var role string

			//connect the database
			db := utils.DbConnect()

			//get url values
			vars := mux.Vars(r)
			id_role := vars["id_role"]

			//create the sql query
			sqlQuery := ("SELECT id, role FROM role WHERE role.id = " + id_role + " ;")

			//execute the sql query
			row := db.QueryRow(sqlQuery)

			//parse the query
			//retrieve the values and check errors
			err := row.Scan(&id, &role)
			utils.CheckErr(err)

			//add the values to the response
			oneRole := structures.GetRole{
				Id:   id,
				Role: role,
			}

			//close the database connection
			db.Close()

			//create the json response
			json.NewEncoder(w).Encode(oneRole)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
	}
}
