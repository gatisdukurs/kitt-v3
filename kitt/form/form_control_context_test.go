package form

import (
	"kitt/kitt/render"
	"testing"
)

func Test_Form_Control_Context(t *testing.T) {
	t.Run("it provides values", func(t *testing.T) {
		e := render.NewEngine()
		control := NewFormControl("emil", e)
		label := NewFormLabel("E-mail", e)
		field := NewFormField("email", e)
		ctx := NewFormControlContext(control)
		control.WithField(field)
		control.WithLabel(label)

		assertEqual(t, ctx.Id(), control.Id())
		assertEqual(t, ctx.Field(), render.AsHtml(field.Render()))
		assertEqual(t, ctx.Label(), render.AsHtml(label.Render()))
		assertEqual(t, ctx.Errors(), render.AsHtml(control.RenderErrors()))
	})
}
