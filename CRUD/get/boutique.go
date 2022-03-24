package CRUD

import (
	"encoding/json"
	databaseTools "inprinte/backend/database"
	"inprinte/backend/structures"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func GetBoutique(w http.ResponseWriter, r *http.Request) {
	db := databaseTools.DbConnect()

	//all sql queries for this path
	newProductsQuery := "SELECT product.id, name, price, description, url FROM product INNER JOIN command_line ON product.id = command_line.id INNER JOIN product_picture ON product.id = product_picture.id_picture INNER JOIN picture ON product_picture.id_picture = picture.id WHERE picture.default = true AND product.pending_validation = false AND product.is_alive = true ORDER BY product.id DESC LIMIT 3;"
	mostWantedProductQuery := "SELECT product.id, COUNT(command_line.id) AS nbrOrder, name, price, description, url FROM product INNER JOIN command_line ON product.id = command_line.id INNER JOIN product_picture ON product.id = product_picture.id_picture INNER JOIN picture ON product_picture.id_picture = picture.id WHERE picture.default = true AND product.is_alive = true AND product.pending_validation = false GROUP BY command_line.id ORDER BY nbrOrder DESC LIMIT 3"
	allProductsQuery := " SELECT product.id, name, price, description, url FROM product INNER JOIN command_line ON product.id = command_line.id INNER JOIN product_picture ON product.id = product_picture.id_picture INNER JOIN picture ON product_picture.id_picture = picture.id WHERE picture.default = true AND product.pending_validation = false AND product.is_alive = true ;"
	categoriesQuery := "SELECT name FROM `category`"

	//----- for the new products -----\\
	newProducts := getNewProducts(db, newProductsQuery)

	//----- for the most wanted products -----\\
	mostWantedProducts := getMostWantedProducts(db, mostWantedProductQuery)

	//----- for all the products -----\\
	allProducts := getAllProducts(db, allProductsQuery)

	categories := getCategories(db, categoriesQuery)

	//----- encode the json response -----\\
	var response = structures.JsonResponseBoutique{
		Type:               "success",
		BoutiqueNews:       newProducts,
		BoutiqueMostWanted: mostWantedProducts,
		AllProducts:        allProducts,
		Categories:         categories,
	}
	json.NewEncoder(w).Encode(response)
}

func GetBoutiqueByCategory(w http.ResponseWriter, r *http.Request) {
	db := databaseTools.DbConnect()

	vars := mux.Vars(r)
	category := vars["category"]

	//all sql queries for this path
	allProductsQuery := "SELECT product.id, product.name, price, description, url FROM product INNER JOIN product_picture ON product.id = product_picture.id_picture INNER JOIN picture ON product_picture.id_picture = picture.id INNER JOIN category ON product.id_category = category.id WHERE picture.default = true AND product.pending_validation = false AND product.is_alive = true AND category.name = \"" + category + "\";"
	log.Println(allProductsQuery)
	//----- for all the products -----\\
	allProducts := getAllProducts(db, allProductsQuery)

	//----- encode the json response -----\\
	var response = structures.JsonResponseBoutique{
		Type:        "success",
		AllProducts: allProducts,
	}
	json.NewEncoder(w).Encode(response)
}
