package crud

import (
	"encoding/json"
	structures "inprinteBackoffice/structures"
	utils "inprinteBackoffice/utils"

	"net/http"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	//global vars
	var Category []structures.Category

	//connect the database
	db := utils.DbConnect()

	//create cors header
	utils.SetCorsHeaders(&w)

	sqlQuery := "SELECT id, name FROM category"
	//execute the sql query and check errors
	rows, err := db.Query(sqlQuery)
	utils.CheckErr(err)

	//parse the query
	for rows.Next() {
		//global vars
		var name string
		var id int

		//retrieve the values and check errors
		err = rows.Scan(&id, &name)
		utils.CheckErr(err)

		//add the values to the response
		Category = append(Category, structures.Category{
			Id:   id,
			Name: name,
		})
	}
	//close the rows

	//create the json response
	utils.SetXTotalCountHeader(&w, len(Category))
	json.NewEncoder(w).Encode(Category)
}
