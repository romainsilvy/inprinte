package authentication

import (
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func generateToken(w http.ResponseWriter, userId int) {
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
}
