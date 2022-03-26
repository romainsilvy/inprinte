package crud

import (
	utils "inprinteBackoffice/utils"
	"log"
	"net/http"
)

func Insert(w http.ResponseWriter, r *http.Request) {
	utils.SetCorsHeaders(&w)
	log.Println("Inserting user")
}
