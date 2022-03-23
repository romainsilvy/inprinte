package CRUD

import (
	"encoding/json"
	"fmt"
	databaseTools "inprinte/backend/database"
	structures "inprinte/backend/structures"
	utils "inprinte/backend/utils"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	db := databaseTools.DbConnect()
	fmt.Println("Getting products ...")
	vars := mux.Vars(r)
	id := vars["id"]

	//Get products infos
	rows, err := db.Query("SELECT name, description, price, AVG(stars_number) AS MOYENNE, picture.url, product_file.id FROM product INNER JOIN rate ON rate.id = product.id INNER JOIN product_picture ON product_picture.id = product.id INNER JOIN picture ON picture.id = product_picture.id INNER JOIN product_file ON product_file.id = product.id WHERE product.id = " + id + " GROUP BY product.id;")

	// check errors
	utils.CheckErr(err)

	// var response []JsonResponse
	var products []structures.Product

	// Foreach product
	for rows.Next() {
		var name string
		var description string
		var price int
		var stars_number float64
		var picture_url string
		var product_file string

		//var picture string
		err = rows.Scan(&name, &description, &price, &stars_number, &picture_url, &product_file)

		// check errors
		utils.CheckErr(err)

		products = append(products, structures.Product{
			Name:         name,
			Description:  description,
			Price:        price,
			Stars_number: stars_number,
			Picture_url:  picture_url,
			Product_file: product_file,
		})
	}

	var response = structures.JsonResponseProduct{
		Type: "success",
		Data: products,
	}

	if response.Data == nil {
		w.WriteHeader(404)
	} else {
		json.NewEncoder(w).Encode(response)
	}
}
