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
	WithSendable(sendable RouteResponseSendable) RouteResponse
	WithStringResponse(response string) RouteResponse
	HTMX() string
}

type routeResponse struct {
	status   int
	sendable RouteResponseSendable
	response string
}

func (rr *routeResponse) WithStatus(status int) RouteResponse {
	rr.status = status
	return rr
}

func (rr *routeResponse) WithSendable(sendable RouteResponseSendable) RouteResponse {
	rr.sendable = sendable
	return rr
}

func (rr *routeResponse) WithStringResponse(response string) RouteResponse {
	rr.response = response
	return rr
}

func (rr routeResponse) Status() int {
	return rr.status
}

func (rr routeResponse) Body() string {
	if rr.response != "" {
		return rr.response
	}

	if rr.sendable != nil {
		return rr.sendable.Render()
	}

	return ""
}

func (rr routeResponse) HTMX() string {
	if rr.response != "" {
		return rr.response
	}

	if rr.sendable != nil {
		return rr.sendable.HTMX()
	}

	return ""
}

func NewRouteResponse() RouteResponse {
	return &routeResponse{
		status: http.StatusOK,
	}
}
