package crud

import (
	"encoding/json"
	"fmt"
	structures "inprinteBackoffice/structures"
	utils "inprinteBackoffice/utils"
	"net/http"
	"strconv"
)

func InsertOne(w http.ResponseWriter, r *http.Request) {
	utils.SetCorsHeaders(&w)

	// parse json from put Request
	var createOneRate structures.CreateOneRate
	err := json.NewDecoder(r.Body).Decode(&createOneRate)
	fmt.Println(createOneRate)

	if err != nil {
		fmt.Println(err)
	}
	// connect the database
	db := utils.DbConnect()

	// create the sql query
	sqlQuery := ("INSERT INTO rate (id_product, id_user, stars_number) VALUES (" + strconv.Itoa(createOneRate.Id_product) + ", " + strconv.Itoa(createOneRate.Id_user) + ", " + strconv.Itoa(createOneRate.Stars_number) + ");")

	// execute the sql query
	_, err = db.Exec(sqlQuery)
	utils.CheckErr(err)
}
