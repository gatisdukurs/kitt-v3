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
		field := NewFormField("email", e)
		control := NewFormControl("field", e)
		label := NewFormLabel("E-mail", e)
		field.WithLabel(label)
		field.WithControl(control)
		form.WithField(field)

		form.WithAttribute("required", "")
		form.WithAttribute("autofocus", "true")

		form.WithError("Error.")

		form.WithSuccess("Success.")

		ctx := NewFormContext(form)

		assertEqual(t, ctx.Action(), form.Action())
		assertEqual(t, ctx.Method(), form.Method())
		assertEqual(t, ctx.Id(), form.Id())
		assertEqual(t, ctx.Fields(), render.AsHtml(form.RenderFields()))
		assertEqual(t, ctx.Error(), render.AsHtml(form.RenderError()))
		assertEqual(t, ctx.Success(), render.AsHtml(form.RenderSuccess()))
		assertEqual(t, ctx.Actions(), render.AsHtml(form.RenderActions()))
		assertEqual(t, ctx.Attributes(), render.AsAttr(form.RenderAttributes()))
	})
}
