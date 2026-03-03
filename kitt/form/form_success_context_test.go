package form

import (
	"kitt/kitt/render"
	"testing"
)

func Test_Form_Success_Context(t *testing.T) {
	t.Run("it provides values", func(t *testing.T) {
		e := render.NewEngine()
		fs := NewFormSuccess("Success.", e)
		ctx := NewFormSuccessContext(fs)

		assertEqual(t, ctx.Message(), fs.Message())
	})
}
