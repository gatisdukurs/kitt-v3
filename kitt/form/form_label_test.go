package form

import (
	"kitt/kitt/render"
	"testing"
)

func Test_Form_Label(t *testing.T) {
	t.Run("it renders", func(t *testing.T) {
		e := render.NewEngine()
		label := NewLabel("E-mail", e)
		assertEqual(t, label.Render(), `<label>E-mail</label>`)
	})

}
