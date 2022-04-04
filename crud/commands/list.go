package crud

import (
	"encoding/json"
	"inprinteBackoffice/structures"
	"inprinteBackoffice/utils"
	"net/http"
)

func GetCommands(w http.ResponseWriter, r *http.Request) {
	//create cors header
	utils.SetCorsHeaders(&w)

	//global vars
	var commands []structures.GetCommands

	//connect the database
	db := utils.DbConnect()

	//get filters values and update the sqlQuery
	orderBy, rangeBy := utils.GetAllParams(r, "command")
	sqlQuery := "SELECT command.id, user.id AS id_user, first_name, last_name, command.date, command.state, SUM(command_line.quantity)AS quantity, SUM(command_line.quantity * product.price) AS price FROM command INNER JOIN user ON command.id_user = user.id INNER JOIN command_line ON command_line.id_command = command.id INNER JOIN product ON product.id = command_line.id_product GROUP BY command.id " + orderBy + rangeBy

	//execute the sql query and check errors
	rows, err := db.Query(sqlQuery)
	utils.CheckErr(err)

	//parse the query
	for rows.Next() {
		//global vars
		var firstname, lastname, date, status string
		var id, id_user, price, quantity int

		//retrieve the values and check errors
		err = rows.Scan(&id, &id_user, &firstname, &lastname, &date, &status, &quantity, &price)
		utils.CheckErr(err)

		//add the values to the response
		commands = append(commands, structures.GetCommands{
			Id:        id,
			Firstname: firstname,
			Lastname:  lastname,
			Id_user:   id_user,
			Price:     price,
			Quantity:  quantity,
			Date:      date,
			Status:    status,
		})
	}
	//close the rows

	//create the json response
	utils.SetXTotalCountHeader(&w, len(commands))
	json.NewEncoder(w).Encode(commands)
}
