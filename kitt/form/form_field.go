package form

import (
	"bytes"
	"kitt/kitt/render"
)

const (
	FIELD_TEXT     = "text"
	FIELD_EMAIL    = "email"
	FIELD_PASSWORD = "password"
	FIELD_TEXTAREA = "textarea"
)

type FormField interface {
	render.Renderable
	Id() string
	Type() string
	Name() string
	WithType(fieldType string) FormField
	WithId(id string) FormField
	WithValue(value string) FormField
	Value() string
}

type formField struct {
	id        string
	name      string
	e         render.Engine
	fieldType string
	value     string
}

func (ff formField) Render() string {
	switch ff.fieldType {
	case FIELD_TEXT:
		return ff.renderText()
	case FIELD_TEXTAREA:
		return ff.renderTextarea()
	default:
		return "unknown field type: " + ff.fieldType
	}
}

func (ff *formField) renderText() string {
	var buf bytes.Buffer

	ff.e.Render(&buf, "form.input", NewFormFieldContext(ff))

	return buf.String()
}

func (ff *formField) renderTextarea() string {
	var buf bytes.Buffer

	ff.e.Render(&buf, "form.textarea", NewFormFieldContext(ff))

	return buf.String()
}

func (ff formField) Type() string {
	return ff.fieldType
}

func (ff formField) Value() string {
	return ff.value
}

func (ff formField) Name() string {
	return ff.name
}
func (ff formField) Id() string {
	return ff.id
}

func (ff *formField) WithType(fieldType string) FormField {
	ff.fieldType = fieldType
	return ff
}

func (ff *formField) WithValue(value string) FormField {
	ff.value = value
	return ff
}

func (ff *formField) WithId(id string) FormField {
	ff.id = id
	return ff
}

func NewFormField(name string, engine render.Engine) FormField {
	inputTpl := `<input class="field" name="{{ .Name }}" id="{{ .Id }}" type="{{ .Type }}" value="{{ .Value }}" />`
	textareaTpl := `<textarea class="field" name="{{ .Name }}" id="{{ .Id }}">{{ .Value }}</textarea>`

	engine.WithTemplate("form.input", inputTpl)
	engine.WithTemplate("form.textarea", textareaTpl)

	return &formField{
		e:         engine,
		name:      name,
		id:        name,
		fieldType: FIELD_TEXT,
	}
}
