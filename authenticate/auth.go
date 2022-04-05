package authenticate

import (
	"encoding/json"
	"fmt"
	"inprinteBackoffice/structures"
	"inprinteBackoffice/utils"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func Authenticate(w http.ResponseWriter, r *http.Request) {
	//set cors headers
	utils.SetCorsHeaders(&w)

	var auth structures.GetAuthData
	var password string
	var is_alive bool
	var id_role int

	//decode json
	err := json.NewDecoder(r.Body).Decode(&auth)
	utils.CheckErr(err)

	//connect the database
	db := utils.DbConnect()

	sqlQuery := ("SELECT password, is_alive, id_role FROM user WHERE email = '" + auth.Email + "'")
	rows, err := db.Query(sqlQuery)
	utils.CheckErr(err)

	//parse the query
	for rows.Next() {
		//retrieve the values and check errors
		err = rows.Scan(&password, &is_alive, &id_role)
		utils.CheckErr(err)
	}

	// close the database connection
	db.Close()

	if password == auth.Password {
		fmt.Println("Password is correct")
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["admin"] = true
		claims["name"] = "admin"
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		tokenString, _ := token.SignedString([]byte("secret"))
		tok := structures.Token{
			Auth: tokenString,
		}
		json.NewEncoder(w).Encode(tok)
	} else {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Invalid credentials"))
	}
}
