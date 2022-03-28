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
	var oneCommand structures.OneCommand
	err := json.NewDecoder(r.Body).Decode(&oneCommand)
	if err != nil {
		fmt.Println(err)
	}

	// connect the database
	db := utils.DbConnect()

	// create the sql query
	sqlQuery := ("UPDATE command SET state = '" + oneCommand.Status + "' WHERE id = " + strconv.Itoa(oneCommand.Id) + ";")

	// execute the sql query
	_, err = db.Exec(sqlQuery)
	utils.CheckErr(err)

	// create the json response
	json.NewEncoder(w).Encode(oneCommand)
}
