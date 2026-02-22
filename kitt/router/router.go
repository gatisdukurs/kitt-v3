package router

import "net/http"

type Routes = map[string]Route

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
	r.routes[route.Pattern()] = route
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
	pattern := method + " " + path

	ctx := NewRouteCtx()
	ctx.WithRequest(NewRequest().WithHttpRequest(request))
	ctx.WithResponse(NewResponse().WithHttpResponse(response))

	if route, ok := r.routes[pattern]; ok {
		route.Execute(ctx)
		return
	}

	if r.handler404 != nil {
		r.handler404(ctx)
		return
	}

	http.NotFound(response, request)
}

func NewRouter() Router {
	return &router{
		routes: make(Routes),
	}
}
