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

type FormControl interface {
	render.Renderable
	Id() string
	Type() string
	Name() string
	WithType(fieldType string) FormControl
	WithId(id string) FormControl
	WithValue(value string) FormControl
	WithValidators(validators ...FormValidator) FormControl
	Value() string
	Validate() (bool, []string)
}

type formControl struct {
	id         string
	name       string
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
	inputTpl := `<input class="field" name="{{ .Name }}" id="{{ .Id }}" type="{{ .Type }}" value="{{ .Value }}" />`
	textareaTpl := `<textarea class="field" name="{{ .Name }}" id="{{ .Id }}">{{ .Value }}</textarea>`

	engine.WithTemplate("form.input", inputTpl)
	engine.WithTemplate("form.textarea", textareaTpl)

	return &formControl{
		e:         engine,
		name:      name,
		id:        name,
		fieldType: FIELD_TEXT,
	}
}
