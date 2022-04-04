package crud

import (
	"encoding/json"
	"inprinteBackoffice/structures"
	utils "inprinteBackoffice/utils"
	"net/http"

	"github.com/gorilla/mux"
)

func DeleteOne(w http.ResponseWriter, r *http.Request) {
	//create cors header
	utils.SetCorsHeaders(&w)

	if r.Method == "DELETE" {

		//connect the database
		db := utils.DbConnect()

		//get url values
		vars := mux.Vars(r)
		id_product := vars["id_product"]

		//delete the command lines where there is id_product linked
		sqlQuery := ("DELETE FROM command_line WHERE id_product = " + id_product + ";")
		_, err := db.Exec(sqlQuery)
		utils.CheckErr(err)

		//delete the favorites where there is id_product linked
		sqlQuery = ("DELETE FROM favorite WHERE id_product = " + id_product + ";")
		_, err = db.Exec(sqlQuery)
		utils.CheckErr(err)

		//delete the product_picture where there is id_product linked
		sqlQuery = ("DELETE FROM product_picture WHERE id_product = " + id_product + ";")
		_, err = db.Exec(sqlQuery)
		utils.CheckErr(err)

		//delete the product_file where there is id_product linked
		sqlQuery = ("DELETE FROM product_file WHERE id_product = " + id_product + ";")
		_, err = db.Exec(sqlQuery)
		utils.CheckErr(err)

		//delete the rate where there is id_product linked
		sqlQuery = ("DELETE FROM rate WHERE id_product = " + id_product + ";")
		_, err = db.Exec(sqlQuery)
		utils.CheckErr(err)

		//create the sql query
		sqlQuery = ("DELETE FROM product WHERE id = " + id_product + ";")
		_, err = db.Exec(sqlQuery)
		utils.CheckErr(err)

		//close the database connection
		db.Close()

		//create the json response
		json.NewEncoder(w).Encode(structures.InsertOneProduct{
			Type: "success",
		})
	}
}
