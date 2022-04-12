package go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// (1) Function handler
func HelloHandler(writer http.ResponseWriter, request *http.Request) {
	// (2) Print for response
	fmt.Fprint(writer, "Hello World.")
}

// (2) Function Test
func TestHttp(t *testing.T) {
	// (1) Create variable request with httptest.NewRequest(method, hosts, body)
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello", nil)
	// (2) Create variable recorder for get response
	recorder := httptest.NewRecorder()

	//(3) Call function HelloHandler for test with send parameter var recoder and request
	HelloHandler(recorder, request)

	// (4) Get response
	response := recorder.Result()

	// (5) Get response body
	body, err := io.ReadAll(response.Body)

	// (6) If error
	if err != nil {
		panic(err)
	}

	// (7) conversi response body to string
	bodyString := string(body)

	// (8) Print response
	fmt.Println(bodyString)

}
