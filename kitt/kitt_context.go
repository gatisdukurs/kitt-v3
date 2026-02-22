package kitt

import "kitt/kitt/render"

type KittBasicContext = render.AnyCtx

type KittContext interface {
	Basic() KittBasicContext
	Set(key string, value interface{}) KittContext
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

func NewKittCtx() KittContext {
	return &kittCtx{
		basic: make(KittBasicContext),
	}
}
