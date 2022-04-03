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
		var response = structures.InsertOneProduct{}
		var oneProduct = structures.CreateOneProduct{}
		var lastInsertID int

		// get body
		err := json.NewDecoder(r.Body).Decode(&oneProduct)
		utils.CheckErr(err)

		// connect the database
		db := utils.DbConnect()

		// create the sql query
		sqlQuery := ("INSERT INTO product (name, price, description, pending_validation, is_alive, id_category, id_user) VALUES ('" + oneProduct.Name + "', " + strconv.Itoa(oneProduct.Price) + " , '" + oneProduct.Description + "' , " + strconv.FormatBool(oneProduct.Pending_validation) + " , " + strconv.FormatBool(oneProduct.Is_alive) + " , (SELECT id FROM category WHERE category.name = '" + oneProduct.Category + "'), " + strconv.Itoa(oneProduct.Id_user) + ");")

		// execute the sql query
		_, err = db.Exec(sqlQuery)
		utils.CheckErr(err)

		// get the last inserted id
		sqlQuery = ("SELECT id FROM product ORDER BY id DESC LIMIT 1;")
		row := db.QueryRow(sqlQuery)
		err = row.Scan(&lastInsertID)
		utils.CheckErr(err)

		// close the database connection
		db.Close()

		// set the response
		response = structures.InsertOneProduct{
			Id:   lastInsertID,
			Type: "success",
			Data: structures.CreateOneProduct{
				Name:               oneProduct.Name,
				Price:              oneProduct.Price,
				Description:        oneProduct.Description,
				Pending_validation: oneProduct.Pending_validation,
				Is_alive:           oneProduct.Is_alive,
				Category:           oneProduct.Category,
				Id_user:            oneProduct.Id_user,
			},
			Message: "New product inserted into DB.",
		}

		// send the response
		json.NewEncoder(w).Encode(response)
	}
}
