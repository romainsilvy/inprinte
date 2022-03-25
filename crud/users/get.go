package crud

import (
	"encoding/json"
	structures "inprinteBackoffice/structures"
	utils "inprinteBackoffice/utils"

	"log"
	"net/http"
)

func GetAllParams(r *http.Request) (string, string) {
	containsOrder := true
	containsSort := true
	containsStart := true
	containsEnd := true
	rangeBy := ""
	orderBy := ""

	urlOrder, ok := r.URL.Query()["_order"]
	if !ok || len(urlOrder[0]) < 1 {
		log.Println("Url Param 'order' is missing")
		containsOrder = false
	}

	urlSort, ok := r.URL.Query()["_sort"]
	if !ok || len(urlSort[0]) < 1 {
		log.Println("Url Param 'sort' is missing")
		containsSort = true
	}

	urlStart, ok := r.URL.Query()["_start"]
	if !ok || len(urlStart[0]) < 1 {
		log.Println("Url Param 'start' is missing")
		containsStart = false
	}

	urlEnd, ok := r.URL.Query()["_end"]
	if !ok || len(urlEnd[0]) < 1 {
		log.Println("Url Param 'End' is missing")
		containsEnd = false
	}

	if containsOrder && containsSort {
		orderBy = " ORDER BY " + urlSort[0] + " " + urlOrder[0]
	}

	if containsStart && containsEnd {
		rangeBy = " LIMIT " + urlStart[0] + "," + urlEnd[0]
	}
	return rangeBy, orderBy

}

func GetAll(w http.ResponseWriter, r *http.Request) {
	//global vars
	var allUsers []structures.User

	//connect the database
	db := utils.DbConnect()

	//create cors header
	utils.SetCorsHeaders(&w)

	//get filters values and update the sqlQuery
	orderBy, rangeBy := GetAllParams(r)
	sqlQuery := "SELECT user.id, email, password, first_name AS firstname, last_name AS lastname, phone, street, city, state, country, zip_code  FROM user INNER JOIN address ON user.id_address = address.id" + orderBy + rangeBy

	//execute the sql request and check errors
	rows, err := db.Query(sqlQuery)
	utils.CheckErr(err)

	//parse the rows
	for rows.Next() {
		//global vars
		var firstname, lastname, email, password, phone, street, city, state, country, zip_code string
		var id int

		//retrieve the values and check errors
		err = rows.Scan(&id, &email, &password, &firstname, &lastname, &phone, &street, &city, &state, &country, &zip_code)
		utils.CheckErr(err)

		//add a new user object to the array with the corresponding values
		allUsers = append(allUsers, structures.User{
			Id:        id,
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
	//close the rows

	//create the json response
	utils.SetXTotalCountHeader(&w, len(allUsers))
	json.NewEncoder(w).Encode(allUsers)
}
