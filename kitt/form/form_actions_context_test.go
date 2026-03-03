package form

import (
	"kitt/kitt/render"
	"testing"
)

func Test_Form_Actions_Context(t *testing.T) {
	t.Run("it provides values", func(t *testing.T) {
		e := render.NewEngine()
		fa := NewFormActions("actions", e)
		action := NewFormAction("button", e)
		fa.WithAction(action)
		ctx := NewFormActionsContext(fa)

		assertEqual(t, ctx.Id(), fa.Id())
		assertEqual(t, ctx.Actions(), render.AsHtml(fa.RenderActions()))
	})
}
