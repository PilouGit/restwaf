package engine

import (
	"testing"
)

func TestInitConfig(t *testing.T) {
	var engine = new(Engine)
	var error = engine.CreateFromConfigurationFile("/home/pilou/goprojects/restwaf/application.properties")
	if error != nil {
		t.Fatalf("Error %v", error)
	}
	error = engine.Init()
	if error != nil {
		t.Fatalf("Error %v", error)
	}

}
