package main

import (
	auth "inprinteBackoffice/authenticate"
	category "inprinteBackoffice/crud/category"
	commandLines "inprinteBackoffice/crud/commandLines"
	commands "inprinteBackoffice/crud/commands"
	products "inprinteBackoffice/crud/products"
	rates "inprinteBackoffice/crud/rates"
	roles "inprinteBackoffice/crud/roles"
	users "inprinteBackoffice/crud/users"
	fetch "inprinteBackoffice/fetch"
	"net/http"

	"github.com/gorilla/mux"
)

func handleUsers(router *mux.Router) {
	router.HandleFunc("/users", users.Insert).Methods("OPTIONS", "POST")
	router.HandleFunc("/users", users.GetUsers).Methods("GET")
	router.HandleFunc("/users/{id_user}", users.GetUser).Methods("GET")
	router.HandleFunc("/users/{id_user}", users.Update).Methods("OPTIONS", "PUT")
	router.HandleFunc("/users/{id_user}", users.Delete).Methods("OPTIONS", "DELETE")
}

func handleCommand(router *mux.Router) {
	router.HandleFunc("/commands", commands.Insert).Methods("OPTIONS", "POST")
	router.HandleFunc("/commands", commands.GetCommands).Methods("GET")
	router.HandleFunc("/commands/{id_command}", commands.GetCommand).Methods("GET")
	router.HandleFunc("/commands/{id_command}", commands.Update).Methods("PUT", "OPTIONS")
	router.HandleFunc("/commands/{id_command}", commands.Delete).Methods("DELETE", "OPTIONS")
}

func handleCommandLine(router *mux.Router) {
	router.HandleFunc("/commandLines", commandLines.Insert).Methods("OPTIONS", "POST")
	router.HandleFunc("/commandLines", commandLines.GetCommandLines).Methods("GET")
	router.HandleFunc("/commandLines/{id_commandLine}", commandLines.GetCommandLine).Methods("GET")
	router.HandleFunc("/commandLines/{id_commandLine}", commandLines.Update).Methods("OPTIONS", "PUT")
	router.HandleFunc("/commandLines/{id_commandLine}", commandLines.Delete).Methods("OPTIONS", "DELETE")
}

func handleRate(router *mux.Router) {
	router.HandleFunc("/rates", rates.Insert).Methods("OPTIONS", "POST")
	router.HandleFunc("/rates", rates.GetRates).Methods("GET")
	router.HandleFunc("/rates/{id_rate}", rates.GetRate).Methods("GET")
	router.HandleFunc("/rates/{id_rate}", rates.Update).Methods("OPTIONS", "PUT")
	router.HandleFunc("/rates/{id_rate}", rates.Delete).Methods("OPTIONS", "DELETE")
}

func handleProducts(router *mux.Router) {
	router.HandleFunc("/products", products.Insert).Methods("OPTIONS", "POST")
	router.HandleFunc("/products", products.GetProducts).Methods("GET")
	router.HandleFunc("/products/{id_product}", products.GetProduct).Methods("GET")
	router.HandleFunc("/products/{id_product}", products.Update).Methods("OPTIONS", "PUT")
	router.HandleFunc("/products/{id_product}", products.Delete).Methods("OPTIONS", "DELETE")
}

func handleCategory(router *mux.Router) {
	router.HandleFunc("/categories", category.Insert).Methods("OPTIONS", "POST")
	router.HandleFunc("/categories", category.GetCategories).Methods("GET")
	router.HandleFunc("/categories/{id_category}", category.GetCategory).Methods("GET")
	router.HandleFunc("/categories/{id_category}", category.Update).Methods("OPTIONS", "PUT")
	router.HandleFunc("/categories/{id_category}", category.Delete).Methods("OPTIONS", "DELETE")
}

func handleRoles(router *mux.Router) {
	router.HandleFunc("/roles", roles.Insert).Methods("OPTIONS", "POST")
	router.HandleFunc("/roles", roles.GetRoles).Methods("GET")
	router.HandleFunc("/roles/{id_role}", roles.GetRole).Methods("GET")
	router.HandleFunc("/roles/{id_role}", roles.Update).Methods("OPTIONS", "PUT")
	router.HandleFunc("/roles/{id_role}", roles.Delete).Methods("OPTIONS", "DELETE")
}

func handleAuth(router *mux.Router) {
	router.HandleFunc("/login", auth.Authenticate).Methods("OPTIONS", "POST")
}

func handleFetch(router *mux.Router) {
	router.HandleFunc("/rolesFetch", fetch.FetchRoles).Methods("GET")
	router.HandleFunc("/categoriesFetch", fetch.FetchCategories).Methods("GET")
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
	handleCategory(router)
	handleRoles(router)
	handleAuth(router)
	handleFetch(router)

	http.ListenAndServe(":8080", router)
}
