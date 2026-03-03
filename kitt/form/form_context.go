package form

import "kitt/kitt/render"

type FormContext interface {
	Action() string
	Method() string
	Id() string
	Controls() render.AsHtml
	Error() render.AsHtml
	Success() render.AsHtml
}

type formCtx struct {
	form Form
}

func (ctx formCtx) Action() string {
	return ctx.form.Action()
}

func (ctx formCtx) Method() string {
	return ctx.form.Method()
}

func (ctx formCtx) Id() string {
	return ctx.form.Id()
}

func (ctx formCtx) Controls() render.AsHtml {
	return render.AsHtml(ctx.form.RenderControls())
}

func (ctx formCtx) Error() render.AsHtml {
	return render.AsHtml(ctx.form.RenderError())
}

func (ctx formCtx) Success() render.AsHtml {
	return render.AsHtml(ctx.form.RenderSuccess())
}

func NewFormContext(form Form) FormContext {
	return &formCtx{
		form: form,
	}
}
