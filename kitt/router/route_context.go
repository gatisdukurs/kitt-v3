package router

import "strconv"

type RouteCtx interface {
	Request() Request
	Response() Response
	WithResponse(response Response) RouteCtx
	WithRequest(request Request) RouteCtx

	Param(key string) string
	Params() map[string]string
	ParamInt(key string) int
	ParamInt64(key string) int64
	WithParams(params map[string]string) RouteCtx
}

type routeCtx struct {
	request  Request
	response Response
	params   map[string]string
}

func (rc routeCtx) Request() Request {
	return rc.request
}

func (rc routeCtx) Response() Response {
	return rc.response
}

func (rc routeCtx) Param(key string) string {
	if param, ok := rc.params[key]; ok {
		return param
	}

	return ""
}

func (rc routeCtx) ParamInt(key string) int {
	param, err := strconv.Atoi(rc.Param(key))

	if err != nil {
		return 0
	}

	return param
}

func (rc routeCtx) Params() map[string]string {
	return rc.params
}

func (rc routeCtx) ParamInt64(key string) int64 {
	param, err := strconv.ParseInt(rc.Param(key), 10, 64)

	if err != nil {
		return 0
	}

	return param
}

func (rc *routeCtx) WithResponse(response Response) RouteCtx {
	rc.response = response
	return rc
}

func (rc *routeCtx) WithRequest(request Request) RouteCtx {
	rc.request = request
	return rc
}

func (rc *routeCtx) WithParams(params map[string]string) RouteCtx {
	rc.params = params
	return rc
}

func NewRouteCtx() RouteCtx {
	return &routeCtx{}
}
