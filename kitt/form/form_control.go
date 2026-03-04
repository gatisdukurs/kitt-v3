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
	Errors() []string
	RenderField() string
	RenderLabel() string
	RenderErrors() string
	WithLabel(label FormLabel) FormControl
	WithField(field FormField) FormControl
	WithErrors(errs []string) FormControl
}

type formControl struct {
	e     render.Engine
	id    string
	label FormLabel
	field FormField
	errs  []string
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

func (f formControl) Errors() []string {
	return f.errs
}

func (f *formControl) WithErrors(errs []string) FormControl {
	f.errs = errs
	return f
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

func (f formControl) RenderErrors() string {
	if len(f.errs) == 0 {
		return ""
	}

	var buf bytes.Buffer

	f.e.Render(&buf, "form.errors", f.errs)

	return buf.String()
}

func NewFormControl(id string, e render.Engine) FormControl {
	control := `<div class="control" id="{{ .Id }}">{{ .Label }}{{ .Field }}{{ .Errors }}</div>`
	errs := `<ul class="errors">{{ range . }}<li>{{ . }}</li>{{ end }}</ul>`

	e.WithTemplate("form.control", control)
	e.WithTemplate("form.errors", errs)

	return &formControl{
		e:  e,
		id: id,
	}
}
