package router

import (
	"net/http"
	"strings"
)

type RouteHandler func(ctx RouteCtx)

type Route interface {
	GET(handler RouteHandler) Route
	POST(handler RouteHandler) Route
	DELETE(handler RouteHandler) Route
	Handler(method string, handler RouteHandler) Route
	Pattern() string
	Match(method string, path string) bool
	Execute(ctx RouteCtx)
}

type route struct {
	method  string
	pattern string
	handler RouteHandler
}

func (r *route) GET(handler RouteHandler) Route {
	return r.Handler(http.MethodGet, handler)
}

func (r *route) POST(handler RouteHandler) Route {
	return r.Handler(http.MethodPost, handler)
}

func (r *route) DELETE(handler RouteHandler) Route {
	return r.Handler(http.MethodDelete, handler)
}

func (r *route) Handler(method string, handler RouteHandler) Route {
	r.method = method
	r.handler = handler
	return r
}

func (r route) Pattern() string {
	return r.method + " " + r.pattern
}

func (r route) Execute(ctx RouteCtx) {
	r.handler(ctx)
}

func (r route) Match(method, path string) bool {
	if r.method != method {
		return false
	}

	// wildcard support: "/assets/*"
	if strings.HasSuffix(r.pattern, "/*") {
		prefix := strings.TrimSuffix(r.pattern, "/*")
		return strings.HasPrefix(path, prefix)
	}

	// exact match
	return r.pattern == strings.TrimSuffix(path, "/")
}

func NewRoute(pattern string) Route {
	return &route{
		method:  http.MethodGet,
		pattern: strings.TrimSuffix(pattern, "/"),
	}
}
