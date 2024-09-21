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
	SrcIp   net.IP
	SrcPort int
	DstIp   net.IP
	DstPort int
	Method  string
	path    string
	Query   string
	Version string
	Url     string
	Body    []byte

	Headers     *map[string]string
	httpRequest *http.Request
}

func CreateRequest(message *message.Message) *Request {
	request := new(Request)
	request.msg = message
	return request
}
func (request *Request) ToHttpRequest() (*http.Request, error) {
	if request.httpRequest == nil {
		result, _ := http.NewRequest(request.Method, request.Url, bytes.NewBuffer(request.Body))
		request.httpRequest = result
		for key, value := range *request.Headers {
			result.Header.Set(key, value)
		}

	}

	return request.httpRequest, nil

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
		request.Method = fmt.Sprint(method)
	}
	url, found := request.msg.KV.Get("full_url")
	if !found {
		return errors.New("  not found url")
	} else {
		log.Printf(" url %v", url)
		if url != nil {
			request.Url = fmt.Sprint(url)
		}
	}
	version, found := request.msg.KV.Get("version")
	if !found {
		return errors.New("  not found url")
	} else {
		log.Printf(" version %v", url)
		if url != nil {
			request.Version = fmt.Sprint(version)
		}
	}

	query, found := request.msg.KV.Get(queryCte)
	if found {
		log.Printf(" query %v", query)
		request.Query = fmt.Sprint(query)
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
		request.Body = body.([]byte)
		log.Printf(" body  %v", string(request.Body))

	}
	headers, found := request.msg.KV.Get(headersCte)
	if !found {
		return errors.New("  not found" + headersCte)
	} else {
		if headers != nil {
			request.Headers = new(map[string]string)
			(*request.Headers) = make(map[string]string)
			headerString := headers.(string)
			headersSplited := strings.Split(headerString, "\r\n")

			for i := 0; i < len(headersSplited); i++ {
				header := headersSplited[i]
				if strings.Contains(header, ":") {
					couple := strings.Split(header, ":")
					(*request.Headers)[strings.TrimSpace(couple[0])] = strings.TrimSpace(couple[1])
				}
			}
			log.Printf(" headers  %v", request.Headers)
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
