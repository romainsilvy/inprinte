package crud

import (
	"database/sql"
	"encoding/json"
	structures "inprinteBackoffice/structures"
	utils "inprinteBackoffice/utils"
	"math"
	"net/http"

	"github.com/gorilla/mux"
)

func GetOne(w http.ResponseWriter, r *http.Request) {
	//global vars
	var oneProduct structures.OneProduct
	var name, description, category, firstname, lastname, role string
	var id, price, id_user int
	var rate float64
	var pending_validation, is_alive bool
	var product_file []structures.FileUrl
	var product_picture []structures.PictureUrl

	vars := mux.Vars(r)
	id_product := vars["id_product"]

	//connect the database
	db := utils.DbConnect()

	//create cors header
	utils.SetCorsHeaders(&w)

	product_file = getProductFile(db, id_product)
	product_picture = getProductPicture(db, id_product)

	sqlQuery := "SELECT product.id, product.name, product.price, product.description, product.pending_validation, product.is_alive, category.name, user.first_name, user.last_name, role.role, user.id AS id_user, AVG(rate.stars_number) AS rate FROM product INNER JOIN category ON category.id = product.id_category INNER JOIN user ON product.id_user = user.id INNER JOIN role ON role.id = user.id_role INNER JOIN rate ON rate.id_product = product.id WHERE product.id = " + id_product + ";"

	row := db.QueryRow(sqlQuery)

	//parse the query
	//retrieve the values and check errors
	err := row.Scan(&id, &name, &price, &description, &pending_validation, &is_alive, &category, &firstname, &lastname, &role, &id_user, &rate)
	utils.CheckErr(err)

	//round the rate
	rate = math.Round(rate*10) / 10

	//add the values to the response
	oneProduct = structures.OneProduct{
		Id:                 id,
		Name:               name,
		Price:              price,
		Description:        description,
		Pending_validation: pending_validation,
		Is_alive:           is_alive,
		Category:           category,
		Firstname:          firstname,
		Lastname:           lastname,
		Role:               role,
		Id_user:            id_user,
		Rate:               rate,
		FileUrl:            product_file,
		PictureUrl:         product_picture,
	}

	//create the json response
	utils.SetXTotalCountHeader(&w, 1)
	json.NewEncoder(w).Encode(oneProduct)
}

func getProductFile(db *sql.DB, id_product string) []structures.FileUrl {
	//global vars
	var product_file []structures.FileUrl
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
		product_file = append(product_file, structures.FileUrl{
			Url: url,
		})
	}

	return product_file
}

func getProductPicture(db *sql.DB, id_product string) []structures.PictureUrl {
	//global vars
	var product_picture []structures.PictureUrl
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
		product_picture = append(product_picture, structures.PictureUrl{
			Url: url,
		})
	}

	return product_picture
}
