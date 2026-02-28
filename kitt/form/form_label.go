package form

import (
	"bytes"
	"kitt/kitt/render"
)

type FormLabel interface {
	render.Renderable
	Name() string
}

type label struct {
	e    render.Engine
	name string
}

func (l label) Render() string {
	var buf bytes.Buffer

	l.e.Render(&buf, "form.label", NewFormLabelContext(l))

	return buf.String()
}

func (l label) Name() string {
	return l.name
}

func NewFormLabel(name string, engine render.Engine) FormLabel {
	template := `<label class="label">{{ .Name }}</label>`
	engine.WithTemplate("form.label", template)

	return &label{
		name: name,
		e:    engine,
	}
}
