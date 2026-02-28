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

		assertEqual(t, f.Render(), `<form action="/pages" method="GET" id="pages"></form>`)
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
