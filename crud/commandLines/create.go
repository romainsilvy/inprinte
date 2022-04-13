package crud

import (
	"encoding/json"
	structures "inprinteBackoffice/structures"
	utils "inprinteBackoffice/utils"
	"net/http"
	"strconv"
)

func Insert(w http.ResponseWriter, r *http.Request) {
	//create cors header
	utils.SetCorsHeaders(&w)

	if r.Method == "POST" {

		if utils.Securized(w, r) {
			// global variables
			var response = structures.JsonReponseCommandLines{}
			var oneCommandLines structures.CreateCommandLine
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

			//create the query
			sqlQuery = ("SELECT id FROM command_line ORDER BY id DESC LIMIT 1;")

			//execute the query
			row := db.QueryRow(sqlQuery)
			err = row.Scan(&lastInsertID)
			utils.CheckErr(err)

			// close the database connection
			db.Close()

			// set the response
			response = structures.JsonReponseCommandLines{
				Id:   lastInsertID,
				Type: "success",
				Data: structures.CreateCommandLine{
					Id_product: oneCommandLines.Id_product,
					Id_command: oneCommandLines.Id_command,
					Quantity:   oneCommandLines.Quantity,
					State:      oneCommandLines.State,
				},
				Message: "New command line inserted into DB.",
			}

			//set the response
			json.NewEncoder(w).Encode(response)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
	}
}
