package go_web

import (
	"fmt"
	"net/http"
)

// Hanlder for set cookie
func SetCookie(writer http.ResponseWriter, request *http.Request) {
	// (1) Create new pointer cookie
	cookie := new(http.Cookie)
	// (2) Cookie name
	cookie.Name = "X-Jabutech-Name"
	// (3) Cookie value
	cookie.Value = request.URL.Query().Get("name")
	// (4) Cookie berlaku di semua path
	cookie.Path = "/"

	//  (5) Set cookie to writer
	http.SetCookie(writer, cookie)
	// (6) Print info
	fmt.Fprint(writer, "Success create cookie")
}

// Handler for get cookie
func GetCookie(writer http.ResponseWriter, request *http.Request) {
	// (1) Get cookie
	cookie, err := request.Cookie("X-Jabutech-Name")
	// (2) If Error
	if err != nil {
		// Yes, print info
		fmt.Fprint(writer, "No Cookie")
	} else {
		// No, get cookie value
		name := cookie.Value
		fmt.Fprintf(writer, "Hello %s", name)
	}
}
