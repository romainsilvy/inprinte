package authentication

import (
	"encoding/json"
	"inprinte/backend/utils"
	"log"
	"net/http"
)

func checkIfUserExists(email string) bool {
	//open the database connection
	db := utils.DbConnect()

	//global vars
	var id int

	//check if the user exists in the database, if so return true
	sql := "SELECT id FROM user WHERE email = \"" + email + "\";"
	err := db.QueryRow(sql).Scan(&id)
	// utils.CheckErr(err)
	if err == nil {
		return true
	} else {
		return false
	}
}

type AddUser struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Phone     string `json:"phone"`
}

func Signup(w http.ResponseWriter, r *http.Request) {
	//decode the json response
	var addUser AddUser
	json.NewDecoder(r.Body).Decode(&addUser)

	//user credentials
	firstname := addUser.Firstname
	lastname := addUser.Lastname
	email := addUser.Email
	password := addUser.Password
	phone := addUser.Phone

	userExists := checkIfUserExists(email)
	if userExists {
		log.Println("User already exists")
	} else {
		//connect the database
		db := utils.DbConnect()

		sql := "INSERT INTO user (first_name, last_name, email, password, phone, is_alive, id_role) VALUES (\"" + firstname + "\", \"" + lastname + "\", \"" + email + "\", \"" + utils.HashPassword(password) + "\", \"" + phone + "\", true, 2);"
		_, err := db.Exec(sql)
		utils.CheckErr(err)
	}
}
