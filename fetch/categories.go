package crud

import (
	"encoding/json"
	structures "inprinteBackoffice/structures"
	utils "inprinteBackoffice/utils"

	"net/http"
)

func FetchCategories(w http.ResponseWriter, r *http.Request) {
	//create cors header
	utils.SetCorsHeaders(&w)

	if r.Method == "GET" {
		//global vars
		var list structures.CategoriesList
		//connect the database
		db := utils.DbConnect()

		sqlQuery := ("SELECT name FROM category ;")

		//execute the sql query and check errors
		rows, err := db.Query(sqlQuery)
		utils.CheckErr(err)

		//global vars
		var categoryOne string
		var categoryList []string
		//parse the query
		for rows.Next() {

			//retrieve the values and check errors
			err = rows.Scan(&categoryOne)
			utils.CheckErr(err)

			//add the values to the response
			categoryList = append(categoryList, categoryOne)
		}

		//add the values to the response
		list = structures.CategoriesList{
			CategoriesList: categoryList,
		}

		//close the rows
		rows.Close()

		//close the database connection
		db.Close()

		//create the json response
		json.NewEncoder(w).Encode(list)
	}
}
