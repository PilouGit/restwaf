package internal

import (
	"fmt"
	"testing"
)

func TestInitConfig(t *testing.T) {
	InitConfig("/home/pilou/goprojects/restwaf/application.properties")
	fmt.Println(Global)
}
