package crud

import (
	"encoding/json"
	"inprinteBackoffice/structures"
	utils "inprinteBackoffice/utils"
	"net/http"

	"github.com/gorilla/mux"
)

func GetOne(w http.ResponseWriter, r *http.Request) {
	//global vars
	var firstname, lastname, email, phone, street, city, state, country, zip_code, role string
	var is_alive bool
	var oneUser structures.OneUser

	//connect the database
	db := utils.DbConnect()

	//create cors header
	utils.SetCorsHeaders(&w)

	//get url values
	vars := mux.Vars(r)
	id_user := vars["id_user"]

	//create the sql query
	sqlQuery := ("SELECT  first_name, last_name, email, phone, is_alive, street, city, state, country, zip_code, role.role FROM user INNER JOIN address ON user.id = address.id INNER JOIN role ON user.id_role = role.id WHERE user.id = " + id_user + ";")

	//execute the sql query
	row := db.QueryRow(sqlQuery)

	//parse the query
	//retrieve the values and check errors
	err := row.Scan(&firstname, &lastname, &email, &phone, &is_alive, &street, &city, &state, &country, &zip_code, &role)
	utils.CheckErr(err)

	//add the values to the response
	oneUser = structures.OneUser{
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

	//create the json response
	json.NewEncoder(w).Encode(oneUser)
}
