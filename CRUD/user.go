package CRUD

import (
	databaseTools "inprinte/backend/database"
	structures "inprinte/backend/structures"
	utils "inprinte/backend/utils"

	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	db := databaseTools.DbConnect()

	//all sql queries for this path
	sqlQuery := "SELECT id_user, email, password, first_name AS firstname, last_name AS lastname, phone FROM user"

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
