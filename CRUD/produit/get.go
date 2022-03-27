package crud

import (
	"database/sql"
	"encoding/json"
	structures "inprinte/backend/structures"
	utils "inprinte/backend/utils"

	"net/http"
)

func getOneProduct(db *sql.DB, id_product string) structures.ProductData {
	//global vars
	var id int
	var name, description, picture, product_file string
	var price float64

	//execute the sql query and check errors
	row := db.QueryRow(`SELECT product.id, name, description, price, picture.url, product_file.url FROM product INNER JOIN rate ON rate.id = product.id INNER JOIN product_picture ON product_picture.id = product.id INNER JOIN picture ON picture.id = product_picture.id INNER JOIN product_file ON product_file.id = product.id WHERE product.id = ? AND pending_validation = false AND product.is_alive = true GROUP BY product.id`, id_product)
	err := row.Scan(&id, &name, &description, &price, &picture, &product_file)
	utils.CheckErr(err)

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

func GetProduct(w http.ResponseWriter, r *http.Request) {
	db := utils.DbConnect()

	productData := getOneProduct(db, "1")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(structures.Product{
		ProductData: productData,
	})
}
