package validator

import "restwaf/internal/model"

type Validator interface {
	processRequest(request *model.Request)
}
type ActionValidatorResponse string

const (
	Drop     = "drop"
	Deny     = "deny"
	Redirect = "redirect"
)

type OpenApiValidatorResponse struct {
	Message string
}
type WafValidatorResponse struct {
	RuleID int

	// Force this status code
	Status int
}
type ValidatorResponse struct {
	// drop, deny, redirect
	Action                   ActionValidatorResponse
	OpenApiValidatorResponse *OpenApiValidatorResponse
	WafValidatorResponse     *WafValidatorResponse
}
