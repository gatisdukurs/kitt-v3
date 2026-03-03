package form

import "kitt/kitt/render"

type FormActionsContext interface {
	Id() string
	Actions() render.AsHtml
}

type formActionsCtx struct {
	fa FormActions
}

func (f formActionsCtx) Id() string {
	return f.fa.Id()
}

func (f formActionsCtx) Actions() render.AsHtml {
	return render.AsHtml(f.fa.RenderActions())
}

func NewFormActionsContext(fa FormActions) FormActionsContext {
	return &formActionsCtx{
		fa: fa,
	}
}
