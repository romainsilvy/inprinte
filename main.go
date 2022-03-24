package main

import (
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
	r.HandleFunc("/users", CRUDget.GetUsers).Methods("GET")

	// READ paths
	r.HandleFunc("/", CRUDget.GetAccueil).Methods("GET")
	r.HandleFunc("/boutique", CRUDget.GetBoutique).Methods("GET")
	r.HandleFunc("/produit/{id_product}", CRUDget.GetOneProduct).Methods("GET")
	r.HandleFunc("/user/{id_user}", CRUDget.GetUser).Methods("GET")
	r.HandleFunc("/boutique/{category}", CRUDget.GetBoutiqueByCategory).Methods("GET")

	// DELETE paths
	r.HandleFunc("/delete/favorite", CRUDdelete.DeleteFavorite).Methods("DELETE")

	// POST paths
	r.HandleFunc("/insert/favorite", CRUDinsert.InsertIntoFavorite).Methods("post")

	http.ListenAndServe(":8080", r)
}
