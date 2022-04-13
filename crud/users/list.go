package crud

import (
	"encoding/json"
	"errors"
	"fmt"
	structures "inprinteBackoffice/structures"
	utils "inprinteBackoffice/utils"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("zebouloux")

func GetUsers(w http.ResponseWriter, r *http.Request) {
	//create cors header
	utils.SetCorsHeaders(&w)

	if r.Method == "GET" {

		// get header values
		token := r.Header.Get("Authorization")
		fmt.Println(token)
		fmt.Println(" ")

		// split token
		tokenSplit := strings.Split(token, "\"")
		// for i := 0; i < len(tokenSplit); i++ {
		// 	fmt.Println(tokenSplit[i])
		// }
		// decode token
		tokenString := tokenSplit[3]
		fmt.Println(tokenString)
		// fmt.Println(" ")

		tokenDecoded, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return jwtKey, nil
		})
		if err != nil {
			fmt.Println(err)
		}

		// validate the essential claims
		if !tokenDecoded.Valid {
			fmt.Println("invalid token")
		}
		// fmt.Println(tokenDecoded.Claims)
		// fmt.Println(" ")

		// get claims
		claims, ok := tokenDecoded.Claims.(jwt.MapClaims)
		if ok && tokenDecoded.Valid {
			fmt.Println("name : ", claims["name"])
			fmt.Println("admin : ", claims["admin"])
			fmt.Println("exp : ", claims["exp"])
		}

		//global vars
		var users []structures.GetUsers

		//connect the database
		db := utils.DbConnect()

		//get filters values and update the sqlQuery
		orderBy, rangeBy := utils.GetAllParams(r, "user")
		sqlQuery := "SELECT user.id, email, password, first_name AS firstname, last_name AS lastname, phone, is_alive, street, city, state, country, zip_code, role.role  FROM user INNER JOIN address ON user.id_address = address.id INNER JOIN role ON user.id_role = role.id " + orderBy + " " + rangeBy

		//execute the sql query and check errors
		rows, err := db.Query(sqlQuery)
		utils.CheckErr(err)

		//parse the query
		for rows.Next() {
			//global vars
			var firstname, lastname, email, password, phone, street, city, state, country, zip_code, role string
			var is_alive bool
			var id int

			//retrieve the values and check errors
			err = rows.Scan(&id, &email, &password, &firstname, &lastname, &phone, &is_alive, &street, &city, &state, &country, &zip_code, &role)
			utils.CheckErr(err)

			//add the values to the response
			users = append(users, structures.GetUsers{
				Id:        id,
				Firstname: firstname,
				Lastname:  lastname,
				Email:     email,
				Phone:     phone,
				IsAlive:   is_alive,
				Role:      role,
				Address: structures.Address{
					Street:  street,
					City:    city,
					State:   state,
					Country: country,
					ZipCode: zip_code,
				},
			})
		}
		//close the rows
		rows.Close()

		//close the database connection
		db.Close()

		//create the json response
		utils.SetXTotalCountHeader(&w, len(users))
		json.NewEncoder(w).Encode(users)
	}
}
