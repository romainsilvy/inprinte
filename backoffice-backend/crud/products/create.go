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
			var oneProduct = structures.CreateProduct{}
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

			// create the sql query
			for i := 0; i < len(oneProduct.FileUrl); i++ {
				sqlQuery = ("INSERT INTO product_file (id_product, url) VALUES (" + strconv.Itoa(lastInsertID) + ", '" + oneProduct.FileUrl[i] + "');")

				// execute the sql query
				_, err = db.Exec(sqlQuery)
				utils.CheckErr(err)
			}

			// create the sql query for picture table
			for i := 0; i < len(oneProduct.PictureUrl); i++ {
				sqlQuery = ("INSERT INTO picture (url) VALUES ('" + oneProduct.PictureUrl[i] + "');")

				// execute the sql query
				_, err = db.Exec(sqlQuery)
				utils.CheckErr(err)

				// create the sql query
				sqlQuery = ("INSERT INTO product_picture (id_picture, id_product) VALUES ((SELECT id FROM picture ORDER BY id DESC LIMIT 1), '" + strconv.Itoa(lastInsertID) + "');")

				// execute the sql query
				_, err = db.Exec(sqlQuery)
				utils.CheckErr(err)
			}

			// close the database connection
			db.Close()

			// set the response
			response := structures.JsonResponseProduct{
				Id:   lastInsertID,
				Type: "success",
				Data: structures.CreateProduct{
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

			// set the response
			json.NewEncoder(w).Encode(response)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
	}
}
