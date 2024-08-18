package serviceconfig

import (
	"testing"
)

func TestInitConfig(t *testing.T) {
	var error = ReadConfiguration("/home/pilou/goprojects/restwaf/application.properties")
	if error != nil {
		t.Fatalf("Error %v", error)
	}
	error = Validate()
	if error != nil {
		t.Fatalf("Error %v", error)
	}
}
