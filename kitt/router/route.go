package router

import (
	"net/http"
	"strings"
)

type RouteHandler func(ctx RouteCtx) RouteResponse

type Route interface {
	GET(handler RouteHandler) Route
	POST(handler RouteHandler) Route
	DELETE(handler RouteHandler) Route
	Handler(method string, handler RouteHandler) Route
	Pattern() string
	Match(method string, path string) bool
	Params(path string) map[string]string
	Execute(ctx RouteCtx) RouteResponse
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

func (r route) Execute(ctx RouteCtx) RouteResponse {
	return r.handler(ctx)
}

func (r route) Match(method, path string) bool {
	if r.method != method {
		return false
	}

	pattern := normalizePath(r.pattern)
	path = normalizePath(path)

	// wildcard support: "/assets/*"
	if strings.HasSuffix(pattern, "/*") {
		prefix := strings.TrimSuffix(pattern, "/*")
		return path == prefix || strings.HasPrefix(path, prefix+"/")
	}

	patternParts := splitPath(pattern)
	pathParts := splitPath(path)

	if len(patternParts) != len(pathParts) {
		return false
	}

	for i := range patternParts {
		pp := patternParts[i]
		p := pathParts[i]

		if isParamSegment(pp) {
			continue
		}

		if pp != p {
			return false
		}
	}

	return true
}

func (r route) Params(path string) map[string]string {
	params := make(map[string]string)

	pattern := normalizePath(r.pattern)
	path = normalizePath(path)

	// no params for wildcard routes
	if strings.HasSuffix(pattern, "/*") {
		return params
	}

	patternParts := splitPath(pattern)
	pathParts := splitPath(path)

	if len(patternParts) != len(pathParts) {
		return params
	}

	for i := range patternParts {
		pp := patternParts[i]
		p := pathParts[i]

		if isParamSegment(pp) {
			name := strings.TrimPrefix(pp, ":")
			if name != "" {
				params[name] = p
			}
		}
	}

	return params
}

func NewRoute(pattern string) Route {
	return &route{
		method:  http.MethodGet,
		pattern: strings.TrimSuffix(pattern, "/"),
	}
}

func normalizePath(path string) string {
	if path == "" {
		return "/"
	}

	path = strings.TrimSuffix(path, "/")
	if path == "" {
		return "/"
	}

	return path
}

func splitPath(path string) []string {
	path = strings.Trim(path, "/")
	if path == "" {
		return nil
	}

	return strings.Split(path, "/")
}

func isParamSegment(segment string) bool {
	return strings.HasPrefix(segment, ":") && len(segment) > 1
}
