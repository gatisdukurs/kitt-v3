package htmx

import (
	"kitt/kitt/form"
	"kitt/kitt/render"
	"testing"
)

func Test_HTMX(t *testing.T) {
	e := render.NewEngine()
	e.WithTemplates("testdata/htmx.html")
	t.Run("it renders", func(t *testing.T) {
		view := render.NewView("htmx", e)
		viewHtmx := NewElement(view)

		have := viewHtmx.Render()
		want := `<div><span>Hello World!</span></div>`

		if have != want {
			t.Fatalf("Not equal: %s -> %s", have, want)
		}
	})

	t.Run("it renders with attributes", func(t *testing.T) {
		view := render.NewView("htmx", e)
		viewHtmx := NewElement(view)
		viewHtmx.WithId("navigation")
		viewHtmx.WithSwap(HTMX_SWAP_DEFAULT)

		have := viewHtmx.Render()
		want := `<div id="navigation" hx-swap-oob="true"><span>Hello World!</span></div>`

		if have != want {
			t.Fatalf("Not equal: %s -> %s", have, want)
		}
	})

	t.Run("it adds attributes to form", func(t *testing.T) {
		e := render.NewEngine()
		form := form.NewForm("login", e)

		form = HTMXForm(form)

		have := form.Render()
		want := `<form class="form" action="/" hx-post="/" hx-swap="outerHTML" id="login" method="POST"></form>`

		if have != want {
			t.Fatalf("not equal: %s -> %s", have, want)
		}
	})
}
