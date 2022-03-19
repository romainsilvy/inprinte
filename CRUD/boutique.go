package CRUD

import (
	"database/sql"
	databaseTools "inprinte/backend/database"
	structures "inprinte/backend/structures"
	utils "inprinte/backend/utils"

	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// func getGlobalCategories(db *sql.DB) []structures.BoutiqueCategories {
// 	rows, err := db.Query("SELECT global_category_name, category_name FROM category INNER JOIN global_category ON category.id_global_category = global_category.id_global_category")

// 	// check errors
// 	utils.CheckErr(err)

// 	// var response []JsonResponse
// 	var categories []structures.BoutiqueCategories

// 	// Foreach product
// 	for rows.Next() {
// 		var globalCategory string
// 		var price string
// 		var description string
// 		//var picture string
// 		err = rows.Scan(&globalCategory, &price, &description)

// 		// check errors
// 		utils.CheckErr(err)

// 		categories = append(categories, structures.BoutiqueCategories{
// 			GlobalCategory: globalCategory,
// 		})
// 	}
// 	return categories
// }

func getNewProducts(db *sql.DB, sqlQuery string) []structures.BoutiqueProduct {
	var name, price, description string
	var newProducts []structures.BoutiqueProduct

	rows := databaseTools.GetRequest(db, sqlQuery)
	for rows.Next() {
		err := rows.Scan(&name, &price, &description)

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

	rows := databaseTools.GetRequest(db, sqlQuery)
	for rows.Next() {
		err := rows.Scan(&nbrOrder, &name, &price, &description)

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

	rows := databaseTools.GetRequest(db, sqlQuery)
	for rows.Next() {
		err := rows.Scan(&name, &price, &description)

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
