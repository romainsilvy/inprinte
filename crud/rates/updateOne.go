package crud

import (
	"encoding/json"
	"fmt"
	"inprinteBackoffice/structures"
	utils "inprinteBackoffice/utils"
	"net/http"
	"strconv"
)

func UpdateOne(w http.ResponseWriter, r *http.Request) {
	var oneRate structures.OneRate
	err := json.NewDecoder(r.Body).Decode(&oneRate)
	if err != nil {
		fmt.Println(err)
	}
	//connect the database
	db := utils.DbConnect()

	//create cors header
	utils.SetCorsHeaders(&w)

	sqlQuery := ("UPDATE rate SET stars_number = " + strconv.Itoa(oneRate.Stars_number) + " WHERE rate.id = " + strconv.Itoa(oneRate.Id) + ";")

	// execute the sql query
	_, err = db.Exec(sqlQuery)
	utils.CheckErr(err)

	//create the json response
	json.NewEncoder(w).Encode(oneRate)

}
