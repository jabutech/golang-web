package go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func RequestHeader(writer http.ResponseWriter, request *http.Request) {
	// (1) Sample get header `content type`
	contentType := request.Header.Get("content-type")
	// (2) print
	fmt.Fprint(writer, contentType)
}

func TestRequestHeader(t *testing.T) {
	// (1) Create variable request with httptest.NewRequest(method, hosts, body)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/", nil)
	// (2) Added new header "Content-Type"
	request.Header.Add("Content-Type", "application/json")

	// (3) Create var recorder for get response writer
	recorder := httptest.NewRecorder()

	// (4) Call function RequestHeader with parameter recorder for get response writer and request
	RequestHeader(recorder, request)

	// (5) Get response
	response := recorder.Result()
	// (6) Read response
	body, _ := io.ReadAll(response.Body)

	// (7) Print and convert body to string
	fmt.Println(string(body))
}

// Function for Response Header
func ResponseHeader(writer http.ResponseWriter, request *http.Request) {
	// (1) Write new response header
	writer.Header().Add("X-Powered-By", "Jabutech")
	// (2) Print to console
	fmt.Fprint(writer, "OK")
}

func TestResponseHeader(t *testing.T) {
	// (1) Create variable request with httptest.NewRequest(method, hosts, body)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/", nil)
	// (2) Added new header "Content-Type"
	request.Header.Add("Content-Type", "application/json")

	// (3) Create var recorder for get response writer
	recorder := httptest.NewRecorder()

	// (4) Call function RequestHeader with parameter recorder for get response writer and request
	ResponseHeader(recorder, request)

	// (5) Get response
	response := recorder.Result()
	// (6) Read response
	body, _ := io.ReadAll(response.Body)

	// (7) Print and convert body to string
	fmt.Println(string(body))

	// (8) Print and Get header x-powered-by
	fmt.Println(response.Header.Get("x-powered-by"))
}
