package cookie

import (
	"encoding/json"
	"inprinteBackoffice/structures"
	"inprinteBackoffice/utils"
	"net/http"
)

// Create the Signin handler
func Signin(w http.ResponseWriter, r *http.Request) {
	// set cors for cookie
	utils.SetCorsHeaders(&w)

	// get authenticate values
	var auth structures.GetAuthData
	err := json.NewDecoder(r.Body).Decode(&auth)
	utils.CheckErr(err)

	if r.Method == "POST" {
		if CheckLogin(auth.Email, auth.Password) {
			token := CreateToken(auth.Email)
			json.NewEncoder(w).Encode(token)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
	}
}
