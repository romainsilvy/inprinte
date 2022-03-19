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

func getNewProducts(db *sql.DB) []structures.BoutiqueProduct {
	rows, err := db.Query("SELECT name, price, description FROM product ORDER BY id_product DESC LIMIT 3")

	// check errors
	utils.CheckErr(err)

	// var response []JsonResponse
	var newProducts []structures.BoutiqueProduct

	// Foreach product
	for rows.Next() {
		var name string
		var price string
		var description string
		//var picture string
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

func getMostWantedProduct(db *sql.DB) []structures.BoutiqueProduct {
	rows, err := db.Query("SELECT COUNT(command_line.id_product) AS nbrOrder, name, price, description FROM product INNER JOIN command_line ON product.id_product = command_line.id_product GROUP BY command_line.id_product ORDER BY nbrOrder DESC LIMIT 3")

	// check errors
	utils.CheckErr(err)

	// var response []JsonResponse
	var mostWantedProduct []structures.BoutiqueProduct

	// Foreach product
	for rows.Next() {
		var nbrOrder int
		var name string
		var price string
		var description string
		//var picture string
		err = rows.Scan(&nbrOrder, &name, &price, &description)

		// check errors
		utils.CheckErr(err)

		mostWantedProduct = append(mostWantedProduct, structures.BoutiqueProduct{
			Name:        name,
			Price:       price,
			Description: description,
		})
	}
	return mostWantedProduct
}

func getAllProducts(db *sql.DB) []structures.BoutiqueProduct {
	rows, err := db.Query("SELECT name, price, description FROM product")

	// check errors
	utils.CheckErr(err)

	// var response []JsonResponse
	var allProducts []structures.BoutiqueProduct

	// Foreach product
	for rows.Next() {
		var name string
		var price string
		var description string
		//var picture string
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

func GetBoutique(w http.ResponseWriter, r *http.Request) {
	db := databaseTools.DbConnect()

	//get boutique infos\\

	//get new products
	newProducts := getNewProducts(db)
	mostWantedProduct := getMostWantedProduct(db)
	allProducts := getAllProducts(db)
	// categories := getAllCategories(db)

	var response = structures.JsonResponseBoutique{
		Type:               "success",
		BoutiqueNews:       newProducts,
		BoutiqueMostWanted: mostWantedProduct,
		AllProducts:        allProducts,
		// Categories:         categories,
	}

	json.NewEncoder(w).Encode(response)
}
