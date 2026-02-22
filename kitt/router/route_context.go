package router

type RouteCtx interface {
	Request() Request
	Response() Response
	WithResponse(response Response) RouteCtx
	WithRequest(request Request) RouteCtx
}

type routeCtx struct {
	request  Request
	response Response
}

func (rc routeCtx) Request() Request {
	return rc.request
}

func (rc routeCtx) Response() Response {
	return rc.response
}

func (rc *routeCtx) WithResponse(response Response) RouteCtx {
	rc.response = response
	return rc
}

func (rc *routeCtx) WithRequest(request Request) RouteCtx {
	rc.request = request
	return rc
}

func NewRouteCtx() RouteCtx {
	return &routeCtx{}
}
