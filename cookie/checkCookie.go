package cookie

import (
	"net/http"
)

func CheckCookie(r *http.Request) bool {
	return true
}
