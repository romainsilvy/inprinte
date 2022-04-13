package crud

import (
	"encoding/json"
	"inprinteBackoffice/structures"
	"inprinteBackoffice/utils"
	"net/http"
	"strconv"
)

func Update(w http.ResponseWriter, r *http.Request) {
	//create cors header
	utils.SetCorsHeaders(&w)

	if r.Method == "PUT" {

		if utils.Securized(w, r) {
			//global variables
			var category structures.UpdateCategory

			// parse json from put Request
			err := json.NewDecoder(r.Body).Decode(&category)
			utils.CheckErr(err)

			// connect the database
			db := utils.DbConnect()

			// create the sql query
			sqlQuery := ("UPDATE category SET name = '" + category.Name + "', is_alive = " + strconv.FormatBool(category.IsAlive) + " WHERE id = " + strconv.Itoa(category.Id) + ";")

			// execute the sql query
			_, err = db.Exec(sqlQuery)
			utils.CheckErr(err)

			// close the database connection
			db.Close()

			// create the json response
			json.NewEncoder(w).Encode(structures.JsonResponseCategory{
				Id:      category.Id,
				Type:    "success",
				Message: "Category updated",
			})
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
	}
}
