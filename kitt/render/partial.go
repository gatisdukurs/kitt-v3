package render

import (
	"bytes"
	"strings"
)

type PartialCtx interface {
	Ctx(name string) interface{}
}

type Partial interface {
	Renderable
	Ctx() AnyCtx
	WithCtx(ctx AnyCtx) Partial
}

type partialCtx struct {
	p Partial
}

func (pc partialCtx) Ctx(name string) interface{} {
	ctx := pc.p.Ctx()
	return ctx[name]
}

type partial struct {
	ctx  AnyCtx
	name string
	e    Engine
}

func (p *partial) WithCtx(ctx AnyCtx) Partial {
	p.ctx = ctx
	return p
}

func (p partial) Ctx() AnyCtx {
	return p.ctx
}

func (p *partial) Render() string {
	var buf bytes.Buffer
	err := p.e.Render(&buf, p.name, newPartialCtx(p))

	if err != nil {
		return err.Error()
	}

	return strings.TrimSpace(buf.String())
}

func NewPartial(name string, e Engine) Partial {
	return &partial{
		name: name,
		e:    e,
	}
}

func newPartialCtx(p Partial) PartialCtx {
	return &partialCtx{
		p: p,
	}
}
