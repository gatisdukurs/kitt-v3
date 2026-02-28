package form

import (
	"kitt/kitt/render"
	"testing"
)

func Test_Form_Label_Context(t *testing.T) {
	t.Run("it returns name", func(t *testing.T) {
		e := render.NewEngine()
		label := NewFormLabel("E-mail", e)
		ctx := NewFormLabelContext(label)

		assertEqual(t, ctx.Name(), "E-mail")
	})
}
