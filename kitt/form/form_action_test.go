package form

import (
	"kitt/kitt/render"
	"testing"
)

func Test_Form_Action(t *testing.T) {
	t.Run("it renders", func(t *testing.T) {
		e := render.NewEngine()
		fa := NewFormAction("submit", e)

		fa.WithId("action_id")
		fa.WithLabel("Save")
		fa.WithValue("save")

		assertEqual(t, fa.Render(), `<button class="btn" id="action_id" name="submit" value="save">Save</button>`)
	})

	t.Run("it sets id", func(t *testing.T) {
		e := render.NewEngine()
		fa := NewFormAction("submit", e)

		assertEqual(t, fa.Id(), "submit")

		fa.WithId("new_id")

		assertEqual(t, fa.Id(), "new_id")
	})

	t.Run("it sets name", func(t *testing.T) {
		e := render.NewEngine()
		fa := NewFormAction("submit", e)

		assertEqual(t, fa.Name(), "submit")
	})

	t.Run("it sets label", func(t *testing.T) {
		e := render.NewEngine()
		fa := NewFormAction("submit", e)
		assertEqual(t, fa.Label(), "submit")

		fa.WithLabel("new_label")
		assertEqual(t, fa.Label(), "new_label")
	})

	t.Run("it sets value", func(t *testing.T) {
		e := render.NewEngine()
		fa := NewFormAction("submit", e)
		fa.WithValue("value")

		assertEqual(t, fa.Value(), "value")
	})
}
