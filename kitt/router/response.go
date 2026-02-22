package router

import (
	"io"
	"net/http"
)

type Renderable interface {
	Render() string
}

type Response interface {
	WithStatus(status int) Response
	WithResponse(response http.ResponseWriter) Response
	Send(interface{})
}

type response struct {
	status   int
	writer   io.Writer
	response http.ResponseWriter
}

func (r *response) WithStatus(status int) Response {
	r.status = status
	return r
}

func (r *response) WithResponse(response http.ResponseWriter) Response {
	r.response = response
	return r
}

func (r *response) Send(raw interface{}) {
	// switch between types
	switch sendable := raw.(type) {
	case string:
		r.response.WriteHeader(r.status)
		r.response.Write([]byte(sendable))
	default:
		// do nothing
	}
}

func NewResponse() Response {
	return &response{
		status: http.StatusOK,
	}
}
