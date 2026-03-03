package form

type FormActionContext interface {
	Id() string
	Name() string
	Value() string
	Label() string
}

type formActionCtx struct {
	fa FormAction
}

func (c formActionCtx) Id() string {
	return c.fa.Id()
}
func (c formActionCtx) Name() string {
	return c.fa.Name()
}
func (c formActionCtx) Value() string {
	return c.fa.Value()
}
func (c formActionCtx) Label() string {
	return c.fa.Label()
}

func NewFormActionContext(fa FormAction) FormActionContext {
	return &formActionCtx{
		fa: fa,
	}
}
