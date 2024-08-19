package config

import (
	"testing"
)

func TestInitConfig(t *testing.T) {
	var configuration = new(Configuration)
	var error = configuration.ReadConfiguration("/home/pilou/goprojects/restwaf/application.properties")
	if error != nil {
		t.Fatalf("Error %v", error)
	}
	error = configuration.Validate()
	if error != nil {
		t.Fatalf("Error %v", error)
	}
	if configuration.GlobalConfiguration.Port != 3000 {
		t.Fatalf("Port is  %v", configuration.GlobalConfiguration.Port)
	}
}
