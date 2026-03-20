package render

import (
	"bytes"
	"strings"
)

type ViewCtx interface {
	Slot(name string) AsHtml
	Ctx(name string) interface{}
}

type View interface {
	Renderable
	Slot(slot string) []Renderable
	Ctx() AnyCtx
	WithPartial(slot string, p Renderable) View
	WithCtx(AnyCtx) View
}

type viewCtx struct {
	view View
}

func (vc viewCtx) Slot(slot string) AsHtml {
	var buf bytes.Buffer

	for _, p := range vc.view.Slot(slot) {
		buf.WriteString(p.Render())
	}

	return AsHtml(buf.String())
}

func (vc viewCtx) Ctx(name string) interface{} {
	ctx := vc.view.Ctx()
	return ctx[name]
}

type view struct {
	e               Engine
	name            string
	slots           map[string][]Renderable
	ctx             AnyCtx
	HTMXcontentSlot string
	HTMXoobSlots    []string
}

func (v *view) Render() string {
	var buf bytes.Buffer
	err := v.e.Render(&buf, v.name, newViewCtx(v))

	if err != nil {
		return err.Error()
	}

	return strings.TrimSpace(buf.String())
}

func (v *view) renderSlot(slot string) string {
	var buf bytes.Buffer
	for _, p := range v.Slot(slot) {
		buf.WriteString(p.Render())
	}
	return strings.TrimSpace(buf.String())
}

func (v view) Ctx() AnyCtx {
	return v.ctx
}

func (v *view) WithPartial(slot string, p Renderable) View {
	v.slots[slot] = append(v.slots[slot], p)
	return v
}

func (v *view) WithCtx(ctx AnyCtx) View {
	v.ctx = ctx
	return v
}

func (v view) Slot(slot string) []Renderable {
	return v.slots[slot]
}

func NewView(name string, e Engine) View {
	v := &view{
		name:  name,
		e:     e,
		slots: make(map[string][]Renderable),
	}
	return v
}

func newViewCtx(v View) ViewCtx {
	return &viewCtx{
		view: v,
	}
}
