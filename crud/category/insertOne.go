package crud

import (
	"encoding/json"
	"fmt"
	structures "inprinteBackoffice/structures"
	utils "inprinteBackoffice/utils"
	"net/http"
)

func InsertOne(w http.ResponseWriter, r *http.Request) {
	utils.SetCorsHeaders(&w)

	// parse json from put Request
	var oneCategory structures.OneCategory
	err := json.NewDecoder(r.Body).Decode(&oneCategory)
	if err != nil {
		fmt.Println(err)
	}

	// connect the database
	db := utils.DbConnect()

	// create the sql query
	sqlQuery := ("INSERT INTO category (name) VALUES ('" + oneCategory.Name + "');")

	// execute the sql query
	_, err = db.Exec(sqlQuery)
	utils.CheckErr(err)
}
