package validator

import (
	"log"
	"restwaf/internal/cache"
	"restwaf/internal/model"

	"github.com/pb33f/libopenapi"
	validator "github.com/pb33f/libopenapi-validator"
)

type OpenApiValidator struct {
	requestCache      *cache.RequestCache
	Document          *libopenapi.Document
	RequestValidator  *validator.Validator
	ResponseValidator *validator.Validator
}

func (openApiValidator *OpenApiValidator) CreateOpenApiValidator(body []byte) error {

	document, docErrs := libopenapi.NewDocument(body)
	if docErrs != nil {
		return docErrs
	}
	openApiValidator.Document = &document
	validator, validatorErrs := validator.NewValidator(document)
	if len(validatorErrs) > 0 {
		return validatorErrs[0]
	}
	openApiValidator.RequestValidator = &validator

	return nil

}
func (openApiValidator *OpenApiValidator) ProcessRequest(request *model.Request) *ValidatorResponse {
	if openApiValidator.RequestValidator != nil {
		requestValidator := *openApiValidator.RequestValidator
		httpRequest, _ := request.ToHttpRequest()
		requestValid, validationErrors := requestValidator.ValidateHttpRequest(httpRequest)
		if !requestValid {
			validationError := validationErrors[0]
			errorMessage := validationError.Error()
			openApiValidatorResponse := OpenApiValidatorResponse{Message: errorMessage}
			return &ValidatorResponse{Action: Deny, OpenApiValidatorResponse: &openApiValidatorResponse, WafValidatorResponse: nil}
		}
	}
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
