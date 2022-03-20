package CRUD

import (
	databaseTools "inprinte/backend/database"
	structures "inprinte/backend/structures"
	utils "inprinte/backend/utils"
	"log"
	"strings"

	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	//manage url vars here
	urlFilter, ok := r.URL.Query()["filter"]
	if !ok || len(urlFilter[0]) < 1 {
		log.Println("Url Param 'filter' is missing")
		return
	}

	urlRange, ok := r.URL.Query()["range"]
	if !ok || len(urlRange[0]) < 1 {
		log.Println("Url Param 'range' is missing")
		return
	}
	urlRange[0] = urlRange[0][1 : len(urlRange[0])-1]

	urlSort, ok := r.URL.Query()["sort"]
	if !ok || len(urlSort[0]) < 1 {
		log.Println("Url Param 'sort' is missing")
		return
	}
	urlSort[0] = urlSort[0][1 : len(urlSort[0])-1]
	urlSort = strings.Split(urlSort[0], ",")
	urlSort[0] = urlSort[0][1 : len(urlSort[0])-1]
	urlSort[1] = urlSort[1][1 : len(urlSort[1])-1]

	if urlSort[0] == "id" {
		urlSort[0] = "id_user"
	}

	// filters := filter[0]
	log.Println("Url Params are : " + string(urlFilter[0]) + string(urlRange[0]) + string(urlSort[0]))

	db := databaseTools.DbConnect()

	//all sql queries for this path
	sqlQuery := "SELECT id_user, email, password, first_name AS firstname, last_name AS lastname, phone FROM user ORDER BY " + urlSort[0] + " " + urlSort[1] + " LIMIT " + urlRange[0]
	log.Println(sqlQuery)
	var id_user int
	var email, password, firstname, lastname, phone string
	var allUsers []structures.User

	rows := databaseTools.GetRequest(db, sqlQuery)
	for rows.Next() {
		err := rows.Scan(&id_user, &email, &password, &firstname, &lastname, &phone)

		// check errors
		utils.CheckErr(err)

		allUsers = append(allUsers, structures.User{
			Id_user:   id_user,
			Email:     email,
			Password:  password,
			Firstname: firstname,
			Lastname:  lastname,
			Phone:     phone,
		})
	}

	//----- encode the json response -----\\
	var response = structures.JsonResponseUsers{
		Type: "success",
		Data: allUsers,
	}
	json.NewEncoder(w).Encode(response)
}
