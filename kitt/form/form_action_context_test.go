package form

import (
	"kitt/kitt/render"
	"testing"
)

func Test_Form_Action_Context(t *testing.T) {
	t.Run("it returns values", func(t *testing.T) {
		e := render.NewEngine()
		f := NewFormAction("Button", e)
		f.WithId("id")
		f.WithLabel("label")
		f.WithValue("value")

		ctx := NewFormActionContext(f)

		assertEqual(t, ctx.Id(), f.Id())
		assertEqual(t, ctx.Name(), f.Name())
		assertEqual(t, ctx.Value(), f.Value())
		assertEqual(t, ctx.Label(), f.Label())
	})
}
