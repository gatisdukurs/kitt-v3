package form

import (
	"kitt/kitt/render"
	"testing"
)

func Test_Form_Field_Context(t *testing.T) {
	t.Run("it provides values", func(t *testing.T) {
		e := render.NewEngine()
		field := NewFormField("emil", e)
		label := NewFormLabel("E-mail", e)
		control := NewFormControl("email", e)
		ctx := NewFormFieldContext(field)
		field.WithControl(control)
		field.WithLabel(label)

		assertEqual(t, ctx.Id(), field.Id())
		assertEqual(t, ctx.Control(), render.AsHtml(control.Render()))
		assertEqual(t, ctx.Label(), render.AsHtml(label.Render()))
		assertEqual(t, ctx.Errors(), render.AsHtml(field.RenderErrors()))
	})
}
