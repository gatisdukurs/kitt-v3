package form

import (
	"kitt/kitt/render"
	"testing"
)

func Test_Field_Context(t *testing.T) {
	t.Run("it returns values", func(t *testing.T) {
		e := render.NewEngine()
		field := NewFormField("email", e)
		ctx := NewFormFieldContext(field)

		assertEqual(t, ctx.Name(), "email")
		assertEqual(t, ctx.Id(), "email")
	})
}
