package cookie

import (
	"encoding/json"
	"inprinteBackoffice/structures"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func Login(w http.ResponseWriter, passwordDB, passwordEntered string, is_alive bool, id_role int) {
	if passwordDB == passwordEntered && is_alive && id_role == 1 {
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["admin"] = true
		claims["name"] = "admin"
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
		tokenString, _ := token.SignedString([]byte("zebouloux"))
		tok := structures.Token{
			Auth: tokenString,
		}
		json.NewEncoder(w).Encode(tok)
	} else {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Invalid credentials"))
	}
}
