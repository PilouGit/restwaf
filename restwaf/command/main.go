package main

import (
	"log"
	"restwaf/internal"
)

func main() {

	/*log.Print("listen 3000")

	listener, err := net.Listen("tcp4", "127.0.0.1:3000")
	if err != nil {
		log.Printf("error create listener, %v", err)
		os.Exit(1)
	}
	defer listener.Close()

	a := agent.New(handler, logger.NewDefaultLog())

	if err := a.Serve(listener); err != nil {
		log.Printf("error agent serve: %+v\n", err)
	}*/
	var error = internal.InitConfig("/home/pilou/goprojects/restwaf/application.properties")
	if error != nil {
		log.Panicf("Error %v", error)
	}

	application, error := internal.CreateRestWaf()
	if error != nil {
		log.Panicf("Error %v", error)
	}
	application.Start()
}

/*func handler(req *request.Request) {

	log.Printf("handle request EngineID: '%v'", req)
	for _, s := range *req.Messages {
		log.Printf("%s", s.Name)
	}

}*/
