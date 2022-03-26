package crud

import (
	"encoding/json"
	"inprinteBackoffice/structures"
	"inprinteBackoffice/utils"
	"net/http"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	//global vars
	var AllCommands []structures.AllCommands

	//connect the database
	db := utils.DbConnect()

	//create cors header
	utils.SetCorsHeaders(&w)

	//get filters values and update the sqlQuery
	orderBy, rangeBy := utils.GetAllParams(r, "command")
	sqlQuery := "SELECT DISTINCT command.id AS id, user.id AS id_user, user.first_name, user.last_name, command.date, command.state, command_line.quantity , product.price FROM user INNER JOIN command ON command.id_user = user.id INNER JOIN command_line ON command_line.id_command = command.id INNER JOIN product ON product.id = command_line.id_product " + orderBy + rangeBy

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
		AllCommands = append(AllCommands, structures.AllCommands{
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
	utils.SetXTotalCountHeader(&w, len(AllCommands))
	json.NewEncoder(w).Encode(AllCommands)
}
