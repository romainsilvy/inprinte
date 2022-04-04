package crud

import (
	"encoding/json"
	structures "inprinteBackoffice/structures"
	utils "inprinteBackoffice/utils"
	"net/http"
)

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// create cors header
		utils.SetCorsHeaders(&w)

		// global variables
		var oneRole structures.CreateRole
		var lastInsertID int

		// get body
		err := json.NewDecoder(r.Body).Decode(&oneRole)
		if err != nil {
			panic(err)
		}

		// connect the database
		db := utils.DbConnect()

		// create the sql query
		sqlQuery := ("INSERT INTO role (role) VALUES ('" + oneRole.Role + "');")
		_, err = db.Exec(sqlQuery)
		utils.CheckErr(err)

		// get the last inserted id
		sqlQuery = ("SELECT id FROM role ORDER BY id DESC LIMIT 1;")

		// execute the query
		row := db.QueryRow(sqlQuery)
		err = row.Scan(&lastInsertID)
		utils.CheckErr(err)

		// close the database connection
		db.Close()

		// set the response
		response := structures.JsonResponseRole{
			Id:   lastInsertID,
			Type: "success",
			Data: structures.CreateRole{
				Id:   lastInsertID,
				Role: oneRole.Role,
			},
			Message: "New role inserted into DB.",
		}

		// send the response
		json.NewEncoder(w).Encode(response)
	}
}
