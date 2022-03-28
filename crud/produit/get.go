package crud

import (
	"database/sql"
	"encoding/json"
	structures "inprinte/backend/structures"
	utils "inprinte/backend/utils"

	"net/http"

	"github.com/gorilla/mux"
)

func getOneProduct(w http.ResponseWriter, db *sql.DB, id_product string) structures.ProductData {
	//global vars
	var id int
	var name, description, picture, product_file string
	var price float64

	//execute the sql query and check errors
	row := db.QueryRow(`SELECT product.id, name, description, price, picture.url, product_file.url FROM product INNER JOIN rate ON rate.id = product.id INNER JOIN product_picture ON product_picture.id = product.id INNER JOIN picture ON picture.id = product_picture.id INNER JOIN product_file ON product_file.id = product.id WHERE product.id = ? AND pending_validation = false AND product.is_alive = true GROUP BY product.id`, id_product)
	err := row.Scan(&id, &name, &description, &price, &picture, &product_file)
	if err == sql.ErrNoRows {
		w.WriteHeader(404)
	} else {
		utils.CheckErr(err)
	}

	//create the json response
	return structures.ProductData{
		Id_product:   id,
		Name:         name,
		Description:  description,
		Price:        price,
		Picture:      picture,
		Product_file: product_file,
	}
}

func getRelatedProduct(db *sql.DB, id_product string) []structures.ProductRelated {
	//global vars
	var productRelated []structures.ProductRelated

	//execute the sql query and check errors
	rows, err := db.Query("SELECT product.id, product.name, product.price, picture.url, AVG(rate.stars_number) AS rate FROM product INNER JOIN product_picture ON product_picture.id = product.id INNER JOIN picture ON picture.id = product_picture.id INNER JOIN rate ON rate.id_product = product.id WHERE product.is_alive = true AND product.pending_validation = false AND product.id_category = (SELECT category.id FROM category INNER JOIN product ON product.id_category = category.id WHERE product.id = ?) GROUP BY product.id ORDER BY rate DESC LIMIT 4", id_product)
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
		productRelated = append(productRelated, structures.ProductRelated{
			Id_product: id,
			Name:       name,
			Price:      price,
			Picture:    picture,
		})
	}
	//close the rows

	//create the json response
	return productRelated
}

func Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id_product := vars["id_product"] // the book title slug

	db := utils.DbConnect()

	productData := getOneProduct(w, db, id_product)
	productRelated := getRelatedProduct(db, id_product)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(structures.Product{
		ProductData:    productData,
		ProductRelated: productRelated,
	})
}
