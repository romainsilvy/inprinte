package crud

import (
	"database/sql"
	"encoding/json"
	structures "inprinteBackoffice/structures"
	utils "inprinteBackoffice/utils"
	"net/http"

	"github.com/gorilla/mux"
)

func GetProduct(w http.ResponseWriter, r *http.Request) {
	//create cors header
	utils.SetCorsHeaders(&w)

	if r.Method == "GET" {
		//global vars
		var oneProduct structures.GetProduct
		var name, description, category string
		var id, price int
		var pending_validation, is_alive bool
		var product_file []string
		var product_picture []string

		vars := mux.Vars(r)
		id_product := vars["id_product"]

		//connect the database
		db := utils.DbConnect()

		//create the sql query
		product_file = getProductFile(db, id_product)
		product_picture = getProductPicture(db, id_product)
		sqlQuery := ("SELECT product.id, product.name, product.price, product.description, product.pending_validation, product.is_alive, category.name FROM product INNER JOIN category ON category.id = product.id_category WHERE product.id = " + id_product + ";")

		//execute the sql query
		row := db.QueryRow(sqlQuery)

		//parse the query
		//retrieve the values and check errors
		err := row.Scan(&id, &name, &price, &description, &pending_validation, &is_alive, &category)
		utils.CheckErr(err)

		oneProduct = structures.GetProduct{
			Id:                 id,
			Name:               name,
			Price:              price,
			Description:        description,
			Pending_validation: pending_validation,
			Is_alive:           is_alive,
			Category:           category,
			FileUrl:            product_file,
			PictureUrl:         product_picture,
		}

		//close the database connection
		db.Close()

		//create the json response
		utils.SetXTotalCountHeader(&w, 1)
		json.NewEncoder(w).Encode(oneProduct)
	}
}

func getProductFile(db *sql.DB, id_product string) []string {
	//global vars
	var product_file []string
	var url string
	//create the sql query
	sqlQuery := ("SELECT product_file.url FROM product_file INNER JOIN product ON product_file.id_product = product.id WHERE product.id = " + id_product + ";")

	rows, err := db.Query(sqlQuery)
	utils.CheckErr(err)

	for rows.Next() {
		//retrieve the values and check errors
		err = rows.Scan(&url)
		utils.CheckErr(err)

		//add the values to the response
		product_file = append(product_file, url)
	}

	return product_file
}

func getProductPicture(db *sql.DB, id_product string) []string {
	//global vars
	var product_picture []string
	var url string
	//create the sql query
	sqlQuery := ("SELECT picture.url FROM picture INNER JOIN product_picture ON product_picture.id_picture = picture.id INNER JOIN product ON product.id = product_picture.id_product WHERE product.id = " + id_product + ";")

	rows, err := db.Query(sqlQuery)
	utils.CheckErr(err)

	for rows.Next() {
		//retrieve the values and check errors
		err = rows.Scan(&url)
		utils.CheckErr(err)

		//add the values to the response
		product_picture = append(product_picture, url)
	}

	return product_picture
}
