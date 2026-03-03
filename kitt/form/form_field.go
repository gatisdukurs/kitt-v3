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
	WithValidators(validators ...FormValidator) FormField
	Value() string
	Validate() (bool, []string)
}

type formField struct {
	id         string
	name       string
	e          render.Engine
	fieldType  string
	value      string
	validators []FormValidator
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

func (ff *formField) WithValidators(validators ...FormValidator) FormField {
	ff.validators = validators
	return ff
}

func (ff *formField) Validate() (bool, []string) {
	errors := []string{}
	value := ff.value

	for _, v := range ff.validators {
		if ok, err := v(value); !ok {
			errors = append(errors, err)
		}
	}

	return len(errors) == 0, errors
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
