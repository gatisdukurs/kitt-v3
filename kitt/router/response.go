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
	WithWriter(writer io.Writer) Response
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

func (r *response) WithWriter(writer io.Writer) Response {
	r.writer = writer
	return r
}

func (r *response) Send(raw interface{}) {
	// switch between types
	switch sendable := raw.(type) {
	case string:
		r.writer.Write([]byte(sendable))
	default:
		// do nothing
	}
}

func NewResponse() Response {
	return &response{
		status: http.StatusOK,
	}
}
