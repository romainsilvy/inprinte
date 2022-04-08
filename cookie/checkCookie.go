package cookie

import (
	"fmt"
	"net/http"
)

func CheckCookie(r *http.Request) bool {
	// get the cookie
	cookie, err := r.Cookie("session")
	if err != nil {
		return false
	}
	fmt.Println(cookie)
	return true
}
