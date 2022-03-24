package CRUD

import (
	"encoding/json"
	"fmt"
	databaseTools "inprinte/backend/database"
	structures "inprinte/backend/structures"
	utils "inprinte/backend/utils"
	"log"
	"strings"

	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func GetUserData(w http.ResponseWriter, r *http.Request) {
	db := databaseTools.DbConnect()
	fmt.Println("Getting products ...")
	vars := mux.Vars(r)
	id_user := vars["id_user"]

	rows, err := db.Query("SELECT  first_name, last_name, email, password, phone, street, city, state, country, zip_code FROM user INNER JOIN address ON user.id = address.id WHERE user.id = " + id_user + " ; ")

	utils.CheckErr(err)

	var users []structures.User

	for rows.Next() {
		var firstname, lastname, email, password, phone, street, city, state, country, zip_code string
		err = rows.Scan(&firstname, &lastname, &email, &password, &phone, &street, &city, &state, &country, &zip_code)

		utils.CheckErr(err)

		users = append(users, structures.User{
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

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	//manage url vars here
	urlFilter, ok := r.URL.Query()["filter"]
	if !ok || len(urlFilter[0]) < 1 {
		log.Println("Url Param 'filter' is missing")
	}

	urlRange, ok := r.URL.Query()["range"]
	if !ok || len(urlRange[0]) < 1 {
		log.Println("Url Param 'range' is missing")
	}
	urlRange[0] = urlRange[0][1 : len(urlRange[0])-1]

	urlSort, ok := r.URL.Query()["sort"]
	if !ok || len(urlSort[0]) < 1 {
		log.Println("Url Param 'sort' is missing")
	}
	urlSort[0] = urlSort[0][1 : len(urlSort[0])-1]
	urlSort = strings.Split(urlSort[0], ",")
	urlSort[0] = urlSort[0][1 : len(urlSort[0])-1]
	urlSort[1] = urlSort[1][1 : len(urlSort[1])-1]

	// filters := filter[0]
	log.Println("Url Params are : " + string(urlFilter[0]) + string(urlRange[0]) + string(urlSort[0]))

	db := databaseTools.DbConnect()

	//all sql queries for this path
	sqlQuery := "SELECT user.id, email, password, first_name AS firstname, last_name AS lastname, phone, street, city, state, country, zip_code  FROM user INNER JOIN address ON user.id_address = address.id ORDER BY " + urlSort[0] + " " + urlSort[1] + " LIMIT " + urlRange[0]
	log.Println(sqlQuery)
	var allUsers []structures.User

	rows, err := db.Query(sqlQuery)
	utils.CheckErr(err)

	for rows.Next() {
		var firstname, lastname, email, password, phone, street, city, state, country, zip_code string
		var id int
		err = rows.Scan(&id, &email, &password, &firstname, &lastname, &phone, &street, &city, &state, &country, &zip_code)

		utils.CheckErr(err)

		allUsers = append(allUsers, structures.User{
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

	//----- encode the json response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(allUsers)
}
