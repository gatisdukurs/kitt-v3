package form

import (
	"bytes"
	"kitt/kitt/render"
)

type FormActions interface {
	render.Renderable
	Id() string
	Actions() []FormAction
	WithAction(action FormAction) FormActions
	RenderActions() string
}

type formActions struct {
	e       render.Engine
	id      string
	actions []FormAction
}

func (fa formActions) Id() string {
	return fa.id
}

func (fa *formActions) Render() string {
	var buf bytes.Buffer

	fa.e.Render(&buf, "form.actions", NewFormActionsContext(fa))

	return buf.String()
}

func (fa formActions) RenderActions() string {
	if len(fa.actions) == 0 {
		return ""
	}

	var buf bytes.Buffer

	for _, a := range fa.actions {
		buf.WriteString(a.Render())
	}

	return buf.String()
}

func (fa formActions) Actions() []FormAction {
	return fa.actions
}

func (fa *formActions) WithAction(action FormAction) FormActions {
	fa.actions = append(fa.actions, action)
	return fa
}

func NewFormActions(id string, e render.Engine) FormActions {
	template := `<div class="actions" id="{{ .Id }}">{{ .Actions }}</div>`

	e.WithTemplate("form.actions", template)

	return &formActions{
		e:  e,
		id: id,
	}
}
