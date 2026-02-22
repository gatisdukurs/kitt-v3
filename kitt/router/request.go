package router

type Request interface{}

type request struct{}

func NewRequest() Request {
	return &request{}
}
