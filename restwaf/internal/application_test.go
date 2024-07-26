package internal

import (
	"testing"
)

func TestReadOpenApiFile(t *testing.T) {
	TestInitConfig(t)
	var application, err = CreateRestWaf()

	if err != nil {
		t.Fatalf(" %v, want ", err)
	}
	err = application.readOpenApiFile()
	if err != nil {
		t.Fatalf(" %v, want ", err)
	}

	if application.Document == nil {
		t.Fatalf("Document is nil")
	}
	application.Start()

}
