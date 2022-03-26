package crud

import (
	"encoding/json"
	structures "inprinteBackoffice/structures"
	utils "inprinteBackoffice/utils"

	"net/http"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	//global vars
	var allRates []structures.AllRates

	//connect the database
	db := utils.DbConnect()

	//create cors header
	utils.SetCorsHeaders(&w)

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
		allRates = append(allRates, structures.AllRates{
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

	//create the json response
	utils.SetXTotalCountHeader(&w, len(allRates))
	json.NewEncoder(w).Encode(allRates)
}
