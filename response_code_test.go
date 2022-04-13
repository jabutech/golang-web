package go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ResponseCode(writer http.ResponseWriter, request *http.Request) {
	// (1) Get query with name "name"
	name := request.URL.Query().Get("name")

	// (2) If name is empty
	if name == "" {
		// Yes
		// (3) write response bad request
		writer.WriteHeader(http.StatusBadRequest)
		// (4) Write info
		fmt.Fprintf(writer, "Name is empty")
	} else {
		// No
		// (5) Return print name
		fmt.Fprintf(writer, "Hello %s", name)
	}
}

// Test Fail
func TestResponseCodeInvalid(t *testing.T) {
	// (1) Create new request
	request := httptest.NewRequest("GET", "http://localhost:8080", nil)
	// (2) Create new recorder for writer
	recorder := httptest.NewRecorder()

	// (3) Call function ResponseCode
	ResponseCode(recorder, request)

	// (4) Get response
	response := recorder.Result()
	body, err := io.ReadAll(response.Body)

	// (5) If error
	if err != nil {
		panic(err)
	}

	// (6) Print status code from function
	fmt.Println(response.StatusCode)
	// (7) Print status code
	fmt.Println(response.Status)
	// (8) Print body
	fmt.Print(string(body))
}

// Test sucess
func TestResponseCodeValid(t *testing.T) {
	// (1) Create new request
	request := httptest.NewRequest("GET", "http://localhost:8080/?name=Rizky", nil)
	// (2) Create new recorder for writer
	recorder := httptest.NewRecorder()

	// (3) Call function ResponseCode
	ResponseCode(recorder, request)

	// (4) Get response
	response := recorder.Result()
	body, err := io.ReadAll(response.Body)

	// (5) If error
	if err != nil {
		panic(err)
	}

	// (6) Print status code from function
	fmt.Println(response.StatusCode)
	// (7) Print status code
	fmt.Println(response.Status)
	// (8) Print body
	fmt.Print(string(body))
}
