package cache

import "restwaf/internal/model"

type RequestCache interface {
	putRequest(request *model.Request)
	getRequest(uid string) *model.Request
}
