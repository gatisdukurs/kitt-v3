package render

import (
	"bytes"
	"fmt"
	"strings"
)

type SupportsHTMX interface {
	HTMX() string
	WithHTMX(content string, oob ...string) Layout
}

type LayoutCtx interface {
	Slot(name string) AsHtml
	Ctx(name string) interface{}
}

type Layout interface {
	Renderable
	SupportsHTMX
	Slot(slot string) []Partial
	Ctx() AnyCtx
	WithPartial(slot string, p Partial) Layout
	WithCtx(AnyCtx) Layout
}

type layoutCtx struct {
	l Layout
}

func (lc layoutCtx) Slot(slot string) AsHtml {
	var buf bytes.Buffer

	for _, p := range lc.l.Slot(slot) {
		buf.WriteString(p.Render())
	}

	return AsHtml(buf.String())
}

func (lc layoutCtx) Ctx(name string) interface{} {
	ctx := lc.l.Ctx()
	return ctx[name]
}

type layout struct {
	e               Engine
	name            string
	slots           map[string][]Partial
	ctx             AnyCtx
	HTMXcontentSlot string
	HTMXoobSlots    []string
}

func (l *layout) Render() string {
	var buf bytes.Buffer
	err := l.e.Render(&buf, l.name, newLayoutCtx(l))

	if err != nil {
		return err.Error()
	}

	return strings.TrimSpace(buf.String())
}

func (l *layout) HTMX() string {
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

func (l *layout) renderSlot(slot string) string {
	var buf bytes.Buffer
	for _, p := range l.Slot(slot) {
		buf.WriteString(p.Render())
	}
	return strings.TrimSpace(buf.String())
}

func (l *layout) WithHTMX(contentSlot string, oobSlots ...string) Layout {
	l.HTMXcontentSlot = contentSlot
	l.HTMXoobSlots = oobSlots
	return l
}

func (l layout) Ctx() AnyCtx {
	return l.ctx
}

func (l *layout) WithPartial(slot string, p Partial) Layout {
	l.slots[slot] = append(l.slots[slot], p)
	return l
}

func (l *layout) WithCtx(ctx AnyCtx) Layout {
	l.ctx = ctx
	return l
}

func (l layout) Slot(slot string) []Partial {
	return l.slots[slot]
}

func NewLayout(name string, e Engine) Layout {
	l := &layout{
		name:  name,
		e:     e,
		slots: make(map[string][]Partial),
	}
	return l
}

func newLayoutCtx(l Layout) LayoutCtx {
	return &layoutCtx{
		l: l,
	}
}
