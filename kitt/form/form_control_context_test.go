package form

import (
	"kitt/kitt/render"
	"testing"
)

func Test_Form_Control_Context(t *testing.T) {
	t.Run("it returns values", func(t *testing.T) {
		e := render.NewEngine()
		control := NewFormControl("email", e)
		ctx := NewFormControlContext(control)

		assertEqual(t, ctx.Name(), "email")
		assertEqual(t, ctx.Id(), "email")
		assertEqual(t, ctx.Attributes(), render.AsAttr(control.RenderAttributes()))
	})
}
