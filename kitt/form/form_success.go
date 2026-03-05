package form

import (
	"bytes"
	"kitt/kitt/render"
)

type FormSuccess interface {
	render.Renderable
	Message() string
	WithMessage(msg string) FormSuccess
}

type formSuccess struct {
	e   render.Engine
	msg string
}

func (f formSuccess) Message() string {
	return f.msg
}

func (f *formSuccess) Render() string {
	var buf bytes.Buffer

	f.e.Render(&buf, "form.success", NewFormSuccessContext(f))

	return buf.String()
}

func (f *formSuccess) WithMessage(msg string) FormSuccess {
	f.msg = msg
	return f
}

func NewFormSuccess(msg string, e render.Engine) FormSuccess {
	tpl := `<div class="alert alert--success flash">{{ .Message }}</div>`

	e.WithTemplate("form.success", tpl)

	return &formSuccess{
		msg: msg,
		e:   e,
	}
}
