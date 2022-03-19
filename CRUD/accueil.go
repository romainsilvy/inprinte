package CRUD

import (
	databaseTools "inprinte/backend/database"
	structures "inprinte/backend/structures"
	utils "inprinte/backend/utils"
	
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func GetAccueil(w http.ResponseWriter, r *http.Request) { 
	db := databaseTools.DbConnect()

	//get accueil infos\\
	//photo, nom, prix, description
	rows, err := db.Query("SELECT name, price, description FROM product LIMIT 3")

	// check errors
	utils.CheckErr(err)

	// var response []JsonResponse
	var accueil []structures.Accueil

	// Foreach product
	for rows.Next() {
		var name string
		var price string
		var description string
		//var picture string
		err = rows.Scan(&name, &price, &description)

		// check errors
		utils.CheckErr(err)

		accueil = append(accueil, structures.Accueil{
			Name: name,
			Price: price,
			Description: description,
		})
	}

	var response = structures.JsonResponseAccueil{
		Type: "success", 
		Data: accueil,
	}

    json.NewEncoder(w).Encode(response)
}

