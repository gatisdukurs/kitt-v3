package router

import "net/http"

type RouteResponseSendable interface {
	Render() string
	HTMX() string
}

type RouteResponse interface {
	Status() int
	Body() string
	WithStatus(status int) RouteResponse
	HTMX() string
}

type routeResponse struct {
	status   int
	sendable RouteResponseSendable
}

func (rr *routeResponse) WithStatus(status int) RouteResponse {
	rr.status = status
	return rr
}

func (rr routeResponse) Status() int {
	return rr.status
}

func (rr routeResponse) Body() string {
	return rr.sendable.Render()
}

func (rr routeResponse) HTMX() string {
	return rr.sendable.HTMX()
}

func NewRouteResponse(sendable RouteResponseSendable) RouteResponse {
	return &routeResponse{
		status:   http.StatusOK,
		sendable: sendable,
	}
}
