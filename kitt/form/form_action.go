package form

import (
	"bytes"
	"kitt/kitt/render"
)

type FormAction interface {
	render.Renderable
	Id() string
	Name() string
	Label() string
	Value() string
	WithId(id string) FormAction
	WithLabel(label string) FormAction
	WithValue(value string) FormAction
}

type formAction struct {
	e     render.Engine
	name  string
	id    string
	label string
	value string
}

func (t formAction) Id() string {
	return t.id
}

func (t formAction) Name() string {
	return t.name
}

func (t formAction) Label() string {
	return t.label
}

func (t formAction) Value() string {
	return t.value
}

func (f *formAction) Render() string {
	var buf bytes.Buffer

	f.e.Render(&buf, "form.button", NewFormActionContext(f))

	return buf.String()
}

func (f *formAction) WithId(id string) FormAction {
	f.id = id
	return f
}

func (f *formAction) WithLabel(label string) FormAction {
	f.label = label
	return f
}

func (f *formAction) WithValue(value string) FormAction {
	f.value = value
	return f
}

func NewFormAction(name string, e render.Engine) FormAction {
	button := `<button class="btn" id="{{ .Id }}" name="{{ .Name }}" value="{{ .Value }}">{{ .Label }}</button>`

	e.WithTemplate("form.button", button)

	return &formAction{
		name:  name,
		label: name,
		id:    name,
		e:     e,
	}
}
