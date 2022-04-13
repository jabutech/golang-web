package go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func FormPost(writer http.ResponseWriter, request *http.Request) {
	// (1) Get data body
	firstName := request.PostFormValue("first_name")
	lastName := request.PostFormValue("last_name")

	// (2) Print
	fmt.Fprintf(writer, "Hello %s %s", firstName, lastName)
}

func TestFormPost(t *testing.T) {
	// (1) Create new reader for body request and save to variable requestBody
	requestBody := strings.NewReader("first_name=Rizky&last_name=Darmawan")

	// (2) Create new request with (method, url/host, body)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080", requestBody)

	// (3) Add header
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// (4) Create recorder for write parameter
	recorder := httptest.NewRecorder()

	// (5) Call function FormPost
	FormPost(recorder, request)

	// (6) Get response result
	response := recorder.Result()
	body, err := io.ReadAll(response.Body)

	// (7) If error
	if err != nil {
		panic(err)
	}

	// (8) Print body
	fmt.Println(string(body))
}
