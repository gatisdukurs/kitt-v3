package form

import (
	"bytes"
	"kitt/kitt/render"
)

type FormError interface {
	render.Renderable
	Message() string
	WithMessage(msg string) FormError
}

type formError struct {
	e   render.Engine
	msg string
}

func (f *formError) Message() string {
	return f.msg
}

func (f *formError) WithMessage(msg string) FormError {
	f.msg = msg
	return f
}

func (f *formError) Render() string {
	var buf bytes.Buffer

	f.e.Render(&buf, "form.error", NewFormErrorContext(f))

	return buf.String()
}

func NewFormError(msg string, e render.Engine) FormError {
	tpl := `<div class="alert alert--danger">{{ .Message }}</div>`

	e.WithTemplate("form.error", tpl)

	return &formError{
		e:   e,
		msg: msg,
	}
}
