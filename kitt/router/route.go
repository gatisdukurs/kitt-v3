package router

import "net/http"

type RouteHandler func(ctx RouteCtx)

type Route interface {
	GET(handler RouteHandler)
	Pattern() string
	Execute(ctx RouteCtx)
}

type route struct {
	method  string
	pattern string
	handler RouteHandler
}

func (r *route) GET(handler RouteHandler) {
	r.method = http.MethodGet
	r.handler = handler
}

func (r route) Pattern() string {
	return r.method + " " + r.pattern
}

func (r route) Execute(ctx RouteCtx) {
	r.handler(ctx)
}

func NewRoute(pattern string) Route {
	return &route{
		method:  http.MethodGet,
		pattern: pattern,
	}
}
