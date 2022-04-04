package crud

import (
	"encoding/json"
	"fmt"
	structures "inprinteBackoffice/structures"
	utils "inprinteBackoffice/utils"
	"net/http"
	"strconv"
)

func Insert(w http.ResponseWriter, r *http.Request) {
	//create cors header
	utils.SetCorsHeaders(&w)

	// global variables
	var response = structures.JsonResponseRate{}
	var oneRate structures.CreateRate
	var lastInsertID int

	if r.Method == "POST" {
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
		response = structures.JsonResponseRate{
			Id:   lastInsertID,
			Type: "success",
			Data: structures.CreateRate{
				Id_product:   oneRate.Id_product,
				Id_user:      oneRate.Id_user,
				Stars_number: oneRate.Stars_number,
			},
			Message: "New rate inserted into DB.",
		}

		// send the response
		json.NewEncoder(w).Encode(response)
	}
}
