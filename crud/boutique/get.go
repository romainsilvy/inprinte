package crud

import (
	"database/sql"
	"encoding/json"
	structures "inprinte/backend/structures"
	utils "inprinte/backend/utils"

	"net/http"
)

func getNewProducts(db *sql.DB) []structures.NewProduct {
	//global vars
	var newProducts []structures.NewProduct

	//execute the sql query and check errors
	rows, err := db.Query("SELECT product.id, name, price, url FROM product INNER JOIN product_picture ON product.id = product_picture.id_picture INNER JOIN picture ON product_picture.id_picture = picture.id WHERE product.pending_validation = false AND product.is_alive = true ORDER BY product.id DESC LIMIT 3")
	utils.CheckErr(err)

	//parse the query
	for rows.Next() {
		//global vars
		var name, picture string
		var price float64
		var id int

		//retrieve the values and check errors
		err = rows.Scan(&id, &name, &price, &picture)
		utils.CheckErr(err)

		//add the values to the response
		newProducts = append(newProducts, structures.NewProduct{
			Id_product: id,
			Name:       name,
			Price:      price,
			Picture:    picture,
		})
	}
	//close the rows

	//create the json response
	return newProducts
}

func getMostSales(db *sql.DB) []structures.MostWantedProduct {
	//global vars
	var mostSales []structures.MostWantedProduct

	//execute the sql query and check errors
	rows, err := db.Query("SELECT product.id, COUNT(command_line.id) AS nbrOrder, name, price, url FROM product INNER JOIN command_line ON product.id = command_line.id INNER JOIN product_picture ON product.id = product_picture.id_picture INNER JOIN picture ON product_picture.id_picture = picture.id WHERE product.is_alive = true AND product.pending_validation = false GROUP BY command_line.id ORDER BY nbrOrder DESC LIMIT 3")
	utils.CheckErr(err)

	//parse the query
	for rows.Next() {
		//global vars
		var name, picture string
		var price float64
		var id, nbrOrder int

		//retrieve the values and check errors
		err = rows.Scan(&id, &nbrOrder, &name, &price, &picture)
		utils.CheckErr(err)

		//add the values to the response
		mostSales = append(mostSales, structures.MostWantedProduct{
			Id_product: id,
			Name:       name,
			Price:      price,
			Picture:    picture,
		})
	}
	//close the rows

	//create the json response
	return mostSales
}

func getAllProducts(r *http.Request, db *sql.DB) []structures.BoutiqueProduct {
	//global vars
	var allProducts []structures.BoutiqueProduct

	//retrieve the url params
	orderBy, rangeBy := utils.GetAllParams(r)

	//execute the sql query and check errors
	rows, err := db.Query("SELECT product.id, name, price, description, url FROM product INNER JOIN product_picture ON product.id = product_picture.id_picture INNER JOIN picture ON product_picture.id_picture = picture.id WHERE product.pending_validation = false AND product.is_alive = true " + orderBy + " " + rangeBy)
	utils.CheckErr(err)

	//parse the query
	for rows.Next() {
		//global vars
		var name, description, picture string
		var price float64
		var id int

		//retrieve the values and check errors
		err = rows.Scan(&id, &name, &price, &description, &picture)
		utils.CheckErr(err)

		//add the values to the response
		allProducts = append(allProducts, structures.BoutiqueProduct{
			Id_product:  id,
			Name:        name,
			Price:       price,
			Picture:     picture,
			Description: description,
		})
	}
	//close the rows

	//create the json response
	return allProducts
}

func getCategories(db *sql.DB) []string {
	//global vars
	var allCategories []string

	//execute the sql query and check errors
	rows, err := db.Query("SELECT name FROM `category`")
	utils.CheckErr(err)

	//parse the query
	for rows.Next() {
		//global vars
		var name string

		//retrieve the values and check errors
		err = rows.Scan(&name)
		utils.CheckErr(err)

		//add the values to the response
		allCategories = append(allCategories, name)
	}
	//close the rows

	//create the json response
	return allCategories
}

func Get(w http.ResponseWriter, r *http.Request) {
	utils.SetCorsHeaders(&w)
	//connect to the database
	db := utils.DbConnect()

	//get informations
	newProducts := getNewProducts(db)
	mostSales := getMostSales(db)
	allProducts := getAllProducts(r, db)
	categories := getCategories(db)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(structures.Boutique{
		NewProducts: newProducts,
		MostSales:   mostSales,
		AllProducts: allProducts,
		Categories:  categories,
	})
}
