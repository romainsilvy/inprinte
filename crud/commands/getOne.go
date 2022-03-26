package crud

// import "net/http"

// func GetOne(w http.ResponseWriter, r *http.Request) {
// 	//global vars
// 	var status string

// 	//connect the database
// 	db := utils.DbConnect()

// 	//create cors header
// 	utils.SetCorsHeaders(&w)

// 	//get url values
// 	vars := mux.Vars(r)
// 	id_user := vars["id_user"]

// 	//create the sql query
// 	sqlQuery := ("SELECT  command.is AS id, command.status FROM command INNER JOIN user ON user.id = command.id_user WHERE user.id = " + id_user + ";")

// 	//execute the sql query
// 	row := db.QueryRow(sqlQuery)

// 	//parse the query
// 	//retrieve the values and check errors
// 	err := row.Scan(&firstname, &lastname, &email, &password, &phone, &is_alive, &street, &city, &state, &country, &zip_code, &role)
// 	utils.CheckErr(err)

// 	//add the values to the response
// 	oneUser = structures.OneUser{
// 		Firstname: firstname,
// 		Lastname:  lastname,
// 		Email:     email,
// 		Phone:     phone,
// 		IsAlive:   is_alive,
// 		Role:      role,
// 		Address: structures.Address{
// 			Street:  street,
// 			City:    city,
// 			State:   state,
// 			Country: country,
// 			ZipCode: zip_code,
// 		},
// 	}

// 	//create the json response
// 	json.NewEncoder(w).Encode(oneUser)
// }