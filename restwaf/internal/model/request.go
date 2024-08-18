package model

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/negasus/haproxy-spoe-go/message"
)

const methodCte = "method"
const appCte = "app"
const idCte = "id"
const bodyCte = "body"
const headersCte = "headers"
const pathCte = "path"
const queryCte = "query"

type Request struct {
	msg     *message.Message
	app     string
	Id      string
	srcIp   net.IP
	srcPort int
	dstIp   net.IP
	dstPort int
	method  string
	path    string
	query   string
	version string
	url     string
	headers string
	body    []byte
}

func CreateRequest(message *message.Message) *Request {
	request := new(Request)
	request.msg = message
	return request
}
func (request *Request) ToHttpRequesst() (*http.Request, error) {
	result, error := http.NewRequest(request.method, request.url, bytes.NewBuffer(request.body))
	headers := strings.Split(request.headers, "\r\n")

	for i := 0; i < len(headers); i++ {
		header := headers[i]
		if strings.Contains(header, ":") {
			couple := strings.Split(header, ":")
			result.Header.Set(strings.TrimSpace(couple[0]), strings.TrimSpace(couple[1]))
		}
	}
	return result, error

}
func (request *Request) Init() error {
	app, found := request.msg.KV.Get(appCte)
	if !found {
		return errors.New("  not found" + appCte)
	} else {
		log.Printf(" app %v", app)
		request.app = fmt.Sprint(app)
	}
	method, found := request.msg.KV.Get(methodCte)
	if !found {
		return errors.New("  not found" + methodCte)
	} else {
		log.Printf(" method %v", method)
		request.method = fmt.Sprint(method)
	}
	url, found := request.msg.KV.Get("full_url")
	if !found {
		return errors.New("  not found url")
	} else {
		log.Printf(" url %v", url)
		if url != nil {
			request.url = fmt.Sprint(url)
		}
	}
	query, found := request.msg.KV.Get(queryCte)
	if found {
		log.Printf(" query %v", query)
		request.query = fmt.Sprint(query)
	}
	id, found := request.msg.KV.Get(idCte)
	if !found {
		return errors.New("  not found" + idCte)
	} else {
		log.Printf(" id %v", id)
		request.Id = fmt.Sprint(id)
	}
	body, found := request.msg.KV.Get(bodyCte)
	if !found {
		return errors.New("  not found" + bodyCte)
	} else {
		request.body = body.([]byte)
		log.Printf(" body  %v", string(request.body))

	}
	headers, found := request.msg.KV.Get(headersCte)
	if !found {
		return errors.New("  not found" + headersCte)
	} else {
		if headers != nil {
			request.headers = headers.(string)
			log.Printf(" headers  %v", string(request.headers))
		}

	}
	path, found := request.msg.KV.Get(pathCte)
	if !found {
		return errors.New("  not found" + pathCte)
	} else {
		request.path = path.(string)
		log.Printf(" headers  %v", string(request.path))

	}
	return nil

}
