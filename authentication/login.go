package authentication

import (
	"database/sql"
	"encoding/json"
	authentication "inprinte/backend/structures"
	"inprinte/backend/utils"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("my_secret_key")

func Login(w http.ResponseWriter, r *http.Request) {
	utils.SetCorsHeaders(&w)

	var creds authentication.Credentials
	// Get the JSON body and decode into credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		// If the structure of the body is wrong, return an HTTP error
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//open the database connection
	db := utils.DbConnect()

	//global vars
	var id int

	//check if the user exists in the database and if the password is correct, if so return true
	sqlQuerry := "SELECT id FROM user WHERE email = \"" + creds.Email + "\" AND password = \"" + utils.HashPassword(creds.Password) + "\";"
	err = db.QueryRow(sqlQuerry).Scan(&id)
	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusUnauthorized)
		return
	} else if err != nil {
		log.Fatal(err)
	}

	// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	expirationTime := time.Now().AddDate(0, 1, 0)
	// Create the JWT claims, which includes the username and expiry time
	claims := &authentication.Claims{
		Id_user: id,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Finally, we set the client cookie for "token" as the JWT we just generated
	// we also set an expiry time which is the same as the token itself
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}
