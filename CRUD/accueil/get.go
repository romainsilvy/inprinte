package crud

import (
	"database/sql"
	"encoding/json"
	structures "inprinte/backend/structures"
	utils "inprinte/backend/utils"

	"net/http"
)

func getMostSales(db *sql.DB) []structures.MostSales {
	//global vars
	var mostSales []structures.MostSales

	//execute the sql query and check errors
	rows, err := db.Query("SELECT COUNT(command_line.id) AS nbrOrder, product.id, name, price, picture.url FROM product INNER JOIN command_line ON product.id = command_line.id_product INNER JOIN product_picture ON product_picture.id = product.id INNER JOIN picture ON picture.id = product_picture.id WHERE pending_validation = false AND product.is_alive = true GROUP BY command_line.id_product ORDER BY nbrOrder DESC LIMIT 3")
	utils.CheckErr(err)

	//parse the query
	for rows.Next() {
		//global vars
		var name, picture string
		var id, nbrOrder int
		var price float64

		//retrieve the values and check errors
		err = rows.Scan(&nbrOrder, &id, &name, &price, &picture)
		utils.CheckErr(err)

		//add the values to the response
		mostSales = append(mostSales, structures.MostSales{
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

func getBestRated(db *sql.DB) []structures.BestRatedProduct {
	//global vars
	var bestRated []structures.BestRatedProduct

	//execute the sql query and check errors
	rows, err := db.Query("SELECT product.id, product.name, product.price, picture.url, AVG(rate.stars_number) AS rate FROM product INNER JOIN product_picture ON product_picture.id = product.id INNER JOIN picture ON picture.id = product_picture.id INNER JOIN rate ON rate.id_product = product.id WHERE product.is_alive = true AND product.pending_validation = false GROUP BY product.id ORDER BY rate DESC LIMIT 3")
	utils.CheckErr(err)

	//parse the query
	for rows.Next() {
		//global vars
		var name, picture string
		var id int
		var price, rate float64

		//retrieve the values and check errors
		err = rows.Scan(&id, &name, &price, &picture, &rate)
		utils.CheckErr(err)

		//add the values to the response
		bestRated = append(bestRated, structures.BestRatedProduct{
			Id_product: id,
			Name:       name,
			Price:      price,
			Picture:    picture,
		})
	}
	//close the rows

	//create the json response
	return bestRated
}

func GetAccueil(w http.ResponseWriter, r *http.Request) {
	db := utils.DbConnect()
	mostSales := getMostSales(db)
	bestRated := getBestRated(db)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(structures.Accueil{
		MostSales: mostSales,
		BestRated: bestRated,
	})
}
