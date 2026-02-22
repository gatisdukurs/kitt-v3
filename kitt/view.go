package kitt

type Renderable interface {
	Render() string
}

type view struct {
	name    string
	content Renderable
	ctx     interface{}
}

func (v view) Content() string {
	return v.content.Render()
}

func (v *view) WithContent(content Renderable) *view {
	v.content = content
	return v
}

func (v *view) WithCtx(ctx interface{}) *view {
	v.ctx = ctx
	return v
}

func View(name string) *view {
	return &view{
		name: name,
	}
}

func (v view) Render() string {
	h := HTML(v.name, v.ctx)
	return h.Render()
}
