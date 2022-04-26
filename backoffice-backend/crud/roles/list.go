package crud

import (
	"encoding/json"
	structures "inprinteBackoffice/structures"
	utils "inprinteBackoffice/utils"

	"net/http"
)

func GetRoles(w http.ResponseWriter, r *http.Request) {
	//create cors header
	utils.SetCorsHeaders(&w)

	if r.Method == "GET" {

		if utils.Securized(w, r) {
			//global vars
			var Roles []structures.GetRoles

			//connect the database
			db := utils.DbConnect()

			sqlQuery := "SELECT id, role FROM role"
			rows, err := db.Query(sqlQuery)
			utils.CheckErr(err)

			//parse the query
			for rows.Next() {
				//global vars
				var role string
				var id int

				//retrieve the values and check errors
				err = rows.Scan(&id, &role)
				utils.CheckErr(err)

				//add the values to the response
				Roles = append(Roles, structures.GetRoles{
					Id:   id,
					Role: role,
				})
			}
			//close the rows
			rows.Close()

			//close the database connection
			db.Close()

			//create the json response
			utils.SetXTotalCountHeader(&w, len(Roles))
			json.NewEncoder(w).Encode(Roles)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
	}
}
