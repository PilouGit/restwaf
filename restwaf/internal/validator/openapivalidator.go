package validator

import (
	"log"
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

func (validator *OpenApiValidator) CreateOpenApiValidator(body []byte) error {

	document, docErrs := libopenapi.NewDocument(body)
	if docErrs != nil {
		return docErrs
	}
	validator.Document = &document

	return nil

}

func (validator *OpenApiValidator) extractPathAndMethod() {
	document := *(validator.Document)
	v3model, errors := document.BuildV3Model()
	if errors != nil {
		log.Printf("Errors in creating model, %v", errors)

	}
	for pathPairs := v3model.Model.Paths.PathItems.First(); pathPairs != nil; pathPairs = pathPairs.Next() {
		pathItem := pathPairs.Key()
		log.Printf("Path Item, %v", pathItem)

	}
}
