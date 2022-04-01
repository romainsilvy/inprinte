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

	// global variables
	var response = structures.InsertOneRate{}
	var oneRate structures.CreateOneRate
	var lastInsertID int

	// get body
	err := json.NewDecoder(r.Body).Decode(&oneRate)
	if err != nil {
		fmt.Println(err)
	}

	// connect to database
	db := utils.DbConnect()

	// create the sql query
	sqlQuery := ("INSERT INTO rate (id_product, id_user, stars_number) VALUES (" + strconv.Itoa(oneRate.Id_product) + ", " + strconv.Itoa(oneRate.Id_user) + ", " + strconv.Itoa(oneRate.Stars_number) + ");")
	// execute the sql query
	_, err = db.Exec(sqlQuery)
	utils.CheckErr(err)

	// get the last inserted id
	sqlQuery = ("SELECT id FROM rate ORDER BY id DESC LIMIT 1;")
	row := db.QueryRow(sqlQuery)
	err = row.Scan(&lastInsertID)
	utils.CheckErr(err)

	// close the database connection
	db.Close()

	// set the response
	response = structures.InsertOneRate{
		Id:   lastInsertID,
		Type: "success",
		Data: structures.CreateOneRate{
			Id_product:   oneRate.Id_product,
			Id_user:      oneRate.Id_user,
			Stars_number: oneRate.Stars_number,
		},
		Message: "New rate inserted into DB.",
	}

	// send the response
	json.NewEncoder(w).Encode(response)
}
