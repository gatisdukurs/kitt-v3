package router

import "net/http"

type Request interface {
	Path() string
	WithHttpRequest(request *http.Request) Request
}

type request struct {
	request *http.Request
}

func (r *request) WithHttpRequest(request *http.Request) Request {
	r.request = request
	return r
}

func (r *request) Path() string {
	return r.request.URL.Path
}

func NewRequest() Request {
	return &request{}
}
