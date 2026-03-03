package form

import (
	"kitt/kitt/render"
	"net/http"
	"net/url"
	"testing"
)

func Test_Form(t *testing.T) {
	t.Run("it renders", func(t *testing.T) {
		engine := render.NewEngine()
		f := NewForm("pages", engine)
		f.WithMethod(http.MethodGet)
		f.WithAction("/pages")
		control := NewFormControl("title", engine)
		field := NewFormField("title", engine)
		label := NewFormLabel("Title", engine)
		control.WithLabel(label)
		control.WithField(field)
		f.WithControl(control)

		assertEqual(t, f.Render(), `<form class="form" action="/pages" method="GET" id="pages"><div class="control" id="title"><label class="label">Title</label><input class="field" name="title" id="title" type="text" value="" /></div></form>`)
	})

	t.Run("it returns control", func(t *testing.T) {
		engine := render.NewEngine()
		f := NewForm("pages", engine)

		control := NewFormControl("title", engine)
		field := NewFormField("title", engine)
		label := NewFormLabel("Title", engine)
		control.WithLabel(label)
		control.WithField(field)
		f.WithControl(control)

		assertEqual(t, f.Control("title").Label().Name(), label.Name())
	})

	t.Run("it renders error", func(t *testing.T) {
		engine := render.NewEngine()
		f := NewForm("pages", engine)
		err := NewFormError("Error.", engine)

		assertEqual(t, f.RenderError(), "")

		f.WithError(err)

		assertEqual(t, f.RenderError(), err.Render())
		assertEqual(t, f.Render(), `<form class="form" action="/" method="POST" id="pages"><div class="error">Error.</div></form>`)
	})

	t.Run("it renders success", func(t *testing.T) {
		engine := render.NewEngine()
		f := NewForm("pages", engine)
		succ := NewFormSuccess("Success.", engine)

		assertEqual(t, f.RenderSuccess(), "")

		f.WithSuccess(succ)

		assertEqual(t, f.RenderSuccess(), succ.Render())
		assertEqual(t, f.Render(), `<form class="form" action="/" method="POST" id="pages"><div class="success">Success.</div></form>`)
	})

	t.Run("it sets values", func(t *testing.T) {
		email := "gatis.dukurs@gmail.com"
		password := "secret"
		engine := render.NewEngine()
		values := url.Values{}
		values.Set("email", email)
		values.Set("password", password)

		f := NewForm("pages", engine)

		control := NewFormControl("email", engine)
		field := NewFormField("email", engine)
		label := NewFormLabel("E-mail", engine)
		control.WithLabel(label)
		control.WithField(field)
		f.WithControl(control)

		control1 := NewFormControl("password", engine)
		field1 := NewFormField("password", engine)
		label1 := NewFormLabel("Password", engine)
		control1.WithLabel(label1)
		control1.WithField(field1)
		f.WithControl(control1)

		f.WithValues(values)

		assertEqual(t, field1.Value(), password)
		assertEqual(t, field.Value(), email)
	})

	t.Run("it sets error message", func(t *testing.T) {
		msg := "Passowrd or username is not correct."
		engine := render.NewEngine()
		err := NewFormError(msg, engine)
		f := NewForm("pages", engine)
		f.WithError(err)

		assertEqual(t, f.Error().Message(), msg)
	})

	t.Run("it sets success message", func(t *testing.T) {
		msg := "Account Created."
		engine := render.NewEngine()
		succ := NewFormSuccess(msg, engine)
		f := NewForm("pages", engine)
		f.WithSuccess(succ)

		assertEqual(t, f.Success().Message(), msg)
	})

	t.Run("it sets id", func(t *testing.T) {
		engine := render.NewEngine()
		f := NewForm("pages", engine)

		assertEqual(t, f.Id(), "pages")
		f.WithId("newId")
		assertEqual(t, f.Id(), "newId")
	})

	t.Run("it sets method", func(t *testing.T) {
		engine := render.NewEngine()
		f := NewForm("pages", engine)

		assertEqual(t, f.Method(), http.MethodPost)
		f.WithMethod(http.MethodGet)
		assertEqual(t, f.Method(), http.MethodGet)
	})

	t.Run("it sets action", func(t *testing.T) {
		engine := render.NewEngine()
		f := NewForm("pages", engine)

		assertEqual(t, f.Action(), "/")
		f.WithAction("/pages")
		assertEqual(t, f.Action(), "/pages")
	})

	t.Run("it renders controls", func(t *testing.T) {
		// engine := render.NewEngine()
		// f := NewForm("pages", engine)
	})
}
