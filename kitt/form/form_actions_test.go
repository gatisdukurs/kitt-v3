package form

import (
	"kitt/kitt/render"
	"testing"
)

func Test_Form_Actions(t *testing.T) {
	t.Run("it renders", func(t *testing.T) {
		e := render.NewEngine()
		fa := NewFormActions("form-actions", e)
		action := NewFormAction("button", e)
		fa.WithAction(action)

		assertEqual(t, fa.Render(), `<div class="actions" id="form-actions"><button class="btn" id="button" name="button" value="">button</button></div>`)
	})

	t.Run("it renders actions", func(t *testing.T) {
		e := render.NewEngine()
		fa := NewFormActions("form-actions", e)
		action := NewFormAction("button", e)
		fa.WithAction(action)

		action1 := NewFormAction("button1", e)
		fa.WithAction(action1)

		assertEqual(t, fa.RenderActions(), `<button class="btn" id="button" name="button" value="">button</button><button class="btn" id="button1" name="button1" value="">button1</button>`)
	})

	t.Run("it sets id", func(t *testing.T) {
		e := render.NewEngine()
		fa := NewFormActions("form-actions", e)

		assertEqual(t, fa.Id(), "form-actions")
	})

	t.Run("it sets actions", func(t *testing.T) {
		e := render.NewEngine()
		fa := NewFormActions("form-actions", e)
		a := NewFormAction("button", e)

		actions := fa.Actions()
		assertEqual(t, len(actions), 0)

		fa.WithAction(a)
		actions = fa.Actions()
		assertEqual(t, len(actions), 1)
	})
}
