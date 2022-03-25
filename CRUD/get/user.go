package CRUD

import (
	"database/sql"
	"encoding/json"
	databaseTools "inprinte/backend/database"
	structures "inprinte/backend/structures"
	"inprinte/backend/utils"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	db := databaseTools.DbConnect()
	vars := mux.Vars(r)
	id_user := vars["id_user"]

	// all request SQL for path user \\
	userDataQuery := ("SELECT  first_name, last_name, email, password, phone, street, city, state, country, zip_code FROM user INNER JOIN address ON user.id = address.id WHERE user.id = " + id_user + ";")
	userFavoriteProductQuery := ("SELECT DISTINCT user.id AS id_user, product.id AS id_product, name, price, picture.url FROM product INNER JOIN favorite ON favorite.id_product = product.id INNER JOIN user ON user.id = favorite.id_user INNER JOIN product_picture ON product_picture.id_product = product.id INNER JOIN picture ON picture.id = product_picture.id_picture AND pending_validation = false AND product.is_alive = true WHERE user.id = " + id_user + ";")
	userOldCommandQuery := ("SELECT user.id AS id_user, name, price, picture.url, quantity FROM product INNER JOIN user ON product.id = user.id INNER JOIN command_line ON product.id = command_line.id INNER JOIN picture ON command_line.id = picture.id WHERE user.id = " + id_user + ";")

	// data of user \\
	userData := GetUserData(db, userDataQuery)

	// favorite product for user \\
	userFavorite := GetUserFavorite(db, userFavoriteProductQuery)

	// old command for user \\
	userOldCommand := GetUserOldCommand(db, userOldCommandQuery)

	//----- encode the json response -----\\
	var response = structures.JsonResponseUsers{
		Data:            userData,
		FavoriteProduct: userFavorite,
		OldCommand:      userOldCommand,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetUserData(db *sql.DB, sqlQuery string) []structures.UserData {
	var firstname, lastname, email, password, phone, street, city, state, country, zip_code string
	var users []structures.UserData

	rows, err := db.Query(sqlQuery)
	utils.CheckErr(err)

	for rows.Next() {
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
	return users
}

func GetUserFavorite(db *sql.DB, sqlQuery string) []structures.UserFavoriteProduct {
	var userFavoriteProduct []structures.UserFavoriteProduct

	rows, err := db.Query(sqlQuery)
	utils.CheckErr(err)

	for rows.Next() {
		var picture, name, price string
		var id_user, id_product int
		err = rows.Scan(&id_user, &id_product, &name, &price, &picture)

		utils.CheckErr(err)

		userFavoriteProduct = append(userFavoriteProduct, structures.UserFavoriteProduct{
			Picture: picture,
			Name:    name,
			Price:   price,
		})
	}
	return userFavoriteProduct
}

func GetUserOldCommand(db *sql.DB, sqlQuery string) []structures.UserOldCommand {
	var userOldCommand []structures.UserOldCommand
	var picture, name, price string
	var id_user, quantity int

	rows, err := db.Query(sqlQuery)
	utils.CheckErr(err)

	for rows.Next() {
		err = rows.Scan(&id_user, &name, &price, &picture, &quantity)

		utils.CheckErr(err)

		userOldCommand = append(userOldCommand, structures.UserOldCommand{
			Picture:  picture,
			Name:     name,
			Price:    price,
			Quantity: quantity,
		})
	}
	return userOldCommand
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	containsOrder := true
	containsSort := true
	containsStart := true
	containsEnd := true

	sqlQuery := "SELECT user.id, email, password, first_name AS firstname, last_name AS lastname, phone, street, city, state, country, zip_code  FROM user INNER JOIN address ON user.id_address = address.id"

	//manage url vars here
	urlOrder, ok := r.URL.Query()["_order"]
	if !ok || len(urlOrder[0]) < 1 {
		log.Println("Url Param 'order' is missing")
		containsOrder = false
	}

	urlSort, ok := r.URL.Query()["_sort"]
	if !ok || len(urlSort[0]) < 1 {
		log.Println("Url Param 'sort' is missing")
		containsSort = true
	}

	urlStart, ok := r.URL.Query()["_start"]
	if !ok || len(urlStart[0]) < 1 {
		log.Println("Url Param 'start' is missing")
		containsStart = false
	}

	urlEnd, ok := r.URL.Query()["_end"]
	if !ok || len(urlEnd[0]) < 1 {
		log.Println("Url Param 'End' is missing")
		containsEnd = false
	}

	if containsOrder && containsSort {
		sqlQuery += " ORDER BY " + urlSort[0] + " " + urlOrder[0]
	}

	if containsStart && containsEnd {
		sqlQuery += " LIMIT " + urlStart[0] + "," + urlEnd[0]
	}

	// filters := filter[0]
	// log.Println("Url Params are : " + string(urlFilter[0]) + string(urlRange[0]) + string(urlSort[0]))

	db := databaseTools.DbConnect()

	//all sql queries for this path
	// sqlQuery := "SELECT user.id, email, password, first_name AS firstname, last_name AS lastname, phone, street, city, state, country, zip_code  FROM user INNER JOIN address ON user.id_address = address.id ORDER BY " + urlSort[0] + " " + urlSort[1] + " LIMIT " + urlRange[0]
	log.Println(sqlQuery)
	var allUsers []structures.UserData

	rows, err := db.Query(sqlQuery)
	utils.CheckErr(err)

	for rows.Next() {
		var firstname, lastname, email, password, phone, street, city, state, country, zip_code string
		var id int
		err = rows.Scan(&id, &email, &password, &firstname, &lastname, &phone, &street, &city, &state, &country, &zip_code)

		utils.CheckErr(err)

		allUsers = append(allUsers, structures.UserData{
			Id:        id,
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

	// res := structures.JsonResponseUsers{
	// 	Data:  allUsers,
	// 	Total: 100,
	// 	Date:  time.Now().Format("2006-01-02 15:04:05"),
	// }

	//----- encode the json response
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Expose-Headers", "X-Total-Count")
	w.Header().Set("X-Total-Count", "234")
	json.NewEncoder(w).Encode(allUsers)

	// w.Header().Set("Access-Control-Allow-Origin", "*")

	// w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	// // return "OKOK"
	// json.NewEncoder(w).Encode("OKOK")
}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
