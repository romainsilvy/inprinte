package cookie

import (
	"encoding/json"
	"fmt"
	"inprinteBackoffice/structures"
	"inprinteBackoffice/utils"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Create the JWT key used to create the signature
var jwtKey = []byte("zebouloux")

// Create the Signin handler
func Signin(w http.ResponseWriter, r *http.Request) {
	// set cors for cookie
	utils.SetCorsHeaders(&w)

	var auth structures.GetAuthData
	err := json.NewDecoder(r.Body).Decode(&auth)
	utils.CheckErr(err)
	if r.Method == "POST" {

		if GetCredentials(auth.Email, auth.Password) {
			fmt.Println("Authentication success")
			tokenJson := jwt.New(jwt.SigningMethodHS256)
			claimsJson := tokenJson.Claims.(jwt.MapClaims)
			claimsJson["name"] = auth.Email
			claimsJson["admin"] = true
			claimsJson["exp"] = time.Now().Add(time.Hour * 72).Unix()
			tokenStringJson, _ := tokenJson.SignedString(jwtKey)
			tokJson := structures.Token{
				Token: tokenStringJson,
			}
			json.NewEncoder(w).Encode(tokJson)
		} else {
			fmt.Println("Authentication failed")
			w.WriteHeader(http.StatusUnauthorized)
		}
	}
}
