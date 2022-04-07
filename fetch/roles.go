package crud

import (
	"encoding/json"
	structures "inprinteBackoffice/structures"
	utils "inprinteBackoffice/utils"

	"net/http"
)

func FetchRoles(w http.ResponseWriter, r *http.Request) {
	//create cors header
	utils.SetCorsHeaders(&w)

	if r.Method == "GET" {
		//global vars
		var list structures.RoleList
		//connect the database
		db := utils.DbConnect()

		sqlQuery := ("SELECT role FROM role ;")

		//execute the sql query and check errors
		rows, err := db.Query(sqlQuery)
		utils.CheckErr(err)

		//global vars
		var roleOne string
		var roleList []string
		//parse the query
		for rows.Next() {

			//retrieve the values and check errors
			err = rows.Scan(&roleOne)
			utils.CheckErr(err)

			//add the values to the response
			roleList = append(roleList, roleOne)
		}

		//add the values to the response
		list = structures.RoleList{
			RoleList: roleList,
		}

		//close the rows
		rows.Close()

		//close the database connection
		db.Close()

		//create the json response
		json.NewEncoder(w).Encode(list)
	}
}
