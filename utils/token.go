package utils

import (
	"errors"
	"inprinteBackoffice/structures"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Create the JWT key used to create the signature
var jwtKey = []byte("zebouloux")

// Create the Signin handler
func CreateToken(email string) structures.Token {
	// Create the token
	tokenJson := jwt.New(jwt.SigningMethodHS256)
	claimsJson := tokenJson.Claims.(jwt.MapClaims)
	claimsJson["name"] = email
	claimsJson["admin"] = true
	claimsJson["exp"] = time.Now().Add(time.Hour * 72).Unix()
	tokenStringJson, _ := tokenJson.SignedString(jwtKey)
	tokJson := structures.Token{
		Token: tokenStringJson,
	}
	return tokJson
}

// Decode the token
func decodeToken(w http.ResponseWriter, r *http.Request) jwt.MapClaims {
	//create cors header
	SetCorsHeaders(&w)

	// get header values
	token := r.Header.Get("Authorization")

	// split token
	tokenSplit := strings.Split(token, "\"")
	tokenString := tokenSplit[3]

	// decode token
	tokenDecoded, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtKey, nil
	})
	CheckErr(err)

	// get claims
	claims, ok := tokenDecoded.Claims.(jwt.MapClaims)
	if !ok {
		return nil
	} else {
		return claims
	}
}

// Check if the token exist
func tokenExist(w http.ResponseWriter, r *http.Request) bool {
	return (len(r.Header.Values("Authorization")) > 0)
}

// Check if the token is valid
func validToken(w http.ResponseWriter, r *http.Request) (bool, jwt.MapClaims) {
	//create cors header
	SetCorsHeaders(&w)

	if (r.Header.Values("Authorization")[0]) != "" {
		return true, decodeToken(w, r)
	} else {
		return false, nil
	}
}

// Check if the token is admin
func isAdminToken(token jwt.MapClaims) bool {
	if token["admin"] == true {
		return true
	} else {
		return false
	}
}

// Check if the user is logged and admin
func Securized(w http.ResponseWriter, r *http.Request) bool {
	if tokenExist(w, r) {
		ok, claims := validToken(w, r)
		if isAdminToken(claims) && ok {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}
