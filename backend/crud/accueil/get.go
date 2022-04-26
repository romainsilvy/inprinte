package crud

import (
	"database/sql"
	"encoding/json"
	structures "inprinte/backend/structures"
	utils "inprinte/backend/utils"
	"strconv"

	"net/http"
)

//Get returns informations for the accueil page
func Get(w http.ResponseWriter, r *http.Request) {
	//set the response header
	utils.SetCorsHeaders(&w)

	//get the database connection
	db := utils.DbConnect()

	//get the informations
	mostSales := getMostSales(db)
	bestRated := getBestRated(db)

	//create the json response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(structures.Accueil{
		MostSales: mostSales,
		BestRated: bestRated,
	})
}

//getMostSales retrieve the three most sold products
func getMostSales(db *sql.DB) []structures.MostSales {
	//global vars
	var mostSales []structures.MostSales
	var product_picture []string

	//execute the sql query and check errors
	rows, err := db.Query("SELECT COUNT(command_line.id_product) AS nbrOrder, command_line.id_product, product.name, product.price FROM command_line INNER JOIN product ON command_line.id_product = product.id WHERE pending_validation = false AND product.is_alive = true GROUP BY command_line.id_product ORDER BY nbrOrder DESC LIMIT 3")
	utils.CheckErr(err)

	//parse the query
	for rows.Next() {
		//global vars
		var name string
		var id, nbrOrder int
		var price float64

		//retrieve the values and check errors
		err = rows.Scan(&nbrOrder, &id, &name, &price)
		utils.CheckErr(err)

		product_picture = getProductPicture(db, strconv.Itoa(id))
		//add the values to the response
		mostSales = append(mostSales, structures.MostSales{
			Id_product: id,
			Name:       name,
			Price:      price,
			Picture:    product_picture[0],
		})
	}
	//close the rows

	//create the json response
	return mostSales
}

//getProductPicture retrieve all the pictures of the given product
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
	//close the rows

	//create the json response
	return product_picture
}

//getBestRated retrieve the three best rated products
func getBestRated(db *sql.DB) []structures.BestRatedProduct {
	//global vars
	var bestRated []structures.BestRatedProduct
	var product_picture []string

	//execute the sql query and check errors
	rows, err := db.Query("SELECT product.id, product.name, product.price, AVG(rate.stars_number) AS rate FROM product INNER JOIN rate ON rate.id_product = product.id WHERE product.is_alive = true AND product.pending_validation = false GROUP BY product.id ORDER BY rate DESC LIMIT 3")
	utils.CheckErr(err)

	//parse the query
	for rows.Next() {
		//global vars
		var name string
		var id int
		var price, rate float64

		//retrieve the values and check errors
		err = rows.Scan(&id, &name, &price, &rate)
		utils.CheckErr(err)

		product_picture = getProductPicture(db, strconv.Itoa(id))
		//add the values to the response
		bestRated = append(bestRated, structures.BestRatedProduct{
			Id_product: id,
			Name:       name,
			Price:      price,
			Picture:    product_picture[0],
		})
	}
	//close the rows

	//create the json response
	return bestRated
}
