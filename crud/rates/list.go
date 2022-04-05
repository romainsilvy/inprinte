package crud

import (
	"encoding/json"
	"inprinteBackoffice/structures"
	utils "inprinteBackoffice/utils"

	"net/http"
)

func GetRates(w http.ResponseWriter, r *http.Request) {
	//create cors header
	utils.SetCorsHeaders(&w)

	if r.Method == "GET" {
		//global vars
		var rates []structures.GetRates

		//connect the database
		db := utils.DbConnect()

		//get filters values and update the sqlQuery
		orderBy, rangeBy := utils.GetAllParams(r, "rate")

		sqlQuery := "select rate.id, product.id, product.name, user.id AS id_user, user.first_name, user.last_name, rate.stars_number FROM rate INNER JOIN product ON product.id = rate.id_product INNER JOIN user ON user.id = rate.id_user " + orderBy + " " + rangeBy

		//execute the sql query and check errors
		rows, err := db.Query(sqlQuery)
		utils.CheckErr(err)

		//parse the query
		for rows.Next() {
			//global vars
			var firstname, lastname, name string
			var id, id_product, id_user, stars_number int

			//retrieve the values and check errors
			err = rows.Scan(&id, &id_product, &name, &id_user, &firstname, &lastname, &stars_number)
			utils.CheckErr(err)

			//add the values to the response
			rates = append(rates, structures.GetRates{
				Id:           id,
				Id_product:   id_product,
				Name:         name,
				Id_user:      id_user,
				Firstname:    firstname,
				Lastname:     lastname,
				Stars_number: stars_number,
			})
		}
		//close the rows
		rows.Close()

		//close the database connection
		db.Close()

		//create the json response
		utils.SetXTotalCountHeader(&w, len(rates))
		json.NewEncoder(w).Encode(rates)
	}
}
