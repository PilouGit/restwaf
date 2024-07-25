package internal

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/negasus/haproxy-spoe-go/agent"
	"github.com/negasus/haproxy-spoe-go/logger"
	"github.com/negasus/haproxy-spoe-go/request"
	"github.com/pb33f/libopenapi"
	validator "github.com/pb33f/libopenapi-validator"
)

type Application struct {
	Document  *libopenapi.Document
	Validator *validator.Validator
}

func (application *Application) readOpenApiFile() error {

	var url string = Global.OpenApi.Url
	response, error := http.Get(url)
	if error != nil {
		fmt.Println(error)
	}
	// read response body
	body, error := io.ReadAll(response.Body)
	if error != nil {
		return error
	}
	// close response body
	response.Body.Close()
	fmt.Println(string(body))
	document, docErrs := libopenapi.NewDocument(body)

	if docErrs != nil {
		return docErrs
	}
	application.Document = &document

	validator, validatorErrs := validator.NewValidator(document)
	if len(validatorErrs) > 0 {
		panic("document is bad")
	}
	application.Validator = &validator
	return nil
}

func (application *Application) Start() {
	log.Print("listen 3000")

	listener, err := net.Listen("tcp4", "127.0.0.1:3000")
	if err != nil {
		log.Printf("error create listener, %v", err)
		os.Exit(1)
	}
	defer listener.Close()

	a := agent.New(handler, logger.NewDefaultLog())

	if err := a.Serve(listener); err != nil {
		log.Printf("error agent serve: %+v\n", err)
	}
}
func handler(req *request.Request) {
}

func New() (*Application, error) {
	var application = new(Application)
	var error = application.readOpenApiFile()

	if error != nil {
		return nil, error

	}

	return application, nil
}
