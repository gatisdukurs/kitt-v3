package form

import (
	"bytes"
	"fmt"
	"kitt/kitt/render"
	"kitt/kitt/router"
	"net/http"
	"net/url"
	"strings"
)

type Form interface {
	router.RouteResponseSendable
	Success() FormSuccess
	Error() FormError
	WithError(msg string) Form
	WithSuccess(msg string) Form
	WithField(control FormField) Form
	WithMethod(method string) Form
	WithAction(action string) Form
	WithActions(actions FormActions) Form
	WithValues(values url.Values) Form
	WithAttribute(key string, value string) Form
	WithHTMXPost() Form
	WithHTMXGet() Form
	WithHTMXTarget(sel string) Form
	WithHTMXSwap(swap string) Form
	WithHTMX() Form
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
}

type form struct {
	e           render.Engine
	attributes  []string
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

func (f *form) WithHTMXPost() Form {
	return f.WithAttribute("hx-post", f.action)
}

func (f *form) WithHTMXGet() Form {
	return f.WithAttribute("hx-get", f.action)
}

func (f *form) WithHTMXTarget(sel string) Form {
	return f.WithAttribute("hx-target", sel)
}

func (f *form) WithHTMXSwap(swap string) Form {
	return f.WithAttribute("hx-swap", swap)
}

func (f *form) WithHTMX() Form {
	f.WithHTMXPost()
	f.WithHTMXSwap("outerHTML")
	f.WithHTMXTarget("#" + f.id)
	return f
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
	if key == "" {
		return f
	}

	var attr string
	if value != "" {
		attr = fmt.Sprintf(`%s="%s"`, key, value)
	} else {
		attr = key
	}
	f.attributes = append(f.attributes, attr)
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

	return " " + strings.Join(f.attributes, " ")
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

func (f form) HTMX() string {
	f.WithHTMX()
	return f.Render()
}

func NewForm(id string, e render.Engine) Form {
	template := `<form class="form" action="{{ .Action }}" method="{{ .Method }}" id="{{ .Id }}"{{ .Attributes }}>{{ .Success }}{{ .Error }}{{ .Fields }}{{ .Actions }}</form>`
	e.WithTemplate("form", template)

	return &form{
		e:      e,
		id:     id,
		method: http.MethodPost,
		action: "/",
	}
}
