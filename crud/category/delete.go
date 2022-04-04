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
		//connect the database
		db := utils.DbConnect()

		//get url values
		vars := mux.Vars(r)
		id_category := vars["id_category"]

		//create the sql query
		sqlQuery := ("UPDATE category SET is_alive = false WHERE id = " + id_category + ";")
		_, err := db.Exec(sqlQuery)
		utils.CheckErr(err)

		//update products
		sqlQuery = ("UPDATE product SET is_alive = false WHERE id_category = " + id_category + ";")
		_, err = db.Exec(sqlQuery)
		utils.CheckErr(err)

		//close the database connection
		db.Close()

		//create the json response
		json.NewEncoder(w).Encode(structures.JsonResponseCategory{
			Type:    "success",
			Message: "Category killed",
		})
	}
}
