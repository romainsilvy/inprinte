package crud

import (
	"encoding/json"
	structures "inprinteBackoffice/structures"
	utils "inprinteBackoffice/utils"

	"net/http"
)

func GetCategories(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		//create cors header
		utils.SetCorsHeaders(&w)

		//global vars
		var categories []structures.GetCategories
		var name string
		var id int
		var is_alive bool

		//connect the database
		db := utils.DbConnect()

		//execute the sql query and check errors
		sqlQuery := "SELECT id, name, is_alive FROM category"
		rows, err := db.Query(sqlQuery)
		utils.CheckErr(err)

		//parse the query
		for rows.Next() {
			//retrieve the values and check errors
			err = rows.Scan(&id, &name, &is_alive)
			utils.CheckErr(err)

			//add the values to the response
			categories = append(categories, structures.GetCategories{
				Id:      id,
				Name:    name,
				IsAlive: is_alive,
			})
		}
		//close the rows
		rows.Close()

		//close the database
		db.Close()

		//create the json response
		utils.SetXTotalCountHeader(&w, len(categories))
		json.NewEncoder(w).Encode(categories)
	}
}
