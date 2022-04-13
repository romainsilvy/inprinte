package authentication

import (
	"encoding/json"
	"inprinte/backend/structures"
	"inprinte/backend/utils"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("my_secret_key")

func getCredentials(r *http.Request) (string, string) {
	var creds structures.Credentials
	// Get the JSON body and decode into credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	utils.CheckErr(err)
	return creds.Email, creds.Password
}

func checkCredentials(r *http.Request, w http.ResponseWriter, email, password string) (bool, int) {
	var id int
	var hashedPassword string

	//open the database connection
	db := utils.DbConnect()

	//check if the user exists in the database and if the password is correct, if so return true
	sqlQuerry := "SELECT id, password FROM user WHERE email = \"" + email + "\";"
	err := db.QueryRow(sqlQuerry).Scan(&id, &hashedPassword)
	if err != nil {
		log.Fatal(err)
		return false, 0
	}
	if utils.CheckPassword(password, hashedPassword) {
		return true, id
	} else {
		return false, 0
	}

}

func createToken(id_user string) structures.Token {
	tokenJson := jwt.New(jwt.SigningMethodHS256)
	claimsJson := tokenJson.Claims.(jwt.MapClaims)
	claimsJson["id_user"] = id_user
	claimsJson["exp"] = time.Now().Add(time.Hour * 72).Unix()

	tokenStringJson, _ := tokenJson.SignedString(jwtKey)
	tokJson := structures.Token{
		Auth: tokenStringJson,
	}

	return tokJson
}

func Login(w http.ResponseWriter, r *http.Request) {
	utils.SetCorsHeaders(&w)

	if r.Method == "POST" {
		email, password := getCredentials(r)
		ok, id := checkCredentials(r, w, email, password)
		if ok {
			tokJson := createToken(strconv.Itoa(id))
			json.NewEncoder(w).Encode(tokJson)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
	}
}
