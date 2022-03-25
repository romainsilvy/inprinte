package main

import (
	"encoding/json"
	CRUDdelete "inprinte/backend/CRUD/delete"
	CRUDget "inprinte/backend/CRUD/get"
	CRUDinsert "inprinte/backend/CRUD/insert"

	utils "inprinte/backend/utils"

	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	// databaseTools.Faker()
	utils.InprinteAscii()
	r := mux.NewRouter()

	// backoffice paths
	r.HandleFunc("/users", CRUDget.GetUsers).Methods("GET", "OPTIONS")

	// READ paths
	r.HandleFunc("/", CRUDget.GetAccueil).Methods("GET")
	r.HandleFunc("/boutique", CRUDget.GetBoutique).Methods("GET")
	r.HandleFunc("/produit/{id_product}", CRUDget.GetOneProduct).Methods("GET")
	// r.HandleFunc("/user/{id_user}", CRUDget.GetUser).Methods("GET")
	r.HandleFunc("/boutique/{category}", CRUDget.GetBoutiqueByCategory).Methods("GET")

	// DELETE paths
	r.HandleFunc("/delete/favorite", CRUDdelete.DeleteFavorite).Methods("DELETE")

	// POST paths
	r.HandleFunc("/insert/favorite", CRUDinsert.InsertIntoFavorite).Methods("post")

	r.HandleFunc("/people", GetPeopleAPI).Methods("GET", "OPTIONS")

	http.ListenAndServe(":8080", r)
}

func GetPeopleAPI(w http.ResponseWriter, r *http.Request) {

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	// return "OKOK"
	json.NewEncoder(w).Encode("OKOK")
}
