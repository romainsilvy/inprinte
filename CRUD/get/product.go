package CRUD

import (
	"database/sql"
	"encoding/json"
	"fmt"
	databaseTools "inprinte/backend/database"
	"inprinte/backend/structures"
	"inprinte/backend/utils"
	"net/http"

	"github.com/gorilla/mux"
)

func GetOneProduct(w http.ResponseWriter, r *http.Request) {
	db := databaseTools.DbConnect()
	fmt.Println("Getting products ...")
	vars := mux.Vars(r)
	id := vars["id"]
	//Get products infos
	rows, err := db.Query("SELECT product.id, name, description, price, AVG(stars_number) AS MOYENNE, picture.url, product_file.id FROM product INNER JOIN rate ON rate.id = product.id INNER JOIN product_picture ON product_picture.id = product.id INNER JOIN picture ON picture.id = product_picture.id INNER JOIN product_file ON product_file.id = product.id WHERE product.id = " + id + " GROUP BY product.id;")

	// check errors
	utils.CheckErr(err)

	// var response []JsonResponse
	var products []structures.Product

	// Foreach product
	for rows.Next() {
		var id_product int
		var name string
		var description string
		var price int
		var stars_number float64
		var picture_url string
		var product_file string

		//var picture string
		err = rows.Scan(&id_product, &name, &description, &price, &stars_number, &picture_url, &product_file)

		// check errors
		utils.CheckErr(err)

		products = append(products, structures.Product{
			Id_product:		id_product,
			Name:         name,
			Description:  description,
			Price:        price,
			Stars_number: stars_number,
			Picture_url:  picture_url,
			Product_file: product_file,
		})
	}

	var response = structures.JsonResponseProduct{
		Type: "success",
		Data: products,
	}

	if response.Data == nil {
		w.WriteHeader(404)
	} else {
		json.NewEncoder(w).Encode(response)
	}
}

func getNewProducts(db *sql.DB, sqlQuery string) []structures.BoutiqueProduct {
	var name, price, description string
	var newProducts []structures.BoutiqueProduct

	rows, err := db.Query(sqlQuery)
	utils.CheckErr(err)

	for rows.Next() {
		err = rows.Scan(&name, &price, &description)

		// check errors
		utils.CheckErr(err)

		newProducts = append(newProducts, structures.BoutiqueProduct{
			Name:        name,
			Price:       price,
			Description: description,
		})
	}
	return newProducts
}

func getMostWantedProducts(db *sql.DB, sqlQuery string) []structures.BoutiqueProduct {
	var name, price, description string
	var nbrOrder int
	var mostWantedProducts []structures.BoutiqueProduct

	rows, err := db.Query(sqlQuery)
	utils.CheckErr(err)

	for rows.Next() {
		err = rows.Scan(&nbrOrder, &name, &price, &description)

		// check errors
		utils.CheckErr(err)

		mostWantedProducts = append(mostWantedProducts, structures.BoutiqueProduct{
			Name:        name,
			Price:       price,
			Description: description,
		})
	}
	return mostWantedProducts
}

func getAllProducts(db *sql.DB, sqlQuery string) []structures.BoutiqueProduct {
	var name, price, description string
	var allProducts []structures.BoutiqueProduct

	rows, err := db.Query(sqlQuery)
	utils.CheckErr(err)

	for rows.Next() {
		err = rows.Scan(&name, &price, &description)

		// check errors
		utils.CheckErr(err)

		allProducts = append(allProducts, structures.BoutiqueProduct{
			Name:        name,
			Price:       price,
			Description: description,
		})
	}
	return allProducts
}