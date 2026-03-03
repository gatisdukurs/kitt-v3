package form

type FormSuccessContext interface {
	Message() string
}

type formSuccessCtx struct {
	fs FormSuccess
}

func (f formSuccessCtx) Message() string {
	return f.fs.Message()
}

func NewFormSuccessContext(fs FormSuccess) FormSuccessContext {
	return &formSuccessCtx{
		fs: fs,
	}
}
