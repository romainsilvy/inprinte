package crud

import (
	"encoding/json"
	"fmt"
	"inprinteBackoffice/structures"
	"inprinteBackoffice/utils"
	"net/http"
	"strconv"
)

func UpdateOne(w http.ResponseWriter, r *http.Request) {
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
	sqlQuery := ("UPDATE category SET name = '" + oneCategory.Name + "' WHERE id = " + strconv.Itoa(oneCategory.Id) + ";")
	// execute the sql query
	_, err = db.Exec(sqlQuery)
	utils.CheckErr(err)

	// create the json response
	json.NewEncoder(w).Encode(oneCategory)
}
