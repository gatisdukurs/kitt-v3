package form

import (
	"kitt/kitt/render"
	"testing"
)

func Test_Form_Error(t *testing.T) {
	t.Run("it renders", func(t *testing.T) {
		e := render.NewEngine()
		err := NewFormError("Error.", e)

		assertEqual(t, err.Render(), `<div class="error">Error.</div>`)
	})

	t.Run("it sets message", func(t *testing.T) {
		e := render.NewEngine()
		err := NewFormError("Error.", e)

		assertEqual(t, err.Message(), "Error.")

		err.WithMessage("New Error.")

		assertEqual(t, err.Message(), "New Error.")
	})
}
