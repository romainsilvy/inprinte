package crud

import (
	"encoding/json"
	structures "inprinteBackoffice/structures"
	utils "inprinteBackoffice/utils"
	"net/http"
	"strconv"
)

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		//create cors header
		utils.SetCorsHeaders(&w)

		// global variables
		var category structures.CreateCategory
		var lastInsertID int

		// get body
		err := json.NewDecoder(r.Body).Decode(&category)
		utils.CheckErr(err)

		// connect to database
		db := utils.DbConnect()

		// create the sql query
		sqlQuery := ("INSERT INTO category SET name = '" + category.Name + "', is_alive = " + strconv.FormatBool(category.IsAlive) + ";")

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
		response := structures.JsonResponseCategory{
			Id:   lastInsertID,
			Type: "success",
			Data: structures.CreateCategory{
				Id:   lastInsertID,
				Name: category.Name,
			},
			Message: "Category created",
		}

		//set the response
		json.NewEncoder(w).Encode(response)
	}
}
