package crud

import (
	"database/sql"
	"encoding/json"
	utils "inprinte/backend/utils"
	"log"

	"net/http"
)

type ItemCart struct {
	Id       string `json:"id"`
	Quantity int    `json:"quantity"`
}

type ReturnItemCart struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Picture     string  `json:"picture"`
	Quantity    int     `json:"quantity"`
}

type AllReturnCart struct {
	AllItems   []ReturnItemCart `json:"allItems"`
	TotalPrice float64          `json:"totalPrice"`
}

func getProductPicture(db *sql.DB, id_product string) string {
	//global vars
	var url string

	//create the sql query
	sqlQuery := ("SELECT picture.url FROM picture INNER JOIN product_picture ON product_picture.id_picture = picture.id INNER JOIN product ON product.id = product_picture.id_product WHERE product.id = " + id_product + ";")

	err := db.QueryRow(sqlQuery).Scan(&url)
	utils.CheckErr(err)

	return url
}

//Get returns informations of the given product
func Get(w http.ResponseWriter, r *http.Request) {
	//retrieve the request url parameters
	cart := r.URL.Query()["cart"]
	allCart := []ItemCart{}
	err := json.Unmarshal([]byte(cart[0]), &allCart)

	if err != nil {
		log.Println(err)
	}

	//set the cors headers
	utils.SetCorsHeaders(&w)

	db := utils.DbConnect()
	//global vars
	var name, description string
	var picture string
	var id int
	var price float64
	var allItemCart []ReturnItemCart
	var totalPrice float64

	for _, item := range allCart {
		//execute the sql query and check errors
		err := db.QueryRow("SELECT product.id, name, price, description FROM product WHERE product.id = ?", item.Id).Scan(&id, &name, &price, &description)
		if err != nil {
			if err == sql.ErrNoRows {
				log.Println("No rows were returned!")
			} else {
				log.Println(err)
			}
		}

		picture = getProductPicture(db, item.Id)

		allItemCart = append(allItemCart, ReturnItemCart{
			Id:          item.Id,
			Name:        name,
			Picture:     picture,
			Description: description,
			Price:       price,
			Quantity:    item.Quantity,
		})
		totalPrice += price * float64(item.Quantity)
	}

	//create the json response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(AllReturnCart{
		AllItems:   allItemCart,
		TotalPrice: totalPrice,
	})
}
