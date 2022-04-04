package authentication

import (
	"encoding/json"
	"fmt"
	"inprinte/backend/utils"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func checkUserCredentials(email string, password string) (bool, int) {
	//open the database connection
	db := utils.DbConnect()

	//global vars
	var id int

	//check if the user exists in the database and if the password is correct, if so return true
	sql := "SELECT id FROM user WHERE email = \"" + email + "\" AND password = \"" + password + "\";"
	err := db.QueryRow(sql).Scan(&id)
	utils.CheckErr(err)
	if err == nil {
		return true, id
	} else {
		return false, id
	}
}

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	//decode the json response
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	fmt.Println(user)

	email := user.Email
	password := user.Password
	userExists, userId := checkUserCredentials(email, password)
	if userExists {
		//generate a new token based on the email
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["id"] = userId
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		//sign the token with our secret
		tokenString, err := token.SignedString([]byte("Zebouloux"))
		if err != nil {
			log.Fatal(err)
		}

		//send the token to the client
		w.Write([]byte(tokenString))
	} else {
		//redirect to the login page with an error message
		http.Redirect(w, r, "/login?error=true", http.StatusFound)
	}
}
