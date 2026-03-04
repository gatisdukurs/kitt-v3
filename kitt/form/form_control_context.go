package form

import "kitt/kitt/render"

type FormControlContext interface {
	Id() string
	Label() render.AsHtml
	Field() render.AsHtml
	Errors() render.AsHtml
}

type formControlCtx struct {
	formControl FormControl
}

func (c formControlCtx) Id() string {
	return c.formControl.Id()
}
func (c formControlCtx) Label() render.AsHtml {
	return render.AsHtml(c.formControl.RenderLabel())
}
func (c formControlCtx) Field() render.AsHtml {
	return render.AsHtml(c.formControl.RenderField())
}

func (c formControlCtx) Errors() render.AsHtml {
	return render.AsHtml(c.formControl.RenderErrors())
}

func NewFormControlContext(formControl FormControl) FormControlContext {
	return &formControlCtx{
		formControl: formControl,
	}
}
