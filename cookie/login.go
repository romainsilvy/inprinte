package cookie

import (
	"encoding/json"
	"fmt"
	"inprinteBackoffice/structures"
	"inprinteBackoffice/utils"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/sessions"
)

// // Create the JWT key used to create the signature
// var jwtKey = []byte("zebouloux")

// // Create a struct to read the username and password from the request body
// type Credentials struct {
// 	Password string `json:"password"`
// 	Email    string `json:"email"`
// }

// // Create a struct that will be encoded to a JWT.
// type Claims struct {
// 	Id int `json:"id"`
// 	jwt.StandardClaims
// }

var (
	key   = []byte("ismatheplatypus@w*")
	store = sessions.NewCookieStore(key)
)

//this structure is used to unmarshall the value of id_th
type MyBody struct {
	Id_th string `json:id_th`
	Value string `json:value`
}

// Create the Signin handler
func Signin(w http.ResponseWriter, r *http.Request) {
	// set cors header
	utils.SetCorsHeaders(&w)

	if r.Method == "POST" {

		session, _ := store.Get(r, "bourses")
		session.Values["authenticated"] = true
		session.Values["user"] = "TEST"
		session.Save(r, w)
		fmt.Println("utilisateur connecté")

		// var creds Credentials
		// // Get the JSON body and decode into credentials
		// err := json.NewDecoder(r.Body).Decode(&creds)
		// if err != nil {
		// 	// If the structure of the body is wrong, return an HTTP error
		// 	w.WriteHeader(http.StatusBadRequest)
		// 	return
		// }

		// // If a password exists for the given user
		// // AND, if it is the same as the password we received, the we can move ahead
		// // if NOT, then we return an "Unauthorized" status
		// hasGoodPassword, id_user := GetCredentials(creds.Email, creds.Password)
		// fmt.Println(hasGoodPassword)
		// fmt.Println(id_user)
		// if !hasGoodPassword {
		// 	w.WriteHeader(http.StatusUnauthorized)
		// 	return
		// }

		// // Declare the expiration time of the token
		// // here, we have kept it as 5 minutes
		// expirationTime := time.Now().AddDate(1, 0, 0)
		// // Create the JWT claims, which includes the username and expiry time
		// claims := &Claims{
		// 	Id: id_user,
		// 	StandardClaims: jwt.StandardClaims{
		// 		// In JWT, the expiry time is expressed as unix milliseconds
		// 		ExpiresAt: expirationTime.Unix(),
		// 	},
		// }

		// // Declare the token with the algorithm used for signing, and the claims
		// token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		// fmt.Println(token)

		// // Create the JWT string
		// tokenString, err := token.SignedString(jwtKey)
		// if err != nil {
		// 	// If there is an error in creating the JWT return an internal server error
		// 	w.WriteHeader(http.StatusInternalServerError)
		// 	return
		// }

		// // Finally, we set the client cookie for "token" as the JWT we just generated
		// // we also set an expiry time which is the same as the token itself
		// http.SetCookie(w, &http.Cookie{
		// 	Name:    "token",
		// 	Value:   tokenString,
		// 	Expires: expirationTime,
		// })

		tokenJson := jwt.New(jwt.SigningMethodHS256)
		claimsJson := tokenJson.Claims.(jwt.MapClaims)
		claimsJson["admin"] = true
		claimsJson["name"] = "admin"
		claimsJson["exp"] = time.Now().Add(time.Hour * 72).Unix()

		tokenStringJson, _ := tokenJson.SignedString([]byte("secret"))
		tokJson := structures.Token{
			Auth: tokenStringJson,
		}
		json.NewEncoder(w).Encode(tokJson)
	}
}

// c, err := r.Cookie("token")
// if err != nil {
// 	if err == http.ErrNoCookie {
// 		w.WriteHeader(http.StatusUnauthorized)
// 		return
// 	}
// 	w.WriteHeader(http.StatusBadRequest)
// 	return
// }

// tknStr := c.Value

// tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
// 	return jwtKey, nil
// })
// if err != nil {
// 	if err == jwt.ErrSignatureInvalid {
// 		w.WriteHeader(http.StatusUnauthorized)
// 		return
// 	}
// 	w.WriteHeader(http.StatusBadRequest)
// 	return
// }
// if !tkn.Valid {
// 	w.WriteHeader(http.StatusUnauthorized)
// 	return
// }
