package go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Function SayHello
func SayHello(writer http.ResponseWriter, request *http.Request) {
	// (1) Get query parameter with name query "name"
	name := request.URL.Query().Get("name")

	// (2) If name is empty
	if name == "" {
		// Yes, Response only "Hello"
		fmt.Fprint(writer, "Hello")
	} else {
		//  No, Response Hello with parameter
		fmt.Fprintf(writer, "Hello %s", name)
	}
}

// Function Test Query
func TestQueryParameter(t *testing.T) {
	// (1) Create var request with httptest.NewRequest(Method, "Host with query parameter", body)
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=Rizky", nil)
	// (2) Create variable recoreder for writer
	recorder := httptest.NewRecorder()

	// (3) Call function SayHello with parameter recorder for writer, request for recorder
	SayHello(recorder, request)

	// (4) Get response from recorder
	response := recorder.Result()

	// (5) convert response to string
	body, err := io.ReadAll(response.Body)

	// (6) IF error
	if err != nil {
		panic(err)
	}

	// (7)  IF no error, print response
	fmt.Println(string(body))
}
