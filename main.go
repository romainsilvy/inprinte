package main

import (
	authentication "inprinte/backend/authentication"
	accueil "inprinte/backend/crud/accueil"
	boutique "inprinte/backend/crud/boutique"
	panier "inprinte/backend/crud/panier"
	produit "inprinte/backend/crud/produit"
	user "inprinte/backend/crud/user"
	"log"

	utils "inprinte/backend/utils"

	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

//handleAccueil is the handle function for the accueil page
func handleAccueil(router *mux.Router) {
	router.HandleFunc("/", accueil.Get).Methods("GET")
}

//handleBoutique is the handle function for the boutique page
func handleBoutique(router *mux.Router) {
	router.HandleFunc("/boutique", boutique.Get).Methods("GET")
}

//handleProduit is the handle function for the produit page
func handleProduit(router *mux.Router) {
	router.HandleFunc("/produit/{id_product}", produit.Get).Methods("GET")
}

//handleUser is the handle function for the user page
func handleUser(router *mux.Router) {
	router.HandleFunc("/user/{id_user}", user.Get).Methods("GET")
}

//handlePanier is the handle function for the panier page
func handlePanier(router *mux.Router) {
	router.HandleFunc("/panier", panier.Get).Methods("GET")
}

//handleAuthentication is the handle function for the authentication page
func handleAuthentication(router *mux.Router) {
	router.HandleFunc("/connexion", authentication.Login)
	router.HandleFunc("/inscription", authentication.Signup)
}

//main is the main function
func main() {
	//print the inprinte logo
	utils.InprinteAscii()

	//create a new mux router
	router := mux.NewRouter()

	//handle all the paths
	handleAccueil(router)
	handleBoutique(router)
	handleProduit(router)
	handleUser(router)
	handlePanier(router)
	handleAuthentication(router)

	//start the server
	log.Fatal(http.ListenAndServe(":8080", router))
}
