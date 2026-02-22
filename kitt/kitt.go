package kitt

import (
	"kitt/kitt/render"
)

type KittTemplatePattern = string
type KittTemplatePatterns = []KittTemplatePattern
type KittTemplateFuncs = render.Funcs
type KittBasicContext = render.AnyCtx

type KittContext interface {
	Basic() KittBasicContext
	Set(key string, value interface{}) KittContext
}

type Kitt interface {
	Layout(name string) render.Layout
	Partial(name string) render.Partial
	Ctx() KittContext
	WithTemplate(name string, str string) Kitt
	WithTemplates(patterns KittTemplatePatterns) Kitt
	WithTemplateFuncs(funcs render.Funcs) Kitt
	InTesting()
}

var kittInstance Kitt

func ensureInitialized() {
	if kittInstance == nil {
		kittInstance = newKitt()
	}
}

type kittCtx struct {
	basic KittBasicContext
}

func (c *kittCtx) Set(key string, value interface{}) KittContext {
	c.basic[key] = value
	return c
}

func (c kittCtx) Basic() KittBasicContext {
	return c.basic
}

type kitt struct {
	renderer render.Engine
}

func (k kitt) Layout(name string) render.Layout {
	l := render.NewLayout(name, k.renderer)
	return l
}

func (k kitt) Partial(name string) render.Partial {
	p := render.NewPartial(name, k.renderer)
	return p
}

func (k kitt) Ctx() KittContext {
	return newKittCtx()
}

func (k *kitt) WithTemplates(patterns KittTemplatePatterns) Kitt {
	for _, pattern := range patterns {
		k.renderer = k.renderer.WithTemplates(pattern)
	}
	return k
}

func (k *kitt) WithTemplate(name string, str string) Kitt {
	k.renderer = k.renderer.WithTemplate(name, str)
	return k
}

func (k *kitt) WithTemplateFuncs(funcs render.Funcs) Kitt {
	k.renderer = k.renderer.WithFuncs(funcs)
	return k
}

func (k *kitt) InTesting() {
	k.init()
}

func (k *kitt) init() {
	k.renderer = render.NewEngine()
}

func K() Kitt {
	ensureInitialized()
	return kittInstance
}

func newKitt() Kitt {
	k := &kitt{}
	k.init()
	return k
}

func newKittCtx() KittContext {
	return &kittCtx{
		basic: make(KittBasicContext),
	}
}
