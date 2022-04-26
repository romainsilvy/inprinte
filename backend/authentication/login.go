//package authentication provides all the functionality needed to authenticate users
package authentication

import (
	"encoding/json"
	"inprinte/backend/structures"
	"inprinte/backend/utils"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//Login is the function that handles the login request
func Login(w http.ResponseWriter, r *http.Request) {
	//set the cors headers
	utils.SetCorsHeaders(&w)

	if r.Method == "POST" {
		//retrieve email and password
		email, password := getCredentials(r)

		//check the credentials
		ok, id := checkCredentials(r, w, email, password)
		if ok {
			//create the token and send it to the client
			tokJson := createToken(strconv.Itoa(id))
			json.NewEncoder(w).Encode(tokJson)
		} else {
			//set the StatusUnauthorized header
			w.WriteHeader(http.StatusUnauthorized)
		}
	}
}

//the secret key used to encrypt passwords
var jwtKey = []byte(os.Getenv("JWT_KEY"))

//getCredentials retrieve the user credentials passed in the request
func getCredentials(r *http.Request) (string, string) {
	//global vars
	var creds structures.Credentials

	// Get the JSON body and decode into credentials
	err := json.NewDecoder(r.Body).Decode(&creds)

	//check error
	utils.CheckErr(err)

	//return the credentials
	return creds.Email, creds.Password
}

//checkCredentials check if the user credentials are correct
func checkCredentials(r *http.Request, w http.ResponseWriter, email, password string) (bool, int) {
	//global vars
	var id int
	var hashedPassword string

	//open the database connection
	db := utils.DbConnect()

	//check if the user exists in the database and retrieve the hashed password
	sqlQuerry := "SELECT id, password FROM user WHERE email = \"" + email + "\";"
	err := db.QueryRow(sqlQuerry).Scan(&id, &hashedPassword)
	if err != nil {
		log.Fatal(err)
		return false, 0
	}

	//compare the hashed password with the password passed in the request
	if utils.CheckPassword(password, hashedPassword) {
		return true, id
	} else {
		return false, 0
	}

}

//createToken creates a new jwt token and send it to the client
func createToken(id_user string) structures.Token {
	//create the token
	tokenJson := jwt.New(jwt.SigningMethodHS256)
	claimsJson := tokenJson.Claims.(jwt.MapClaims)

	//set the token claims
	claimsJson["id_user"] = id_user
	claimsJson["exp"] = time.Now().Add(time.Hour * 72).Unix()

	//sign the token
	tokenStringJson, _ := tokenJson.SignedString(jwtKey)

	//create the token object
	tokJson := structures.Token{
		Auth: tokenStringJson,
	}

	//return the token
	return tokJson
}
