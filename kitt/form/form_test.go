package form

import (
	"kitt/kitt/render"
	"net/http"
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
