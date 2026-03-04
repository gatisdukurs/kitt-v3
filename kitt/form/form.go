package form

import (
	"bytes"
	"kitt/kitt/render"
	"net/http"
	"net/url"
)

type Form interface {
	render.Renderable
	Success() FormSuccess
	Error() FormError
	WithError(err FormError) Form
	WithSuccess(succ FormSuccess) Form
	WithControl(control FormField) Form
	WithMethod(method string) Form
	WithAction(action string) Form
	WithValues(values url.Values) Form
	WithId(id string) Form
	RenderControls() string
	RenderError() string
	RenderSuccess() string
	Action() string
	Method() string
	Id() string
	Control(id string) FormField
	Validate() bool
}

type form struct {
	e           render.Engine
	controls    []FormField
	method      string
	action      string
	id          string
	formError   FormError
	formSuccess FormSuccess
}

func (f form) Error() FormError {
	return f.formError
}

func (f form) Success() FormSuccess {
	return f.formSuccess
}

func (f form) Validate() bool {
	isValid := true

	for _, c := range f.controls {
		if ok, errs := c.Control().Validate(); !ok {
			c.WithErrors(errs)
			isValid = false
		}
	}

	return isValid
}

func (f *form) WithError(err FormError) Form {
	f.formError = err
	return f
}

func (f *form) WithSuccess(succ FormSuccess) Form {
	f.formSuccess = succ
	return f
}

func (f *form) WithValues(values url.Values) Form {
	for _, c := range f.controls {
		if c.Control() != nil {
			value := values.Get(c.Control().Name())
			c.Control().WithValue(value)
		}
	}

	return f
}

func (f *form) WithControl(control FormField) Form {
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

func (f form) RenderSuccess() string {
	if f.formSuccess == nil {
		return ""
	}
	return f.formSuccess.Render()
}

func (f form) Control(id string) FormField {
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
	template := `<form class="form" action="{{ .Action }}" method="{{ .Method }}" id="{{ .Id }}">{{ .Success }}{{ .Error }}{{ .Controls }}</form>`
	e.WithTemplate("form", template)

	return &form{
		e:      e,
		id:     id,
		method: http.MethodPost,
		action: "/",
	}
}
