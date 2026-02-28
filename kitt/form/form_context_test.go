package form

import (
	"kitt/kitt/render"
	"net/http"
	"testing"
)

func Test_Form_Context(t *testing.T) {
	t.Run("it returns values", func(t *testing.T) {
		e := render.NewEngine()
		form := NewForm("login", e)
		form.WithAction("/action")
		form.WithMethod(http.MethodGet)
		form.WithId("login-form")
		control := NewFormControl("email", e)
		field := NewFormField("field", e)
		label := NewFormLabel("E-mail", e)
		control.WithLabel(label)
		control.WithField(field)
		form.WithControl(control)

		ctx := NewFormContext(form)

		assertEqual(t, ctx.Action(), form.Action())
		assertEqual(t, ctx.Method(), form.Method())
		assertEqual(t, ctx.Id(), form.Id())
		assertEqual(t, ctx.Controls(), render.AsHtml(form.RenderControls()))
	})
}
