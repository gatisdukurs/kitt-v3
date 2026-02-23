package render

import (
	"bytes"
	"strings"
)

type LayoutCtx interface {
	Slot(name string) AsHtml
	Ctx(name string) interface{}
}

type Layout interface {
	Renderable
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
	e     Engine
	name  string
	slots map[string][]Partial
	ctx   AnyCtx
}

func (l *layout) Render() string {
	var buf bytes.Buffer
	err := l.e.Render(&buf, l.name, newLayoutCtx(l))

	if err != nil {
		return err.Error()
	}

	return strings.TrimSpace(buf.String())
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
