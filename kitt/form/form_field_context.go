package form

type FormFieldContext interface {
	Name() string
	Id() string
	Type() string
	Value() string
}

type fieldCtx struct {
	field FormField
}

func (fc fieldCtx) Name() string {
	return fc.field.Name()
}

func (fc fieldCtx) Id() string {
	return fc.field.Id()
}

func (fc fieldCtx) Type() string {
	return fc.field.Type()
}

func (fc fieldCtx) Value() string {
	return fc.field.Value()
}

// inputTpl := `<input name="{{ .Name }}" id="{{ .Id }}" type="{{ .Type }}" value={{ .Value }} />`
func NewFormFieldContext(field FormField) FormFieldContext {
	return &fieldCtx{
		field: field,
	}
}
