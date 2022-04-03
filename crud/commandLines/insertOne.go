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
		var response = structures.InsertOneCommandLines{}
		var oneCommandLines structures.CreateOneCommandLines
		var lastInsertID int

		// get body
		err := json.NewDecoder(r.Body).Decode(&oneCommandLines)
		utils.CheckErr(err)

		// connect the database
		db := utils.DbConnect()

		// create the sql query
		sqlQuery := ("INSERT INTO command_line (id_product, id_command, quantity, state) VALUES (" + strconv.Itoa(oneCommandLines.Id_product) + ", " + strconv.Itoa(oneCommandLines.Id_command) + " , " + strconv.Itoa(oneCommandLines.Quantity) + " , '" + oneCommandLines.State + "');")

		// execute the sql query
		_, err = db.Exec(sqlQuery)
		utils.CheckErr(err)

		// get the last inserted id
		sqlQuery = ("SELECT id FROM command_line ORDER BY id DESC LIMIT 1;")
		row := db.QueryRow(sqlQuery)
		err = row.Scan(&lastInsertID)
		utils.CheckErr(err)

		// close the database connection
		db.Close()

		// set the response
		response = structures.InsertOneCommandLines{
			Id:   lastInsertID,
			Type: "success",
			Data: structures.CreateOneCommandLines{
				Id_product: oneCommandLines.Id_product,
				Id_command: oneCommandLines.Id_command,
				Quantity:   oneCommandLines.Quantity,
				State:      oneCommandLines.State,
			},
			Message: "New command line inserted into DB.",
		}

		// send the response
		json.NewEncoder(w).Encode(response)
	}
}
