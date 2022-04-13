package cookie

import (
	"encoding/json"
	"inprinteBackoffice/structures"
	"inprinteBackoffice/utils"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/sessions"
)

// Create the JWT key used to create the signature
var jwtKey = []byte("zebouloux")

var store = sessions.NewCookieStore(jwtKey)

// Create the Signin handler
func Signin(w http.ResponseWriter, r *http.Request) {
	// set cors for cookie
	utils.SetCorsHeaders(&w)

	if r.Method == "POST" {
		tokenJson := jwt.New(jwt.SigningMethodHS256)
		claimsJson := tokenJson.Claims.(jwt.MapClaims)
		claimsJson["admin"] = true
		claimsJson["name"] = "admin"
		claimsJson["exp"] = time.Now().Add(time.Hour * 72).Unix()
		tokenStringJson, _ := tokenJson.SignedString(jwtKey)
		tokJson := structures.Token{
			Auth: tokenStringJson,
		}
		json.NewEncoder(w).Encode(tokJson)
	}
}
