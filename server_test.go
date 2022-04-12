package go_web

import (
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	// (1) Create server
	server := http.Server{
		// (2) Create address host with localhost and port with 8080
		Addr: "localhost:8080",
	}

	// (3) Run server
	err := server.ListenAndServe()
	// (4) If error
	if err != nil {
		// Return panic
		panic(err)
	}

}
