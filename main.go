package main

import (
	commandLines "inprinteBackoffice/crud/commandLines"
	commands "inprinteBackoffice/crud/commands"
	products "inprinteBackoffice/crud/products"
	rates "inprinteBackoffice/crud/rates"
	users "inprinteBackoffice/crud/users"

	"net/http"

	"github.com/gorilla/mux"
)

func handleUsers(router *mux.Router) {
	router.HandleFunc("/users", users.Insert).Methods("POST", "OPTIONS")
	router.HandleFunc("/users", users.GetAll).Methods("GET")
	router.HandleFunc("/users/{id_user}", users.GetOne).Methods("GET")
	// router.HandleFunc("/users/{id_user}", users.update).Methods("UPDATE")
	// router.HandleFunc("/users/{id_user}", users.delete).Methods("DELETE")
}

func handleCommand(router *mux.Router) {
	router.HandleFunc("/commands", commands.GetAll).Methods("GET")
	// router.HandleFunc("/command/{id_command}", command.getOne).Methods("GET")
	// router.HandleFunc("/command/{id_command}", command.update).Methods("UPDATE")
}

func handleCommandLine(router *mux.Router) {
	router.HandleFunc("/commandLines", commandLines.GetAll).Methods("GET")
	// router.HandleFunc("/command-line/{id_commandLine}", commandLine.getOne).Methods("GET")
	// router.HandleFunc("/command-line/{id_commandLine}", commandLine.update).Methods("UPDATE")
}

func handleRate(router *mux.Router) {
	router.HandleFunc("/rates", rates.GetAll).Methods("GET")
	// router.HandleFunc("/rate/{id_rate}", rate.getOne).Methods("GET")
	// router.HandleFunc("/rate/{id_rate}", rate.update).Methods("UPDATE")
}

func handleProducts(router *mux.Router) {
	// router.HandleFunc("/products", products.insert).Methods("POST")
	router.HandleFunc("/products", products.GetAll).Methods("GET")
	// router.HandleFunc("/products/{id_product}", products.getOne).Methods("GET")
	// router.HandleFunc("/products/{id_product}", products.update).Methods("UPDATE")
	// router.HandleFunc("/products/{id_product}", products.delete).Methods("DELETE")
}

func handleCategory(router *mux.Router) {
	// router.HandleFunc("/category", category.insert).Methods("POST")
	// router.HandleFunc("/category", category.getAll).Methods("GET")
	// router.HandleFunc("/category/{id_product}", category.getOne).Methods("GET")
	// router.HandleFunc("/category/{id_product}", category.update).Methods("UPDATE")
	// router.HandleFunc("/category/{id_product}", category.delete).Methods("DELETE")
}

func handleRole(router *mux.Router) {
	// router.HandleFunc("/role", role.insert).Methods("POST")
	// router.HandleFunc("/role", role.getAll).Methods("GET")
	// router.HandleFunc("/role/{id_product}", role.getOne).Methods("GET")
	// router.HandleFunc("/role/{id_product}", role.update).Methods("UPDATE")
	// router.HandleFunc("/role/{id_product}", role.delete).Methods("DELETE")
}

func main() {
	//create a new mux router
	router := mux.NewRouter()

	//handle all the paths
	handleUsers(router)
	handleCommand(router)
	handleCommandLine(router)
	handleRate(router)
	handleProducts(router)

	http.ListenAndServe(":8080", router)
}
