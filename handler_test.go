package go_web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHandler(t *testing.T) {
	//  (1) Create variable handler with value http.HanlderFunc
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		// (2) Logic web with print hello world
		fmt.Fprint(writer, "Hello World")
	}

	// (3) Create Server
	server := http.Server{
		// (4) Set address with host localhost and port 8080
		Addr: "localhost:8080",
		// (5) Set Handler with use var handller
		Handler: handler,
	}

	// (6) Run Server
	err := server.ListenAndServe()
	// (7) If err
	if err != nil {
		// Return panic
		panic(err)
	}
}

func TestServeMux(t *testing.T) {
	// (1) Create new handler erver with mux
	mux := http.NewServeMux()

	// (2) Create endpoint root
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		// Print Hello World
		fmt.Fprint(writer, "Hello World")
	})

	// (3) Create endpoint hi
	mux.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		// Print hi
		fmt.Fprint(w, "Hi")
	})

	// (3) Create endpoint hi
	mux.HandleFunc("/get-request", func(w http.ResponseWriter, r *http.Request) {
		// Get requst method
		fmt.Fprint(w, r.Method)
		// Get requst URI
		fmt.Fprint(w, r.RequestURI)
	})

	// (4) Create server
	server := http.Server{
		// Set address with host localhost and port 8080
		Addr: "localhost:8080",
		// Set handler with mux
		Handler: mux,
	}

	// (5) Run server
	err := server.ListenAndServe()
	// (6) If error
	if err != nil {
		// Return panic
		panic(err)
	}

}
