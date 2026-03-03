package form

type FormErrorContext interface {
	Message() string
}

type formErrorCtx struct {
	fe FormError
}

func (c formErrorCtx) Message() string {
	return c.fe.Message()
}

func NewFormErrorContext(fe FormError) FormErrorContext {
	return &formErrorCtx{
		fe: fe,
	}
}
