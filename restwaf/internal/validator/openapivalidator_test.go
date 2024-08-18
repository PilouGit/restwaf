package validator

import (
	"fmt"
	"io"
	"net/http"
	"testing"
)

func TestOpenApiValidator(t *testing.T) {

	var url string = "http://localhost:8282/v3/api-docs"
	response, error := http.Get(url)
	if error != nil {
		fmt.Println(error)
	}
	// read response body
	body, error := io.ReadAll(response.Body)
	if error != nil {
		t.Fatalf("request is nil %v", error)
	}
	// close response body
	response.Body.Close()
	fmt.Println(string(body))
	openapivalidator, error := CreateOpenApiValidator(body)
	if error != nil {
		t.Fatalf("request is nil %v", error)
	}
	openapivalidator.extractPathAndMethod()
}
