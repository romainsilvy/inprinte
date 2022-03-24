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

/* func manageRoutes() {
// 	router := mux.NewRouter()

// 	router.HandleFunc("/users", CRUD.GetUsers).Methods("GET")

// 	router.HandleFunc("/user/{userid}", CreateUser).Methods("POST")
// 	router.HandleFunc("/user/{userid}", CRUD.GetUsers).Methods("GET")
// 	router.HandleFunc("/user/{userid}", UpdateUser).Methods("UPDATE")
// 	router.HandleFunc("/user/{userid}", DeleteUser).Methods("DELETE")

// 	fmt.Println("Server at 8080")
// 	log.Fatal(http.ListenAndServe(":8000", router))
}*/

func main() {
	// databaseTools.Faker()
	utils.InprinteAscii()

	r := mux.NewRouter()

	// backoffice paths

	r.HandleFunc("/users", CRUDget.GetUsers).Methods("GET")

	// normal paths

	r.HandleFunc("/", CRUDget.GetAccueil).Methods("GET")
	r.HandleFunc("/boutique", CRUDget.GetBoutique).Methods("GET")
	r.HandleFunc("/produit/{id}", CRUDget.GetOneProduct).Methods("GET")
	r.HandleFunc("/user/{id_user}", CRUDget.GetUserData).Methods("GET")
  r.HandleFunc("/boutique/{category}", CRUD.GetBoutiqueByCategory).Methods("GET")
  r.HandleFunc("/delete/favorite", CRUDdelete.DeleteFavorite).Methods("DELETE")
	r.HandleFunc("/insert/favorite", CRUDinsert.InsertIntoFavorite).Methods("post")

	http.ListenAndServe(":8080", r)

}
