package crud

import (
	"encoding/json"
	structures "inprinteBackoffice/structures"
	utils "inprinteBackoffice/utils"
	"net/http"
)

func InsertOne(w http.ResponseWriter, r *http.Request) {
	utils.SetCorsHeaders(&w)

	// global variables
	var response = structures.InsertOne{}
	var oneCategory structures.OneCategory
	var lastInsertID int

	// get body
	err := json.NewDecoder(r.Body).Decode(&oneCategory)
	utils.CheckErr(err)

	// connect to database
	db := utils.DbConnect()

	// create the sql query
	sqlQuery := ("INSERT INTO category SET name = '" + oneCategory.Name + "';")

	// execute the sql query
	_, err = db.Exec(sqlQuery)
	utils.CheckErr(err)

	// get the last inserted id
	sqlQuery = ("SELECT id FROM category ORDER BY id DESC LIMIT 1;")
	row := db.QueryRow(sqlQuery)
	err = row.Scan(&lastInsertID)
	utils.CheckErr(err)

	// close the database connection
	db.Close()

	// set the response
	response = structures.InsertOne{
		Id:   lastInsertID,
		Type: "success",
		Data: structures.OneCategory{
			Id:   lastInsertID,
			Name: oneCategory.Name,
		},
		Message: "New categorie inserted into DB.",
	}

	// send the response
	json.NewEncoder(w).Encode(response)
}
