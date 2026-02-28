package form

import (
	"bytes"
	"kitt/kitt/render"
	"net/http"
)

type Form interface {
	render.Renderable
	WithControl(control FormControl) Form
	WithMethod(method string) Form
	WithAction(action string) Form
	WithId(id string) Form
	RenderControls() string
	Action() string
	Method() string
	Id() string
}

type form struct {
	e        render.Engine
	controls []FormControl
	method   string
	action   string
	id       string
}

func (f *form) WithControl(control FormControl) Form {
	return f
}

func (f *form) WithMethod(method string) Form {
	f.method = method
	return f
}

func (f *form) WithAction(action string) Form {
	f.action = action
	return f
}

func (f *form) WithId(id string) Form {
	f.id = id
	return f
}

func (f *form) Render() string {
	var buf bytes.Buffer
	f.e.Render(&buf, "form", NewFormContext(f))
	return buf.String()
}

func (f form) RenderControls() string {
	var buf bytes.Buffer

	for _, c := range f.controls {
		buf.WriteString(c.Render())
	}

	return buf.String()
}

func (f form) Action() string {
	return f.action
}

func (f form) Method() string {
	return f.method
}

func (f form) Id() string {
	return f.id
}

func NewForm(id string, e render.Engine) Form {
	template := `<form action="{{ .Action }}" method="{{ .Method }}" id="{{ .Id }}">{{ .Controls }}</form>`
	e.WithTemplate("form", template)

	return &form{
		e:      e,
		id:     id,
		method: http.MethodPost,
		action: "/",
	}
}
