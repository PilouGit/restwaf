package internal

import (
	"fmt"
	"log"
	"net"
	"os"
	"restwaf/internal/engine"
	"restwaf/internal/model"
	"restwaf/internal/validator"

	"github.com/negasus/haproxy-spoe-go/action"
	"github.com/negasus/haproxy-spoe-go/agent"
	"github.com/negasus/haproxy-spoe-go/logger"
	"github.com/negasus/haproxy-spoe-go/message"
	"github.com/negasus/haproxy-spoe-go/request"
)

type Application struct {
	engine *engine.Engine
}

func (application *Application) InitFrom(filename string) error {
	application.engine = new(engine.Engine)
	error := application.engine.CreateFromConfigurationFile(filename)
	if error != nil {
		return error
	}
	return application.engine.Init()

}
func (application *Application) Start() {
	configuration := application.engine.GetConfiguration()
	log.Printf("listen %d", configuration.GlobalConfiguration.Port)
	concatenated := fmt.Sprintf("%s:%d", configuration.GlobalConfiguration.Adress, configuration.GlobalConfiguration.Port)

	listener, err := net.Listen("tcp4", concatenated)
	if err != nil {
		log.Printf("error create listener, %v", err)
		os.Exit(1)
	}
	defer listener.Close()

	a := agent.New(application.handler, logger.NewDefaultLog())

	if err := a.Serve(listener); err != nil {
		log.Printf("error agent serve: %+v\n", err)
	}
}
func (application *Application) handler(req *request.Request) {
	log.Printf("handle request EngineID: '%v'", req)
	messsage, error := req.Messages.GetByName("coraza-req")
	if error != nil {
		log.Printf("var method  not found in message")
		return
	}
	if messsage != nil {
		validadorresponse := application.processRequest(messsage)
		if validadorresponse != nil {
			req.Actions.SetVar(action.ScopeTransaction, "action", "deny")

		}
	}
}

func (application *Application) processRequest(message *message.Message) *validator.ValidatorResponse {

	request := model.CreateRequest(message)
	error := request.Init()
	if error != nil {
		log.Printf("error agent serve: %+v\n", error)
	}
	return application.engine.ProcessRequest(request)
}
