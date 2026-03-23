package form

import (
	"bytes"
	"fmt"
	"kitt/kitt/render"
	"kitt/kitt/router"
	"net/http"
	"net/url"
	"sort"
	"strings"
)

type Form interface {
	router.RouteResponseSendable

	Success() FormSuccess
	Error() FormError

	WithError(msg string) Form
	WithSuccess(msg string) Form
	WithField(control FormField) Form
	WithActions(actions FormActions) Form
	WithValues(values url.Values) Form

	WithMethod(method string) Form
	WithAction(action string) Form
	WithAttribute(key string, value string) Form
	WithId(id string) Form

	RenderFields() string
	RenderError() string
	RenderSuccess() string
	RenderActions() string
	RenderAttributes() string

	Action() string
	Method() string
	Id() string
	Field(id string) FormField

	Validate() bool
	Reset()
}

type form struct {
	e           render.Engine
	attributes  map[string]string
	fields      []FormField
	method      string
	action      string
	actions     FormActions
	id          string
	formError   FormError
	formSuccess FormSuccess
}

func (f form) Error() FormError {
	return f.formError
}

func (f form) Success() FormSuccess {
	return f.formSuccess
}

func (f form) Validate() bool {
	isValid := true

	for _, field := range f.fields {
		if ok, errs := field.Control().Validate(); !ok {
			field.WithErrors(errs)
			isValid = false
		}
	}

	return isValid
}

func (f *form) WithError(msg string) Form {
	f.formError = NewFormError(msg, f.e)
	return f
}

func (f *form) WithSuccess(msg string) Form {
	f.formSuccess = NewFormSuccess(msg, f.e)
	return f
}

func (f *form) WithAttribute(key string, value string) Form {
	f.attributes[key] = value
	return f
}

func (f *form) WithValues(values url.Values) Form {
	for _, field := range f.fields {
		if field.Control() != nil {
			value := values.Get(field.Control().Name())
			field.Control().WithValue(value)
		}
	}

	return f
}

func (f *form) Reset() {
	for _, field := range f.fields {
		if field.Control() != nil {
			field.Control().WithValue("")
		}
	}
}

func (f *form) WithField(field FormField) Form {
	f.fields = append(f.fields, field)
	return f
}

func (f *form) WithMethod(method string) Form {
	f.method = method
	return f
}

func (f *form) WithAction(action string) Form {
	f.action = action
	return f
}

func (f *form) WithActions(actions FormActions) Form {
	f.actions = actions
	return f
}

func (f *form) WithId(id string) Form {
	f.id = id
	return f
}

func (f *form) Render() string {
	var buf bytes.Buffer
	f.e.Render(&buf, "form", NewFormContext(f))
	return buf.String()
}

func (f form) RenderFields() string {
	if len(f.fields) == 0 {
		return ""
	}

	var buf bytes.Buffer

	for _, c := range f.fields {
		buf.WriteString(c.Render())
	}

	return buf.String()
}

func (f form) RenderError() string {
	if f.formError == nil {
		return ""
	}
	return f.formError.Render()
}

func (f form) RenderSuccess() string {
	if f.formSuccess == nil {
		return ""
	}
	return f.formSuccess.Render()
}

func (f form) RenderActions() string {
	if f.actions == nil {
		return ""
	}
	return f.actions.Render()
}

func (f form) RenderAttributes() string {
	if len(f.attributes) == 0 {
		return ""
	}

	keys := make([]string, 0, len(f.attributes))
	for k := range f.attributes {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	attrs := make([]string, 0, len(keys))

	for _, k := range keys {
		v := f.attributes[k]
		if v == "" {
			attrs = append(attrs, k)
		} else {
			attrs = append(attrs, fmt.Sprintf(`%s="%s"`, k, v))
		}
	}

	return " " + strings.Join(attrs, " ")
}

func (f form) Field(id string) FormField {
	for _, field := range f.fields {
		if field.Id() == id {
			return field
		}
	}

	return nil
}

func (f form) Action() string {
	return f.action
}

func (f form) Method() string {
	return f.method
}

func (f form) Id() string {
	return f.id
}

func NewForm(id string, e render.Engine) Form {
	template := `<form class="form" action="{{ .Action }}" method="{{ .Method }}" id="{{ .Id }}"{{ .Attributes }}>{{ .Success }}{{ .Error }}{{ .Fields }}{{ .Actions }}</form>`
	e.WithTemplate("form", template)

	return &form{
		e:          e,
		id:         id,
		method:     http.MethodPost,
		action:     "/",
		attributes: make(map[string]string),
	}
}
