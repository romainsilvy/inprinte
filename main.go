package main

import (
	paths "inprinteBackoffice/paths"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	//create a new mux router
	router := mux.NewRouter()

	//handle all the paths
	paths.HandleUsers(router)
	paths.HandleCommand(router)
	paths.HandleCommandLine(router)
	paths.HandleRate(router)
	paths.HandleProducts(router)
	paths.HandleCategory(router)
	paths.HandleRoles(router)
	paths.HandleAuth(router)
	paths.HandleFetch(router)

	http.ListenAndServe(":8080", router)
}
