//package crud is the package that contains all the functions to interact with the database
package crud

import (
	"database/sql"
	"encoding/json"
	structures "inprinte/backend/structures"
	utils "inprinte/backend/utils"
	"strconv"

	"net/http"

	"github.com/gorilla/mux"
)

//getAllPictures retrieve all the pictures of a product
func getAllPictures(db *sql.DB, id_product string) []string {

	//global vars
	var allPictures []string

	//execute the sql query and check errors
	rows, err := db.Query("SELECT url FROM picture WHERE id IN (SELECT id_picture FROM product_picture WHERE id_product IN (SELECT id FROM product WHERE pending_validation = false AND is_alive = true AND product.id = ? ) ); ", id_product)
	utils.CheckErr(err)

	//parse the query
	for rows.Next() {
		//global vars
		var picture string

		//retrieve the values and check errors
		err = rows.Scan(&picture)
		utils.CheckErr(err)

		//add the values to the response
		allPictures = append(allPictures, picture)
	}
	//close the rows

	//return all the pictures
	return allPictures

}

//getOneProduct retrieve all the data of a product
func getOneProduct(w http.ResponseWriter, db *sql.DB, id_product string) structures.ProductData {
	//global vars
	var id int
	var name, description string
	var price float64

	//execute the sql query and check errors
	row := db.QueryRow(`SELECT product.id, name, description, price FROM product INNER JOIN rate ON rate.id = product.id WHERE product.id = ? AND pending_validation = false AND product.is_alive = true;`, id_product)
	err := row.Scan(&id, &name, &description, &price)
	if err == sql.ErrNoRows {
		w.WriteHeader(404)
	} else {
		utils.CheckErr(err)
	}

	//retrieve all the pictures of the product
	allPictures := getAllPictures(db, id_product)

	//create the json response
	return structures.ProductData{
		Id_product:  id,
		Name:        name,
		Description: description,
		Price:       price,
		Pictures:    allPictures,
	}
}

//getRelatedProduct retrieve all the related products of a product
func getRelatedProduct(r *http.Request, db *sql.DB, id_product string) []structures.ProductRelated {
	//global vars
	var productRelated []structures.ProductRelated

	//retrieve the range parameter from the url
	rangeBy := utils.GetRangeParam(r)

	//execute the sql query and check errors
	rows, err := db.Query("SELECT product.id, product.name, product.price, AVG(rate.stars_number) AS rate FROM product INNER JOIN rate ON rate.id_product = product.id WHERE product.is_alive = true AND product.pending_validation = false AND product.id_category = (SELECT category.id FROM category INNER JOIN product ON product.id_category = category.id WHERE product.id = " + id_product + " ) GROUP BY product.id ORDER BY rate DESC " + rangeBy)
	utils.CheckErr(err)

	//parse the query
	for rows.Next() {
		//global vars
		var name, picture string
		var id int
		var price, rate float64

		//retrieve the values and check errors
		err = rows.Scan(&id, &name, &price, &rate)
		utils.CheckErr(err)
		picture = getProductPicture(db, strconv.Itoa(id))

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

//getProductPicture retrieve the first picture of a product
func getProductPicture(db *sql.DB, id_product string) string {
	//global vars
	var url string

	//create the sql query
	sqlQuery := ("SELECT picture.url FROM picture INNER JOIN product_picture ON product_picture.id_picture = picture.id INNER JOIN product ON product.id = product_picture.id_product WHERE product.id = " + id_product + " LIMIT 1;")

	//execute the sql query and check errors
	err := db.QueryRow(sqlQuery).Scan(&url)
	utils.CheckErr(err)

	//return the picture
	return url
}

//Get retrieve all the usefull informations for the product page
func Get(w http.ResponseWriter, r *http.Request) {
	//set the header
	utils.SetCorsHeaders(&w)

	//get the url parameters
	vars := mux.Vars(r)
	id_product := vars["id_product"]

	//connect to the database
	db := utils.DbConnect()

	//retrieve the data
	productData := getOneProduct(w, db, id_product)
	productRelated := getRelatedProduct(r, db, id_product)

	//create the json response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(structures.Product{
		ProductData:    productData,
		ProductRelated: productRelated,
	})
}
