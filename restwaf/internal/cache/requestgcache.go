package cache

import (
	"restwaf/internal/model"

	"github.com/bluele/gcache"
)

type RequestGCache struct {
	Gc *gcache.Cache
}

func CreateRequestGCache(nbTransation uint) *RequestGCache {
	result := new(RequestGCache)
	gcache := gcache.New(int(nbTransation)).LRU().
		Build()
	result.Gc = &gcache
	return result
}

func (cache *RequestGCache) putRequest(request *model.Request) {
	gcache := cache.Gc
	(*gcache).Set((*request).Id, request)
}
func (cache *RequestGCache) getRequest(uid string) *model.Request {
	gcache := cache.Gc
	value, error := (*gcache).Get(uid)
	if error == nil {
		return value.(*model.Request)
	} else {
		return nil
	}
}
