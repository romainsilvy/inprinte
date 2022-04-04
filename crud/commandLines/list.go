package crud

import (
	"encoding/json"
	structures "inprinteBackoffice/structures"
	utils "inprinteBackoffice/utils"

	"net/http"
)

func GetCommandLines(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		//create cors header
		utils.SetCorsHeaders(&w)

		//global vars
		var commandLines []structures.GetCommandLines

		//connect the database
		db := utils.DbConnect()

		//get filters values and update the sqlQuery
		orderBy, rangeBy := utils.GetAllParams(r, "commandLine")
		sqlQuery := "SELECT command_line.id, user.first_name, user.last_name, product.id AS id_product, product.name, command.id AS id_command, user.id AS id_user, quantity, (command_line.quantity * product.price) AS price, command.date, command_line.state FROM command_line INNER JOIN command ON command_line.id_command = command.id INNER JOIN user ON user.id = command.id_user INNER JOIN product ON command_line.id_product = product.id " + orderBy + " " + rangeBy

		//execute the sql query and check errors
		rows, err := db.Query(sqlQuery)
		utils.CheckErr(err)

		//parse the query
		for rows.Next() {
			//global vars
			var firstname, lastname, state, name string
			var id, id_product, id_command, id_user, quantity, price, date int

			//retrieve the values and check errors
			err = rows.Scan(&id, &firstname, &lastname, &id_product, &name, &id_command, &id_user, &quantity, &price, &date, &state)
			utils.CheckErr(err)

			//add the values to the response
			commandLines = append(commandLines, structures.GetCommandLines{
				Id:         id,
				Firstname:  firstname,
				Lastname:   lastname,
				Id_product: id_product,
				Id_command: id_command,
				Id_user:    id_user,
				Quantity:   quantity,
				Price:      price,
				Date:       date,
				State:      state,
				Name:       name,
			})
		}
		//close the rows
		rows.Close()

		//close the database connection
		db.Close()

		//create the json response
		utils.SetXTotalCountHeader(&w, len(commandLines))
		json.NewEncoder(w).Encode(commandLines)
	}
}
