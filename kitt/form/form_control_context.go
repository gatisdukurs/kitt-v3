package form

type FormControlContext interface {
	Name() string
	Id() string
	Type() string
	Value() string
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

func NewFormControlContext(control FormControl) FormControlContext {
	return &formControlCtx{
		control: control,
	}
}
