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

type ValidatorResponse struct {
	RuleID int

	// drop, deny, redirect
	Action ActionValidatorResponse

	// Force this status code
	Status int
}
