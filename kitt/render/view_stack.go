package render

import (
	"bytes"
	"strings"
)

type ViewStack interface {
	Renderable
	WithRenderable(view Renderable) ViewStack
}

type viewStack struct {
	views []Renderable
}

func (vs *viewStack) WithRenderable(view Renderable) ViewStack {
	vs.views = append(vs.views, view)
	return vs
}

func (vs viewStack) Render() string {
	var buf bytes.Buffer

	for _, v := range vs.views {
		content := v.Render()
		if content == "" {
			continue
		}

		buf.WriteString(content)
	}

	return strings.TrimSpace(buf.String())
}

func NewViewStack() ViewStack {
	return &viewStack{}
}
