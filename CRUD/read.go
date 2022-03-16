package CRUD

import (
	utils "inprinte/backend/utils"
	"fmt"
	"net/http"
	"encoding/json"
	databaseTools "inprinte/backend/database"
	_ "github.com/go-sql-driver/mysql"
)


func GetUsers(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Getting users ...")
    db := databaseTools.DbConnect()
	fmt.Println("DB OK")

    //get users\\ 

    rows, err := db.Query("SELECT * FROM users")
	fmt.Println("DB req ok ")

    // check errors
    utils.CheckErr(err)
	fmt.Println("check err ok")

    // var response []JsonResponse
    var users []User

    // Foreach user
    for rows.Next() {
		var id_user int	
		var photo string
		var email string
		var password string
		var firstname string
		var lastname string
		var phone string

        err = rows.Scan(&id_user, &photo, &email, &password, &firstname, &lastname, &phone)

        // check errors
        utils.CheckErr(err)

        users = append(users, User{
			id_user: id_user, 
			photo: photo,
			email: email, 
			password: password,
			firstname: firstname, 
			lastname: lastname,
			phone: phone,
		})
    }

    var response = AllUsers{allUsers: users}

    json.NewEncoder(w).Encode(response)
}


type User struct {
	id_user		int			`json:"id_user"`
	photo		string		`json:"photo"`
	email		string		`json:"email"`
	password	string		`json:"password "`
	firstname	string		`json:"firstname"`
	lastname	string		`json:"lastname"`
	phone		string		`json:"phone"`
	//adress		[]Adress 	`json:"adress"`
}

type AllUsers struct {
	allUsers	[]User	`json:"allUsers"`
}