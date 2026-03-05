package form

import (
	"bytes"
	"html"
	"kitt/kitt/render"
	"strings"
)

const (
	FIELD_TEXT     = "text"
	FIELD_EMAIL    = "email"
	FIELD_PASSWORD = "password"
	FIELD_TEXTAREA = "textarea"
)

type FormControl interface {
	render.Renderable
	Id() string
	Type() string
	Name() string
	WithType(fieldType string) FormControl
	WithId(id string) FormControl
	WithValue(value string) FormControl
	WithValidators(validators ...FormValidator) FormControl
	WithAttribute(key string, value string) FormControl
	RenderAttributes() string
	Value() string
	Validate() (bool, []string)
}

type formControl struct {
	id         string
	name       string
	attributes map[string]string
	e          render.Engine
	fieldType  string
	value      string
	validators []FormValidator
}

func (fc formControl) Render() string {
	switch fc.fieldType {
	case FIELD_TEXT:
		return fc.renderText()
	case FIELD_TEXTAREA:
		return fc.renderTextarea()
	default:
		return "unknown field type: " + fc.fieldType
	}
}

func (fc *formControl) renderText() string {
	var buf bytes.Buffer

	fc.e.Render(&buf, "form.input", NewFormControlContext(fc))

	return buf.String()
}

func (fc *formControl) renderTextarea() string {
	var buf bytes.Buffer

	fc.e.Render(&buf, "form.textarea", NewFormControlContext(fc))

	return buf.String()
}

func (fc formControl) RenderAttributes() string {
	if len(fc.attributes) == 0 {
		return ""
	}

	var b strings.Builder

	for k, v := range fc.attributes {
		if v == "" {
			// boolean attribute
			b.WriteString(" ")
			b.WriteString(k)
			continue
		}

		b.WriteString(" ")
		b.WriteString(k)
		b.WriteString(`="`)
		b.WriteString(html.EscapeString(v))
		b.WriteString(`"`)
	}

	return b.String()
}

func (fc formControl) Type() string {
	return fc.fieldType
}

func (fc formControl) Value() string {
	return fc.value
}

func (fc formControl) Name() string {
	return fc.name
}

func (fc formControl) Id() string {
	return fc.id
}

func (fc *formControl) WithType(fieldType string) FormControl {
	fc.fieldType = fieldType
	return fc
}

func (fc *formControl) WithValue(value string) FormControl {
	fc.value = value
	return fc
}

func (fc *formControl) WithId(id string) FormControl {
	fc.id = id
	return fc
}

func (fc *formControl) WithValidators(validators ...FormValidator) FormControl {
	fc.validators = validators
	return fc
}

func (fc *formControl) WithAttribute(key string, value string) FormControl {
	fc.attributes[key] = value
	return fc
}

func (fc *formControl) Validate() (bool, []string) {
	errors := []string{}
	value := fc.value

	for _, v := range fc.validators {
		if ok, err := v(value); !ok {
			errors = append(errors, err)
		}
	}

	return len(errors) == 0, errors
}

func NewFormControl(name string, engine render.Engine) FormControl {
	inputTpl := `<input class="control" name="{{ .Name }}" id="{{ .Id }}" type="{{ .Type }}" value="{{ .Value }}"{{ .Attributes }}/>`
	textareaTpl := `<textarea class="control" name="{{ .Name }}" id="{{ .Id }}"{{ .Attributes }}>{{ .Value }}</textarea>`

	engine.WithTemplate("form.input", inputTpl)
	engine.WithTemplate("form.textarea", textareaTpl)

	return &formControl{
		attributes: make(map[string]string),
		e:          engine,
		name:       name,
		id:         name,
		fieldType:  FIELD_TEXT,
	}
}
