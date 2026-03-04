package form

import "kitt/kitt/render"

type FormFieldContext interface {
	Id() string
	Label() render.AsHtml
	Control() render.AsHtml
	Errors() render.AsHtml
}

type formFieldCtx struct {
	formField FormField
}

func (c formFieldCtx) Id() string {
	return c.formField.Id()
}

func (c formFieldCtx) Label() render.AsHtml {
	return render.AsHtml(c.formField.RenderLabel())
}

func (c formFieldCtx) Control() render.AsHtml {
	return render.AsHtml(c.formField.RenderControl())
}

func (c formFieldCtx) Errors() render.AsHtml {
	return render.AsHtml(c.formField.RenderErrors())
}

func NewFormFieldContext(formField FormField) FormFieldContext {
	return &formFieldCtx{
		formField: formField,
	}
}
