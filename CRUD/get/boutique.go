package CRUD

import (
	"encoding/json"
	databaseTools "inprinte/backend/database"
	"inprinte/backend/structures"
	"net/http"
)

func GetBoutique(w http.ResponseWriter, r *http.Request) {
	db := databaseTools.DbConnect()

	//all sql queries for this path
	newProductsQuery := "SELECT name, price, description FROM product ORDER BY id_product DESC LIMIT 3;"
	mostWantedProductQuery := "SELECT COUNT(command_line.id_product) AS nbrOrder, name, price, description FROM product INNER JOIN command_line ON product.id_product = command_line.id_product GROUP BY command_line.id_product ORDER BY nbrOrder DESC LIMIT 3"
	allProductsQuery := " SELECT name, price, description FROM product;"

	//----- for the new products -----\\
	newProducts := getNewProducts(db, newProductsQuery)

	//----- for the most wanted products -----\\
	mostWantedProducts := getMostWantedProducts(db, mostWantedProductQuery)

	//----- for all the products -----\\
	allProducts := getAllProducts(db, allProductsQuery)

	// categories := getAllCategories(db)

	//----- encode the json response -----\\
	var response = structures.JsonResponseBoutique{
		Type:               "success",
		BoutiqueNews:       newProducts,
		BoutiqueMostWanted: mostWantedProducts,
		AllProducts:        allProducts,
		// Categories:         categories,
	}
	json.NewEncoder(w).Encode(response)
}