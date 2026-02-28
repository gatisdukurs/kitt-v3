package form

type FormLabelContext interface {
	Name() string
}

type formLabelContext struct {
	label FormLabel
}

func (lc formLabelContext) Name() string {
	return lc.label.Name()
}

func NewFormLabelContext(label FormLabel) FormLabelContext {
	return &formLabelContext{
		label: label,
	}
}
