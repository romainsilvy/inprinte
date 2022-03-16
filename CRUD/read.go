package CRUD

import (
	"encoding/json"
	"fmt"
	databaseTools "inprinte/backend/database"
	utils "inprinte/backend/utils"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Getting users ...")
	db := databaseTools.DbConnect()
	fmt.Println("DB OK")

	//get users\\

	rows, err := db.Query("SELECT id_user, email, password, firstname, lastname, phone FROM user")
	fmt.Println("DB req ok ")

	// check errors
	utils.CheckErr(err)
	fmt.Println("check err ok")

	// var response []JsonResponse
	var users []User

	// Foreach user
	for rows.Next() {
		var id_user int
		//var photo string
		var email string
		var password string
		var firstname string
		var lastname string
		var phone string

		err = rows.Scan(&id_user, &email, &password, &firstname, &lastname, &phone)

		// check errors
		utils.CheckErr(err)

		users = append(users, User{
			Id_user: id_user,
			//photo: photo,
			Email:     email,
			Password:  password,
			Firstname: firstname,
			Lastname:  lastname,
			Phone:     phone,
		})
	}

	var response = AllUsers{AllUsers: users}
	json.NewEncoder(w).Encode(response)
}

type User struct {
	Id_user   int    `json:"id_user"`
	Photo     string `json:"photo"`
	Email     string `json:"email"`
	Password  string `json:"password "`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Phone     string `json:"phone"`
	//adress		[]Adress 	`json:"adress"`
}

type AllUsers struct {
	AllUsers []User `json:"allUsers"`
}
