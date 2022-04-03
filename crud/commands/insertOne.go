package crud

import (
	"encoding/json"
	structures "inprinteBackoffice/structures"
	utils "inprinteBackoffice/utils"
	"net/http"
	"strconv"
)

func InsertOne(w http.ResponseWriter, r *http.Request) {
	utils.SetCorsHeaders(&w)
	if r.Method == "POST" {
		// global variables
		var response = structures.InsertOneCommand{}
		var oneCommand structures.CreateOneCommand
		var lastInsertID int

		// get body
		err := json.NewDecoder(r.Body).Decode(&oneCommand)
		utils.CheckErr(err)

		// connect the database
		db := utils.DbConnect()

		// create the sql query
		sqlQuery := ("INSERT INTO command (id_user, date, state) VALUES (" + strconv.Itoa(oneCommand.Id_user) + ", '" + oneCommand.Date + "' , '" + oneCommand.State + "');")

		// execute the sql query
		_, err = db.Exec(sqlQuery)
		utils.CheckErr(err)

		// get the last inserted id
		sqlQuery = ("SELECT id FROM command ORDER BY id DESC LIMIT 1;")
		row := db.QueryRow(sqlQuery)
		err = row.Scan(&lastInsertID)
		utils.CheckErr(err)

		// close the database connection
		db.Close()

		// set the response
		response = structures.InsertOneCommand{
			Id:   lastInsertID,
			Type: "success",
			Data: structures.CreateOneCommand{
				Id_user: oneCommand.Id_user,
				Date:    oneCommand.Date,
				State:   oneCommand.State,
			},
			Message: "New command inserted into DB.",
		}

		// send the response
		json.NewEncoder(w).Encode(response)
	}
}
