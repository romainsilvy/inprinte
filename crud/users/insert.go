package crud

import (
	"log"
	"net/http"
)

func Insert(w http.ResponseWriter, r *http.Request) {
	log.Println("Inserting user")
	log.Println(w)
}
