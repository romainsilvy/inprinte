package crud

import (
	"encoding/json"
	"inprinteBackoffice/structures"
	utils "inprinteBackoffice/utils"
	"net/http"

	"github.com/gorilla/mux"
)

func GetUser(w http.ResponseWriter, r *http.Request) {

	//global vars
	var firstname, lastname, email, phone, street, city, state, country, zip_code, role string
	var is_alive bool
	var oneUser structures.GetUser
	var id int

	//connect the database
	db := utils.DbConnect()

	//create cors header
	utils.SetCorsHeaders(&w)

	//get url values
	vars := mux.Vars(r)
	id_user := vars["id_user"]

	//create the sql query
	sqlQuery := ("SELECT user.id, first_name, last_name, email, phone, is_alive, street, city, state, country, zip_code, role.role FROM user INNER JOIN address ON user.id_address = address.id INNER JOIN role ON user.id_role = role.id WHERE user.id = " + id_user + " ;")

	err := db.QueryRow(sqlQuery).Scan(&id, &firstname, &lastname, &email, &phone, &is_alive, &street, &city, &state, &country, &zip_code, &role)
	utils.CheckErr(err)

	//add the values to the response
	oneUser = structures.GetUser{
		Id:        id,
		Firstname: firstname,
		Lastname:  lastname,
		Email:     email,
		Phone:     phone,
		IsAlive:   is_alive,
		Role:      role,
		Address: structures.Address{
			Street:  street,
			City:    city,
			State:   state,
			Country: country,
			ZipCode: zip_code,
		},
	}

	// Close the database connection
	db.Close()

	//create the json response
	json.NewEncoder(w).Encode(oneUser)

}
