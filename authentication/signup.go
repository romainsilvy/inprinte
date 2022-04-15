//package authentication provides all the functionality needed to authenticate users
package authentication

import (
	"encoding/json"
	"inprinte/backend/structures"
	"inprinte/backend/utils"
	"net/http"
)

//checkIfUserExists checks if the user exists in the database based on the email address
func checkIfUserExists(email string) bool {
	//open the database connection
	db := utils.DbConnect()

	//global vars
	var id int

	//check if the user exists in the database, if so return true
	sql := "SELECT id FROM user WHERE email = \"" + email + "\";"
	err := db.QueryRow(sql).Scan(&id)
	if err == nil {
		return true
	} else {
		return false
	}
}

//Signup creates a new user in the database
func Signup(w http.ResponseWriter, r *http.Request) {
	//set the cors headers
	utils.SetCorsHeaders(&w)

	//decode the json response
	var addUser structures.AddUser
	json.NewDecoder(r.Body).Decode(&addUser)

	//declare user credentials
	firstname := addUser.Firstname
	lastname := addUser.Lastname
	email := addUser.Email
	password := addUser.Password
	phone := addUser.Phone

	//check if the user exists in the database
	userExists := checkIfUserExists(email)
	if userExists {
		//return exists error
		json.NewEncoder(w).Encode("exists")
	} else {
		//connect the database
		db := utils.DbConnect()
		//execute the query and check errors
		sql := "INSERT INTO user (first_name, last_name, email, password, phone, is_alive, id_role) VALUES (\"" + firstname + "\", \"" + lastname + "\", \"" + email + "\", \"" + utils.HashPassword(password) + "\", \"" + phone + "\", true, 2);"
		_, err := db.Exec(sql)
		utils.CheckErr(err)

		//return created status
		json.NewEncoder(w).Encode("created")
	}
}
