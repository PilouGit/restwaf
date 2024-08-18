package cache

import (
	"restwaf/internal/model"
	"testing"
)

func createMockRequest() *model.Request {
	result := new(model.Request)
	result.Id = "titi"
	return result
}
func TestGache(t *testing.T) {

	cache := CreateRequestGCache(10)
	cache.putRequest(createMockRequest())
	request := cache.getRequest("titi")
	if request == nil {
		t.Fatalf("request is nil")
	}
}
