package crud

import (
	"encoding/json"
	"inprinteBackoffice/cookie"
	structures "inprinteBackoffice/structures"
	utils "inprinteBackoffice/utils"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	//create cors header
	utils.SetCorsHeaders(&w)

	if r.Method == "GET" {
		if cookie.Securized(w, r) {
			//global vars
			var users []structures.GetUsers
			//connect the database
			db := utils.DbConnect()

			//get filters values and update the sqlQuery
			orderBy, rangeBy := utils.GetAllParams(r, "user")
			sqlQuery := "SELECT user.id, email, password, first_name AS firstname, last_name AS lastname, phone, is_alive, street, city, state, country, zip_code, role.role  FROM user INNER JOIN address ON user.id_address = address.id INNER JOIN role ON user.id_role = role.id " + orderBy + " " + rangeBy

			//execute the sql query and check errors
			rows, err := db.Query(sqlQuery)
			utils.CheckErr(err)

			//parse the query
			for rows.Next() {
				//global vars
				var firstname, lastname, email, password, phone, street, city, state, country, zip_code, role string
				var is_alive bool
				var id int

				//retrieve the values and check errors
				err = rows.Scan(&id, &email, &password, &firstname, &lastname, &phone, &is_alive, &street, &city, &state, &country, &zip_code, &role)
				utils.CheckErr(err)

				//add the values to the response
				users = append(users, structures.GetUsers{
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
				})
			}
			//close the rows
			rows.Close()

			//close the database connection
			db.Close()

			//create the json response
			utils.SetXTotalCountHeader(&w, len(users))
			json.NewEncoder(w).Encode(users)

		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
	}
}
