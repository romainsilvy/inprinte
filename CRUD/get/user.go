package CRUD

import (
	"encoding/json"
	"fmt"
	databaseTools "inprinte/backend/database"
	structures "inprinte/backend/structures"
	"inprinte/backend/utils"

	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// func GetUser(w http.ResponseWriter, r *http.Request) {
// 	db := databaseTools.DbConnect()

// 	userData := getUserData(db, userDataQuery)

// 	userFavorite := getUserFavorite(db, userFavoriteProductQuery)

// 	userOldCommand := getUserOldCommand(db, userOldCommandQuery)

// 	var response = structures.JsonResponseUsers{
// 		Type:       "success",
// 		UserData:   userData,
// 		Favorite:   userFavorite,
// 		OldCommand: userOldCommand,
// 	}
// 	json.NewEncoder(w).Encode(response)
// }

func GetUserData(w http.ResponseWriter, r *http.Request) {
	db := databaseTools.DbConnect()
	fmt.Println("Getting products ...")
	vars := mux.Vars(r)
	id_user := vars["id_user"]

	rows, err := db.Query("SELECT  first_name, last_name, email, password, phone, street, city, state, country, zip_code FROM user INNER JOIN address ON user.id = address.id WHERE user.id = " + id_user + " ; ")
	utils.CheckErr(err)

	var users []structures.UserData

	for rows.Next() {
		var firstname, lastname, email, password, phone, street, city, state, country, zip_code string
		err = rows.Scan(&firstname, &lastname, &email, &password, &phone, &street, &city, &state, &country, &zip_code)

		utils.CheckErr(err)

		users = append(users, structures.UserData{
			Firstname: firstname,
			Lastname:  lastname,
			Email:     email,
			Password:  password,
			Phone:     phone,
			Address: structures.Address{
				Street:  street,
				City:    city,
				State:   state,
				Country: country,
				ZipCode: zip_code,
			},
		})
	}

	var favorite := getUserFavorite()
	var response = structures.JsonResponseUsers{
		Type:     "success",
		UserData: users,
	}

	if response.UserData == nil {
		w.WriteHeader(404)
	} else {
		json.NewEncoder(w).Encode(response)
	}
}

func GetUserFavorite(w http.ResponseWriter, r *http.Request) {
	db := databaseTools.DbConnect()
	vars := mux.Vars(r)
	id_user := vars["id_user"]

	rows, err := db.Query("SELECT DISTINCT user.id AS id_user, product.id AS id_product, name, price, picture.url FROM product INNER JOIN favorite ON favorite.id_product = product.id INNER JOIN user ON user.id = favorite.id_user INNER JOIN product_picture ON product_picture.id_product = product.id INNER JOIN picture ON picture.id = product_picture.id_picture AND pending_validation = false AND product.is_alive = true WHERE user.id = " + id_user + ";")
	utils.CheckErr(err)

	var userFavoriteProduct []structures.UserProductData

	for rows.Next() {
		var picture, name, price string
		var id_user, id_product int
		err = rows.Scan(&id_user, &id_product, &name, &price, &picture)

		utils.CheckErr(err)

		userFavoriteProduct = append(userFavoriteProduct, structures.UserProductData{
			Picture: picture,
			Name:    name,
			Price:   price,
		})
	}

	var response = structures.JsonResponseUsers{
		Type:     "success",
		Favorite: userFavoriteProduct,
	}

	return userFavoriteProduct
}

func GetUserOldCommand(w http.ResponseWriter, r *http.Request) {
	db := databaseTools.DbConnect()
	vars := mux.Vars(r)
	id_user := vars["id_user"]

	rows, err := db.Query("SELECT url, name, price FROM product INNER JOIN picture ON product.id = picture.id WHERE user.id = " + id_user + " ; ")
	utils.CheckErr(err)

	var userOldCommand []structures.UserProductData

	for rows.Next() {
		var picture, name, price string
		err = rows.Scan(&picture, &name, &price)

		utils.CheckErr(err)

		userOldCommand = append(userOldCommand, structures.UserProductData{
			Picture: picture,
			Name:    name,
			Price:   price,
		})
	}

	var response = structures.JsonResponseUsers{
		Type:       "success",
		OldCommand: userOldCommand,
	}

	if response.UserData == nil {
		w.WriteHeader(404)
	} else {
		json.NewEncoder(w).Encode(response)
	}

}
