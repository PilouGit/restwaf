package validator

import (
	"restwaf/internal/cache"

	"github.com/pb33f/libopenapi"
	validator "github.com/pb33f/libopenapi-validator"
)

type OpenApiValidator struct {
	requestCache      *cache.RequestCache
	Document          *libopenapi.Document
	RequestValidator  *validator.Validator
	ResponseValidator *validator.Validator
}

func CreateOpenApiValidator(body []byte) (*OpenApiValidator, error) {
	openApiValidator := new(OpenApiValidator)

	document, docErrs := libopenapi.NewDocument(body)
	if docErrs != nil {
		return nil, docErrs
	}
	openApiValidator.Document = &document

	return openApiValidator, nil

}
