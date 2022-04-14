package go_web

import (
	"net/http"
	"testing"
)

// Handler ServeFile
func ServeFile(writer http.ResponseWriter, request *http.Request) {
	// (1) If query parameter "name" is not empty
	if request.URL.Query().Get("name") != "" {
		// Yes, Print page ok
		http.ServeFile(writer, request, "./resources/ok.html")
	} else {
		// No, print page not found
		http.ServeFile(writer, request, "./resources/notfound.html")
	}
}

// Function test serve file
func TestServeFileServer(t *testing.T) {
	// (1) Create server config
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(ServeFile),
	}

	// (2) Run Server
	err := server.ListenAndServe()
	// (3) If error
	if err != nil {
		panic(err)
	}
}
