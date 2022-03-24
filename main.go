package main

import (
	CRUD "inprinte/backend/CRUD/get"
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
	r.HandleFunc("/users", CRUD.GetUsers).Methods("GET")

	// normal paths
	r.HandleFunc("/", CRUD.GetAccueil).Methods("GET")
	r.HandleFunc("/boutique", CRUD.GetBoutique).Methods("GET")
	r.HandleFunc("/produit/{id}", CRUD.GetOneProduct).Methods("GET")
	r.HandleFunc("/user/{id_user}", CRUD.GetUserData).Methods("GET")

	http.ListenAndServe(":8080", r)

}
