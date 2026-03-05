package form

import "kitt/kitt/render"

type FormControlContext interface {
	Name() string
	Id() string
	Type() string
	Value() string
	Attributes() render.AsAttr
}

type formControlCtx struct {
	control FormControl
}

func (c formControlCtx) Name() string {
	return c.control.Name()
}

func (c formControlCtx) Id() string {
	return c.control.Id()
}

func (c formControlCtx) Type() string {
	return c.control.Type()
}

func (c formControlCtx) Value() string {
	return c.control.Value()
}

func (c formControlCtx) Attributes() render.AsAttr {
	return render.AsAttr(c.control.RenderAttributes())
}

func NewFormControlContext(control FormControl) FormControlContext {
	return &formControlCtx{
		control: control,
	}
}
