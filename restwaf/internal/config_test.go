package internal

import (
	"testing"
)

func TestInitConfig(t *testing.T) {
	var error = InitConfig("/home/pilou/goprojects/restwaf/application.properties")
	if error != nil {
		t.Fatalf("Error %v", error)
	}
}
