package validator

import "restwaf/internal/model"

type Validator interface {
	processRequest(request *model.Request)
}
type ActionValidatorResponse int

const (
	Drop     ActionValidatorResponse = 0
	Deny                             = 1
	Redirect                         = 2
)

type ValidatorResponse struct {
	RuleID int

	// drop, deny, redirect
	Action ActionValidatorResponse

	// Force this status code
	Status int
}
