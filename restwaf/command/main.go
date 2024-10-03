package main

import (
	"log"
	"restwaf/internal"
)

func main() {

	var application = new(internal.Application)
	var error = application.InitFrom("/home/pilou/goprojects/restwaf/application.json")
	if error != nil {
		log.Fatalf("Error %v", error)
	}

	application.Start()
}
