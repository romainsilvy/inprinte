package cookie

import (
	"errors"
	"inprinteBackoffice/structures"
	"inprinteBackoffice/utils"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Create the JWT key used to create the signature
var jwtKey = []byte("zebouloux")

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

func decodeToken(w http.ResponseWriter, r *http.Request) jwt.MapClaims {
	//create cors header
	utils.SetCorsHeaders(&w)

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
	utils.CheckErr(err)

	// get claims
	claims, ok := tokenDecoded.Claims.(jwt.MapClaims)
	if !ok {
		return nil
	} else {
		return claims
	}
}

func tokenExist(w http.ResponseWriter, r *http.Request) bool {
	return (len(r.Header.Values("Authorization")) > 0)
}

func validToken(w http.ResponseWriter, r *http.Request) (bool, jwt.MapClaims) {
	//create cors header
	utils.SetCorsHeaders(&w)

	if (r.Header.Values("Authorization")[0]) != "" {
		return true, decodeToken(w, r)
	} else {
		return false, nil
	}
}

func isAdminToken(token jwt.MapClaims) bool {
	if token["admin"] == true {
		return true
	} else {
		return false
	}
}

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
