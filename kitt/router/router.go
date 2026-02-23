package router

import (
	"net/http"
	"strings"
)

type Routes = []Route

type Router interface {
	HttpHandler
	To(route Route) Router
	With404(handler RouteHandler) Router
	Routes() Routes
}

type router struct {
	routes     Routes
	handler404 RouteHandler
}

func (r *router) To(route Route) Router {
	r.routes = append(r.routes, route)
	return r
}

func (r *router) With404(handler RouteHandler) Router {
	r.handler404 = handler
	return r
}

func (r router) Routes() Routes {
	return r.routes
}

func (r router) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	method := request.Method
	path := request.URL.Path

	ctx := NewRouteCtx()
	ctx.WithRequest(NewRequest().WithHttpRequest(request))
	ctx.WithResponse(NewResponse().WithHttpResponse(response))

	for _, route := range r.routes {
		if route.Match(method, path) {
			route.Execute(ctx)
			return
		}
	}

	if r.handler404 != nil {
		r.handler404(ctx)
		return
	}

	http.NotFound(response, request)
}

func NewRouter() Router {
	return &router{}
}

func NewStaticRoute(prefix, dir string) Route {
	if !strings.HasSuffix(prefix, "/") {
		prefix += "/"
	}

	fs := http.StripPrefix(prefix, http.FileServer(http.Dir(dir)))

	return NewRoute(prefix + "*").GET(func(ctx RouteCtx) {
		fs.ServeHTTP(
			ctx.Response().HttpResponse(),
			ctx.Request().HttpRequest(),
		)
	})
}
