package crud

import (
	"encoding/json"
	structures "inprinteBackoffice/structures"
	utils "inprinteBackoffice/utils"
	"net/http"
	"strconv"
)

func InsertOne(w http.ResponseWriter, r *http.Request) {
	utils.SetCorsHeaders(&w)
	if r.Method == "POST" {
		// global variables
		var response = structures.InsertOneUser{}
		var oneUser = structures.CreateOneUser{}
		var lastInsertID int
		var lastInsertIDAddress int

		// get body
		err := json.NewDecoder(r.Body).Decode(&oneUser)
		utils.CheckErr(err)

		// connect the database
		db := utils.DbConnect()

		sql := ("INSERT INTO address (street,city,state,country,zip_code) VALUES (\"" + oneUser.Address.Street + "\", \"" + oneUser.Address.City + "\", \"" + oneUser.Address.State + "\", \"" + oneUser.Address.Country + "\", \"" + oneUser.Address.ZipCode + "\");")
		_, er := db.Exec(sql)
		utils.CheckErr(er)

		// get the last inserted id
		sqlQuer := ("SELECT id FROM address ORDER BY id DESC LIMIT 1;")
		row := db.QueryRow(sqlQuer)
		err = row.Scan(&lastInsertIDAddress)
		utils.CheckErr(err)

		// create the sql query
		sqlQuery := ("INSERT INTO user (email, first_name, last_name, password, phone, is_alive, id_role, id_address) VALUES ('" + oneUser.Email + "', '" + oneUser.Firstname + "', '" + oneUser.Lastname + "', '" + oneUser.Password + "', '" + oneUser.Phone + "', " + strconv.FormatBool(oneUser.IsAlive) + ", (SELECT id FROM role WHERE role = '" + oneUser.Role + "'), " + strconv.Itoa(lastInsertIDAddress) + ");")

		// execute the sql query
		_, err = db.Exec(sqlQuery)
		utils.CheckErr(err)

		// get the last inserted id
		sqlQuery = ("SELECT id FROM user ORDER BY id DESC LIMIT 1;")
		row = db.QueryRow(sqlQuery)
		err = row.Scan(&lastInsertID)
		utils.CheckErr(err)

		// close the database connection
		db.Close()

		// set the response
		response = structures.InsertOneUser{
			Id:   lastInsertID,
			Type: "success",
			Data: structures.CreateOneUser{
				Email:     oneUser.Email,
				Firstname: oneUser.Firstname,
				Lastname:  oneUser.Lastname,
				Password:  oneUser.Password,
				Phone:     oneUser.Phone,
				IsAlive:   oneUser.IsAlive,
				Role:      oneUser.Role,
				Address: structures.Address{
					Street:  oneUser.Address.Street,
					City:    oneUser.Address.City,
					State:   oneUser.Address.State,
					Country: oneUser.Address.Country,
					ZipCode: oneUser.Address.ZipCode,
				},
			},
			Message: "New user inserted into DB.",
		}

		// send the response
		json.NewEncoder(w).Encode(response)
	}
}
