package paths

import (
	cookie "inprinteBackoffice/cookie"
	category "inprinteBackoffice/crud/category"
	commandLines "inprinteBackoffice/crud/commandLines"
	commands "inprinteBackoffice/crud/commands"
	products "inprinteBackoffice/crud/products"
	rates "inprinteBackoffice/crud/rates"
	roles "inprinteBackoffice/crud/roles"
	users "inprinteBackoffice/crud/users"
	fetch "inprinteBackoffice/fetch"

	"github.com/gorilla/mux"
)

func HandleUsers(router *mux.Router) {
	router.HandleFunc("/api/users", users.Insert).Methods("OPTIONS", "POST")
	router.HandleFunc("/api/users", users.GetUsers).Methods("GET")
	router.HandleFunc("/api/users/{id_user}", users.GetUser).Methods("GET")
	router.HandleFunc("/api/users/{id_user}", users.Update).Methods("OPTIONS", "PUT")
	router.HandleFunc("/api/users/{id_user}", users.Delete).Methods("OPTIONS", "DELETE")
}

func HandleCommand(router *mux.Router) {
	router.HandleFunc("/api/commands", commands.Insert).Methods("OPTIONS", "POST")
	router.HandleFunc("/api/commands", commands.GetCommands).Methods("GET")
	router.HandleFunc("/api/commands/{id_command}", commands.GetCommand).Methods("GET")
	router.HandleFunc("/api/commands/{id_command}", commands.Update).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/commands/{id_command}", commands.Delete).Methods("DELETE", "OPTIONS")
}

func HandleCommandLine(router *mux.Router) {
	router.HandleFunc("/api/commandLines", commandLines.Insert).Methods("OPTIONS", "POST")
	router.HandleFunc("/api/commandLines", commandLines.GetCommandLines).Methods("GET")
	router.HandleFunc("/api/commandLines/{id_commandLine}", commandLines.GetCommandLine).Methods("GET")
	router.HandleFunc("/api/commandLines/{id_commandLine}", commandLines.Update).Methods("OPTIONS", "PUT")
	router.HandleFunc("/api/commandLines/{id_commandLine}", commandLines.Delete).Methods("OPTIONS", "DELETE")
}

func HandleRate(router *mux.Router) {
	router.HandleFunc("/api/rates", rates.Insert).Methods("OPTIONS", "POST")
	router.HandleFunc("/api/rates", rates.GetRates).Methods("GET")
	router.HandleFunc("/api/rates/{id_rate}", rates.GetRate).Methods("GET")
	router.HandleFunc("/api/rates/{id_rate}", rates.Update).Methods("OPTIONS", "PUT")
	router.HandleFunc("/api/rates/{id_rate}", rates.Delete).Methods("OPTIONS", "DELETE")
}

func HandleProducts(router *mux.Router) {
	router.HandleFunc("/api/products", products.Insert).Methods("OPTIONS", "POST")
	router.HandleFunc("/api/products", products.GetProducts).Methods("GET")
	router.HandleFunc("/api/products/{id_product}", products.GetProduct).Methods("GET")
	router.HandleFunc("/api/products/{id_product}", products.Update).Methods("OPTIONS", "PUT")
	router.HandleFunc("/api/products/{id_product}", products.Delete).Methods("OPTIONS", "DELETE")
}

func HandleCategory(router *mux.Router) {
	router.HandleFunc("/api/categories", category.Insert).Methods("OPTIONS", "POST")
	router.HandleFunc("/api/categories", category.GetCategories).Methods("GET")
	router.HandleFunc("/api/categories/{id_category}", category.GetCategory).Methods("GET")
	router.HandleFunc("/api/categories/{id_category}", category.Update).Methods("OPTIONS", "PUT")
	router.HandleFunc("/api/categories/{id_category}", category.Delete).Methods("OPTIONS", "DELETE")
}

func HandleRoles(router *mux.Router) {
	router.HandleFunc("/api/roles", roles.Insert).Methods("OPTIONS", "POST")
	router.HandleFunc("/api/roles", roles.GetRoles).Methods("GET")
	router.HandleFunc("/api/roles/{id_role}", roles.GetRole).Methods("GET")
	router.HandleFunc("/api/roles/{id_role}", roles.Update).Methods("OPTIONS", "PUT")
	router.HandleFunc("/api/roles/{id_role}", roles.Delete).Methods("OPTIONS", "DELETE")
}

func HandleAuth(router *mux.Router) {
	router.HandleFunc("/api/login", cookie.Authenticate).Methods("OPTIONS", "POST")
}

func HandleFetch(router *mux.Router) {
	router.HandleFunc("/api/rolesFetch", fetch.FetchRoles).Methods("GET")
	router.HandleFunc("/api/categoriesFetch", fetch.FetchCategories).Methods("GET")
}
