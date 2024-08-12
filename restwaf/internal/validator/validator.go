package validator

import "restwaf/internal/model"

type Validator interface {
	processRequest(request *model.Request)
}
