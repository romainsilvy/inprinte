package cookie

import (
	"fmt"
	"net/http"
)

func CheckCookie(r *http.Request) bool {
	// get the cookie admin
	fmt.Println(r.Cookies())
	for _, cookie := range r.Cookies() {
		fmt.Println("Found a cookie named:", cookie.Name)
	}
	return true
}
