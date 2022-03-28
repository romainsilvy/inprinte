package main

import (
	accueil "inprinte/backend/crud/accueil"
	boutique "inprinte/backend/crud/boutique"
	produit "inprinte/backend/crud/produit"
	user "inprinte/backend/crud/user"

	utils "inprinte/backend/utils"

	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func handleAccueil(router *mux.Router) {
	router.HandleFunc("/", accueil.Get).Methods("GET")
}

func handleBoutique(router *mux.Router) {
	router.HandleFunc("/boutique", boutique.Get).Methods("GET")
}

func handleProduit(router *mux.Router) {
	router.HandleFunc("/produit/{id_product}", produit.Get).Methods("GET")
	// router.HandleFunc("/produit", produit.Insert).Methods("INSERT")
	// router.HandleFunc("/produit", produit.Delete).Methods("DELETE")
}

func handleUser(router *mux.Router) {
	router.HandleFunc("/user/{id_user}", user.Get).Methods("GET")
	// router.HandleFunc("/user", user.Update).Methods("UPDATE")
	// router.HandleFunc("/user", user.Delete).Methods("DELETE")
}

func handlePanier(router *mux.Router) {
	// router.HandleFunc("/panier", panier.Get).Methods("GET")
	// router.HandleFunc("/panier", panier.Insert).Methods("INSERT")
}

func handleDemandeDesign(router *mux.Router) {
	// router.HandleFunc("/demande-design", demandeDesign.Get).Methods("GET")
	// router.HandleFunc("/demande-design", demandeDesign.Insert).Methods("INSERT")
	// router.HandleFunc("/demande-design", demandeDesign.Delete).Methods("DELETE")
}

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
	handleDemandeDesign(router)

	http.ListenAndServe(":8080", router)
}

// READ paths
// r.HandleFunc("/", CRUDget.GetAccueil).Methods("GET")
// r.HandleFunc("/boutique", CRUDget.GetBoutique).Methods("GET")
// r.HandleFunc("/produit/{id_product}", CRUDget.GetOneProduct).Methods("GET")
// r.HandleFunc("/user/{id_user}", CRUDget.GetUser).Methods("GET")
// r.HandleFunc("/boutique/{category}", CRUDget.GetBoutiqueByCategory).Methods("GET")

// DELETE paths
// r.HandleFunc("/delete/favorite", CRUDdelete.DeleteFavorite).Methods("DELETE")

// POST paths
// r.HandleFunc("/insert/favorite", CRUDinsert.InsertIntoFavorite).Methods("post")
// backoffice paths
// r.HandleFunc("/users", CRUDget.GetUsers).Methods("GET", "OPTIONS")
