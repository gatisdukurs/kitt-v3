package form

import (
	"kitt/kitt/render"
	"testing"
)

func Test_Form_Error_Context(t *testing.T) {
	t.Run("it provides values", func(t *testing.T) {
		e := render.NewEngine()
		err := NewFormError("Error.", e)
		ctx := NewFormErrorContext(err)

		assertEqual(t, ctx.Message(), err.Message())
	})
}
