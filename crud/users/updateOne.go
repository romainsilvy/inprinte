package crud

import (
	"encoding/json"
	"fmt"
	structures "inprinteBackoffice/structures"
	utils "inprinteBackoffice/utils"
	"net/http"
	"strconv"
)

func UpdateOne(w http.ResponseWriter, r *http.Request) {
	utils.SetCorsHeaders(&w)

	// parse json from put Request
	var oneUser structures.OneUser
	err := json.NewDecoder(r.Body).Decode(&oneUser)
	if err != nil {
		fmt.Println(err)
	}

	// connect the database
	db := utils.DbConnect()

	// create the sql query
	sqlQuery := ("UPDATE user INNER JOIN address ON user.id_address = address.id SET first_name = '" + oneUser.Firstname + "', last_name = '" + oneUser.Lastname + "', email = '" + oneUser.Email + "', phone = '" + oneUser.Phone + "', is_alive = " + strconv.FormatBool(oneUser.IsAlive) + ", street = '" + oneUser.Address.Street + "', city = '" + oneUser.Address.City + "', state = '" + oneUser.Address.State + "', country = '" + oneUser.Address.Country + "', zip_code = '" + oneUser.Address.ZipCode + "', id_role = (SELECT role.id FROM role WHERE role.role = '" + oneUser.Role + "') WHERE user.id = '" + strconv.Itoa(oneUser.Id) + "'; ")

	// execute the sql query
	_, err = db.Exec(sqlQuery)
	utils.CheckErr(err)

	// create the json response
	json.NewEncoder(w).Encode(oneUser)
}
