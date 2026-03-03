package router

import (
	"net/http"
	"net/url"
	"strings"
)

type Request interface {
	Path() string
	WithHttpRequest(request *http.Request) Request
	FormValues() url.Values
	HttpRequest() *http.Request
	HTMX() bool
}

type request struct {
	request *http.Request
}

func (r request) HttpRequest() *http.Request {
	return r.request
}

func (r *request) WithHttpRequest(request *http.Request) Request {
	r.request = request
	return r
}

func (r *request) FormValues() url.Values {
	r.request.ParseForm()
	return r.request.Form
}

func (r *request) HTMX() bool {
	return r.request.Header.Get("HX-Request") == "true"
}

func (r *request) Path() string {
	return strings.TrimSuffix(r.request.URL.Path, "/")
}

func NewRequest() Request {
	return &request{}
}
