package CRUD

import (
	"encoding/json"
	databaseTools "inprinte/backend/database"
	"inprinte/backend/structures"
	"inprinte/backend/utils"
	"net/http"
)

func GetAccueil(w http.ResponseWriter, r *http.Request) {
	db := databaseTools.DbConnect()

	//get accueil infos\\
	//photo, nom, prix, descriptions
	rows, err := db.Query("SELECT COUNT(command_line.id) AS nbrOrder, product.id, name, price, description, picture.url FROM product INNER JOIN command_line ON product.id = command_line.id_product INNER JOIN product_picture ON product_picture.id = product.id INNER JOIN picture ON picture.id = product_picture.id WHERE pending_validation = false AND product.is_alive = true GROUP BY command_line.id_product ORDER BY nbrOrder DESC LIMIT 3")

	// check errors
	utils.CheckErr(err)

	// var response []JsonResponse
	var accueil []structures.Accueil

	// Foreach product
	for rows.Next() {
		var name, description, picture string
		var nbrOrder, id, price int
		//var picture string
		err = rows.Scan(&nbrOrder, &id, &name, &price, &description, &picture)

		// check errors
		utils.CheckErr(err)

		accueil = append(accueil, structures.Accueil{
			Name:        name,
			Price:       price,
			Description: description,
			Picture:     picture,
			Id_product:  id,
		})
	}

	var response = structures.JsonResponseAccueil{
		Type: "success",
		Data: accueil,
	}

	json.NewEncoder(w).Encode(response)
}
