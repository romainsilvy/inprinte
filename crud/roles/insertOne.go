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

	// global variables
	var response = structures.InsertOneRole{}
	var oneRole structures.CreateOneRole
	var lastInsertID int

	// get body
	err := json.NewDecoder(r.Body).Decode(&oneRole)
	if err != nil {
		fmt.Println(err)
	}

	// connect the database
	db := utils.DbConnect()

	// create the sql query
	sqlQuery := ("INSERT INTO role (role) VALUES ('" + oneRole.Role + "');")

	// execute the sql query
	_, err = db.Exec(sqlQuery)
	utils.CheckErr(err)

	// get the last inserted id
	sqlQuery = ("SELECT id FROM role ORDER BY id DESC LIMIT 1;")
	row := db.QueryRow(sqlQuery)
	err = row.Scan(&lastInsertID)
	utils.CheckErr(err)

	// close the database connection
	db.Close()

	// set the response
	response = structures.InsertOneRole{
		Id:   lastInsertID,
		Type: "success",
		Data: structures.CreateOneRole{
			Id:   lastInsertID,
			Role: oneRole.Role,
		},
		Message: "New role inserted into DB.",
	}

	// send the response
	json.NewEncoder(w).Encode(response)
}
