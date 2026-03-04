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
	WithField(control FormField) Form
	WithMethod(method string) Form
	WithAction(action string) Form
	WithValues(values url.Values) Form
	WithId(id string) Form
	RenderFields() string
	RenderError() string
	RenderSuccess() string
	Action() string
	Method() string
	Id() string
	Field(id string) FormField
	Validate() bool
}

type form struct {
	e           render.Engine
	fields      []FormField
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

	for _, field := range f.fields {
		if ok, errs := field.Control().Validate(); !ok {
			field.WithErrors(errs)
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
	for _, field := range f.fields {
		if field.Control() != nil {
			value := values.Get(field.Control().Name())
			field.Control().WithValue(value)
		}
	}

	return f
}

func (f *form) WithField(field FormField) Form {
	f.fields = append(f.fields, field)
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

func (f form) RenderFields() string {
	if len(f.fields) == 0 {
		return ""
	}

	var buf bytes.Buffer

	for _, c := range f.fields {
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

func (f form) Field(id string) FormField {
	for _, field := range f.fields {
		if field.Id() == id {
			return field
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
	template := `<form class="form" action="{{ .Action }}" method="{{ .Method }}" id="{{ .Id }}">{{ .Success }}{{ .Error }}{{ .Fields }}</form>`
	e.WithTemplate("form", template)

	return &form{
		e:      e,
		id:     id,
		method: http.MethodPost,
		action: "/",
	}
}
