package router

import "net/http"

type HttpHandler interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}
