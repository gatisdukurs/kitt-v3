package form

import (
	"bytes"
	"kitt/kitt/render"
	"net/http"
	"net/url"
)

type Form interface {
	render.Renderable
	Error() FormError
	SuccessMsg() string
	WithError(err FormError) Form
	WithSuccess(message string) Form
	WithControl(control FormControl) Form
	WithMethod(method string) Form
	WithAction(action string) Form
	WithValues(values url.Values) Form
	WithId(id string) Form
	RenderControls() string
	RenderError() string
	Action() string
	Method() string
	Id() string
	Control(id string) FormControl
}

type form struct {
	e              render.Engine
	controls       []FormControl
	method         string
	action         string
	id             string
	formError      FormError
	successMessage string
}

func (f *form) Error() FormError {
	return f.formError
}

func (f *form) SuccessMsg() string {
	return f.successMessage
}

func (f *form) WithError(err FormError) Form {
	f.formError = err
	return f
}

func (f *form) WithSuccess(message string) Form {
	f.successMessage = message
	return f
}

func (f *form) WithValues(values url.Values) Form {
	for _, c := range f.controls {
		if c.Field() != nil {
			value := values.Get(c.Field().Name())
			c.Field().WithValue(value)
		}
	}

	return f
}

func (f *form) WithControl(control FormControl) Form {
	f.controls = append(f.controls, control)
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
	if len(f.controls) == 0 {
		return ""
	}

	var buf bytes.Buffer

	for _, c := range f.controls {
		buf.WriteString(c.Render())
	}

	return buf.String()
}

func (f form) RenderError() string {
	if f.formError == nil {
		return ""
	}
	return f.formError.Render()
}

func (f form) Control(id string) FormControl {
	for _, c := range f.controls {
		if c.Id() == id {
			return c
		}
	}

	return nil
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
	template := `<form class="form" action="{{ .Action }}" method="{{ .Method }}" id="{{ .Id }}">{{ .Error }}{{ .Controls }}</form>`
	e.WithTemplate("form", template)

	return &form{
		e:      e,
		id:     id,
		method: http.MethodPost,
		action: "/",
	}
}
