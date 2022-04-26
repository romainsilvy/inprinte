package crud

import (
	"encoding/json"
	"fmt"
	structures "inprinteBackoffice/structures"
	utils "inprinteBackoffice/utils"
	"net/http"
	"strconv"
)

func Update(w http.ResponseWriter, r *http.Request) {
	//create cors header
	utils.SetCorsHeaders(&w)

	if r.Method == "PUT" {

		if utils.Securized(w, r) {
			// parse json from put Request
			var user structures.GetUser
			err := json.NewDecoder(r.Body).Decode(&user)
			if err != nil {
				fmt.Println(err)
			}

			// connect the database
			db := utils.DbConnect()

			// create the sql query
			sqlQuery := ("UPDATE user INNER JOIN address ON user.id_address = address.id SET first_name = '" + user.Firstname + "', last_name = '" + user.Lastname + "', email = '" + user.Email + "', phone = '" + user.Phone + "', is_alive = " + strconv.FormatBool(user.IsAlive) + ", street = '" + user.Address.Street + "', city = '" + user.Address.City + "', state = '" + user.Address.State + "', country = '" + user.Address.Country + "', zip_code = '" + user.Address.ZipCode + "', id_role = (SELECT role.id FROM role WHERE role.role = '" + user.Role + "') WHERE user.id = '" + strconv.Itoa(user.Id) + "'; ")

			// execute the sql query
			_, err = db.Exec(sqlQuery)
			utils.CheckErr(err)

			// close the database connection
			db.Close()

			// create the json response
			json.NewEncoder(w).Encode(user)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
	}
}
