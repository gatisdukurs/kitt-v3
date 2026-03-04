package render

import (
	"bytes"
	"fmt"
	"strings"
)

type SupportsHTMX interface {
	HTMX() string
	WithHTMX(content string, oob ...string) View
}

type ViewCtx interface {
	Slot(name string) AsHtml
	Ctx(name string) interface{}
}

type View interface {
	Renderable
	SupportsHTMX
	Slot(slot string) []Renderable
	Ctx() AnyCtx
	WithPartial(slot string, p Renderable) View
	WithCtx(AnyCtx) View
}

type viewCtx struct {
	l View
}

func (lc viewCtx) Slot(slot string) AsHtml {
	var buf bytes.Buffer

	for _, p := range lc.l.Slot(slot) {
		buf.WriteString(p.Render())
	}

	return AsHtml(buf.String())
}

func (lc viewCtx) Ctx(name string) interface{} {
	ctx := lc.l.Ctx()
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

func (l *view) Render() string {
	var buf bytes.Buffer
	err := l.e.Render(&buf, l.name, newViewCtx(l))

	if err != nil {
		return err.Error()
	}

	return strings.TrimSpace(buf.String())
}

func (l *view) HTMX() string {
	if l.HTMXcontentSlot == "" {
		return l.Render()
	}

	var buf bytes.Buffer

	buf.WriteString(l.renderSlot(l.HTMXcontentSlot))

	for _, slot := range l.HTMXoobSlots {
		buf.WriteString(
			fmt.
				Sprintf(`<div id="%s" hx-swap-oob="true">%s</div>`,
					slot,
					l.renderSlot(slot),
				),
		)
	}

	return strings.TrimSpace(buf.String())
}

func (l *view) renderSlot(slot string) string {
	var buf bytes.Buffer
	for _, p := range l.Slot(slot) {
		buf.WriteString(p.Render())
	}
	return strings.TrimSpace(buf.String())
}

func (l *view) WithHTMX(contentSlot string, oobSlots ...string) View {
	l.HTMXcontentSlot = contentSlot
	l.HTMXoobSlots = oobSlots
	return l
}

func (l view) Ctx() AnyCtx {
	return l.ctx
}

func (l *view) WithPartial(slot string, p Renderable) View {
	l.slots[slot] = append(l.slots[slot], p)
	return l
}

func (l *view) WithCtx(ctx AnyCtx) View {
	l.ctx = ctx
	return l
}

func (l view) Slot(slot string) []Renderable {
	return l.slots[slot]
}

func NewView(name string, e Engine) View {
	l := &view{
		name:  name,
		e:     e,
		slots: make(map[string][]Renderable),
	}
	return l
}

func newViewCtx(l View) ViewCtx {
	return &viewCtx{
		l: l,
	}
}
