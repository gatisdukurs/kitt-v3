package form

import (
	"bytes"
	"kitt/kitt/render"
)

type FormField interface {
	render.Renderable
	Id() string
	Label() FormLabel
	Control() FormControl
	Errors() []string
	RenderControl() string
	RenderLabel() string
	RenderErrors() string
	WithLabel(label FormLabel) FormField
	WithControl(control FormControl) FormField
	WithErrors(errs []string) FormField
}

type formField struct {
	e       render.Engine
	id      string
	label   FormLabel
	control FormControl
	errs    []string
}

func (f formField) Id() string {
	return f.id
}

func (f formField) Label() FormLabel {
	return f.label
}

func (f formField) Control() FormControl {
	return f.control
}

func (f formField) Errors() []string {
	return f.errs
}

func (f *formField) WithErrors(errs []string) FormField {
	f.errs = errs
	return f
}

func (f *formField) WithLabel(label FormLabel) FormField {
	f.label = label
	return f
}

func (f *formField) WithControl(control FormControl) FormField {
	f.control = control
	return f
}

func (f *formField) Render() string {
	var buf bytes.Buffer

	f.e.Render(&buf, "form.control", NewFormFieldContext(f))

	return buf.String()
}

func (f formField) RenderControl() string {
	if f.control == nil {
		return ""
	}
	return f.control.Render()
}

func (f formField) RenderLabel() string {
	if f.label == nil {
		return ""
	}
	return f.label.Render()
}

func (f formField) RenderErrors() string {
	if len(f.errs) == 0 {
		return ""
	}

	var buf bytes.Buffer

	f.e.Render(&buf, "form.errors", f.errs)

	return buf.String()
}

func NewFormField(id string, e render.Engine) FormField {
	controlTpl := `<div class="control" id="{{ .Id }}">{{ .Label }}{{ .Control }}{{ .Errors }}</div>`
	errsTpl := `<ul class="errors">{{ range . }}<li>{{ . }}</li>{{ end }}</ul>`

	e.WithTemplate("form.control", controlTpl)
	e.WithTemplate("form.errors", errsTpl)

	return &formField{
		e:  e,
		id: id,
	}
}
