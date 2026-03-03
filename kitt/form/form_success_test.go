package form

import (
	"kitt/kitt/render"
	"testing"
)

func Test_Form_Success(t *testing.T) {
	t.Run("it renders", func(t *testing.T) {
		e := render.NewEngine()
		fs := NewFormSuccess("Account Created.", e)

		assertEqual(t, fs.Render(), `<div class="success">Account Created.</div>`)
	})

	t.Run("it sets message", func(t *testing.T) {
		msg := "Account Created."
		e := render.NewEngine()
		fs := NewFormSuccess(msg, e)

		assertEqual(t, fs.Message(), msg)

		fs.WithMessage("New Success.")

		assertEqual(t, fs.Message(), "New Success.")
	})
}
