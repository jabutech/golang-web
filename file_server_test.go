package go_web

import (
	"net/http"
	"testing"
)

func TestFileServer(t *testing.T) {
	// (1) Init directory resources
	directory := http.Dir("./resources")
	// (2) Set File server
	fileServer := http.FileServer(directory)

	// (3) Create new server with mux
	mux := http.NewServeMux()
	// (4) Create endpoint static
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// (5) Server config
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	// (6) Run server
	err := server.ListenAndServe()
	// (7) If error
	if err != nil {
		panic(err)
	}
}
