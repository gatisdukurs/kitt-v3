package kitt

import (
	"fmt"
	"net/http"
)

const EventRouterOnRequest = "router.onRequest"

type RouteResponseSendable interface {
	Render() string
}

type RouteResponse struct {
	w      http.ResponseWriter
	sent   bool
	status int
}

func (r *RouteResponse) Send(s RouteResponseSendable, statusCode int) {
	if r.sent {
		return
	}
	r.w.WriteHeader(statusCode)
	r.w.Write([]byte(s.Render()))
	r.sent = true
	r.status = statusCode
}

func (r RouteResponse) Status() int {
	return r.status
}

func (r RouteResponse) Sent() bool {
	return r.sent
}

func (r *RouteRequest) Method() string {
	return r.r.Method
}

func (r *RouteResponse) Redirect(url string) {
	if r.sent {
		return
	}

	statusCode := http.StatusSeeOther // 303 (best for POST -> GET)

	r.w.Header().Set("Location", url)
	r.w.WriteHeader(statusCode)

	r.sent = true
	r.status = statusCode
}

type RouteRequest struct {
	r          *http.Request
	formParsed bool
}

func (r *RouteResponse) Header(key, value string) {
	if r.sent {
		return
	}
	r.w.Header().Set(key, value)
}

func (r *RouteRequest) HTMX() bool {
	// HTMX sends: HX-Request: true
	return r.r.Header.Get("HX-Request") == "true"
}

func (r RouteRequest) Path() string {
	return r.r.URL.Path
}

func (r RouteRequest) Host() string {
	return r.r.Host
}

func (r RouteRequest) Query() string {
	return r.r.URL.RawQuery
}

func (r *RouteRequest) ensureFormParsed() error {
	if r.formParsed {
		return nil
	}

	r.r.Body = http.MaxBytesReader(nil, r.r.Body, 1<<20) // 1MB

	if err := r.r.ParseForm(); err != nil {
		return err
	}

	r.formParsed = true
	return nil
}

func (r RouteRequest) formValue(key string) string {
	_ = r.ensureFormParsed()
	return r.r.FormValue(key)
}

func (r *RouteRequest) Input(key string) string {
	return r.formValue(key)
}

func (r *RouteRequest) Has(key string) bool {
	r.ensureFormParsed()
	_, ok := r.r.Form[key]
	return ok
}

func (r *RouteRequest) Inputs() map[string][]string {
	_ = r.ensureFormParsed()
	return r.r.Form
}

type RouteCtx struct {
	route    Route
	response *RouteResponse
	request  *RouteRequest
	Vars     map[string]interface{}
}

func (r RouteCtx) Route() Route {
	return r.route
}

func (r RouteCtx) Response() *RouteResponse {
	return r.response
}

func (r RouteCtx) Request() *RouteRequest {
	return r.request
}

func (r *RouteCtx) SetVar(key string, vars interface{}) {
	if r.Vars == nil {
		r.Vars = make(map[string]interface{})
	}
	r.Vars[key] = vars
}

func (r *RouteCtx) WithVars(key string, vars interface{}) *RouteCtx {
	r.SetVar(key, vars)
	return r
}

type Route struct {
	Module  string
	Method  string
	Pattern string
	Handler func(ctx *RouteCtx)
}

type Router struct {
	Routes []Route
}

func (r *Router) To(route Route) {
	r.Routes = append(r.Routes, route)
}

func MuxIt(r *Router, mux *http.ServeMux) {
	for _, route := range r.Routes {
		mux.HandleFunc(fmt.Sprintf("%s %s", route.Method, route.Pattern),
			func(w http.ResponseWriter, r *http.Request) {
				req := &RouteRequest{
					r: r,
				}

				res := &RouteResponse{
					w: w,
				}

				ctx := &RouteCtx{
					route:    route,
					response: res,
					request:  req,
				}

				Publish(Event{
					Key: EventRouterOnRequest,
					Ctx: ctx,
				})

				route.Handler(ctx)

				if ctx.Response().Status() != http.StatusNoContent && !ctx.Response().Sent() {
					ctx.Response().Send(TEXT("No response from: "+route.Method+" "+route.Pattern), http.StatusNotImplemented)
				}
			})
	}
}
