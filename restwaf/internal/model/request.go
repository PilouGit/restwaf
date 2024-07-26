package model

import (
	"errors"
	"fmt"
	"log"
	"net"

	"github.com/negasus/haproxy-spoe-go/message"
)

const app = "app"
const id = "id"
const body = "body"
const headers = "headers"

type Request struct {
	msg     *message.Message
	app     string
	id      string
	srcIp   net.IP
	srcPort int
	dstIp   net.IP
	dstPort int
	method  string
	path    string
	query   string
	version string
	headers string
	body    []byte
}

func CreateRequest(message *message.Message) *Request {
	request := new(Request)
	request.msg = message
	return request
}

func (request *Request) Init() error {
	method, found := request.msg.KV.Get(app)
	if !found {
		return errors.New("  not found" + app)
	} else {
		log.Printf(" app %v", method)
		request.app = fmt.Sprint(method)
	}
	method, found = request.msg.KV.Get(id)
	if !found {
		return errors.New("  not found" + id)
	} else {
		log.Printf(" id %v", method)
		request.id = fmt.Sprint(method)
	}
	method, found = request.msg.KV.Get(body)
	if !found {
		return errors.New("  not found" + id)
	} else {
		request.body = method.([]byte)
		log.Printf(" body  %v", string(request.body))

	}
	method, found = request.msg.KV.Get(headers)
	if !found {
		return errors.New("  not found" + id)
	} else {
		request.headers = method.(string)
		log.Printf(" body  %v", string(request.headers))

	}
	return nil

}
