package form

import (
	"bytes"
	"kitt/kitt/render"
)

type FormControl interface {
	render.Renderable
	Id() string
	Label() FormLabel
	Field() FormField
	RenderField() string
	RenderLabel() string
	WithLabel(label FormLabel) FormControl
	WithField(field FormField) FormControl
}

type formControl struct {
	e     render.Engine
	id    string
	label FormLabel
	field FormField
}

func (f formControl) Id() string {
	return f.id
}

func (f formControl) Label() FormLabel {
	return f.label
}

func (f formControl) Field() FormField {
	return f.field
}

func (f *formControl) WithLabel(label FormLabel) FormControl {
	f.label = label
	return f
}

func (f *formControl) WithField(field FormField) FormControl {
	f.field = field
	return f
}

func (f *formControl) Render() string {
	var buf bytes.Buffer

	f.e.Render(&buf, "form.control", NewFormControlContext(f))

	return buf.String()
}

func (f formControl) RenderField() string {
	if f.field == nil {
		return ""
	}
	return f.field.Render()
}

func (f formControl) RenderLabel() string {
	if f.label == nil {
		return ""
	}
	return f.label.Render()
}

func NewFormControl(id string, e render.Engine) FormControl {
	template := `<div class="control" id="{{ .Id }}">{{ .Label }}{{ .Field }}</div>`
	e.WithTemplate("form.control", template)

	return &formControl{
		e:  e,
		id: id,
	}
}
