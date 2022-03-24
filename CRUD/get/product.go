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
	id := vars["id_product"]
	//Get products infos
	rows, err := db.Query(`SELECT product.id, name, description, price, AVG(stars_number) AS MOYENNE, picture.url, product_file.id FROM product INNER JOIN rate ON rate.id = product.id INNER JOIN product_picture ON product_picture.id = product.id INNER JOIN picture ON picture.id = product_picture.id INNER JOIN product_file ON product_file.id = product.id WHERE product.id = ? AND pending_validation = false AND product.is_alive = true GROUP BY product.id;`, id)

	// check errors
	utils.CheckErr(err)

	// var response []JsonResponse
	var products []structures.Product

	// Foreach product
	for rows.Next() {
		var id_product, price int
		var stars_number float64
		var name, description, picture_url, product_file string

		//var picture string
		err = rows.Scan(&id_product, &name, &description, &price, &stars_number, &picture_url, &product_file)

		// check errors
		utils.CheckErr(err)

		products = append(products, structures.Product{
			Id_product:   id_product,
			Name:         name,
			Description:  description,
			Price:        price,
			Stars_number: stars_number,
			Picture_url:  picture_url,
			Product_file: product_file,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func getNewProducts(db *sql.DB, sqlQuery string) []structures.BoutiqueProduct {
	var id_product int
	var name, price, description, picture string
	var newProducts []structures.BoutiqueProduct

	rows, err := db.Query(sqlQuery)
	utils.CheckErr(err)

	for rows.Next() {
		err = rows.Scan(&id_product, &name, &price, &description, &picture)

		// check errors
		utils.CheckErr(err)

		newProducts = append(newProducts, structures.BoutiqueProduct{
			Id_product:  id_product,
			Name:        name,
			Price:       price,
			Description: description,
			Picture:     picture,
		})
	}
	return newProducts
}

func getMostWantedProducts(db *sql.DB, sqlQuery string) []structures.BoutiqueProduct {
	var name, price, description, picture string
	var nbrOrder, id_product int
	var mostWantedProducts []structures.BoutiqueProduct

	rows, err := db.Query(sqlQuery)
	utils.CheckErr(err)

	for rows.Next() {
		err = rows.Scan(&id_product, &nbrOrder, &name, &price, &description, &picture)

		// check errors
		utils.CheckErr(err)

		mostWantedProducts = append(mostWantedProducts, structures.BoutiqueProduct{
			Id_product:  id_product,
			Name:        name,
			Price:       price,
			Description: description,
			Picture:     picture,
		})
	}
	return mostWantedProducts
}

func getAllProducts(db *sql.DB, sqlQuery string) []structures.BoutiqueProduct {
	var id_product int
	var name, price, description, picture string
	var allProducts []structures.BoutiqueProduct

	rows, err := db.Query(sqlQuery)
	utils.CheckErr(err)

	for rows.Next() {
		err = rows.Scan(&id_product, &name, &price, &description, &picture)

		// check errors
		utils.CheckErr(err)

		allProducts = append(allProducts, structures.BoutiqueProduct{
			Id_product:  id_product,
			Name:        name,
			Price:       price,
			Description: description,
			Picture:     picture,
		})
	}
	return allProducts
}
