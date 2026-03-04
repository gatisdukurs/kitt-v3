package render

import "testing"

func Test_Layout(t *testing.T) {
	e := NewEngine()
	e.WithTemplates("testdata/layout.html")
	e.WithTemplates("testdata/layout_partials.html")
	e.WithTemplates("testdata/partial.html")
	e.WithTemplates("testdata/layout_ctx.html")
	e.WithTemplates("testdata/layout_htmx.html")
	e.WithTemplates("testdata/htmx_content.html")
	e.WithTemplates("testdata/htmx_navigation.html")

	t.Run("it renders", func(t *testing.T) {
		l := NewView("layout", e)
		assertNotNil(t, l)
		assertEqual(t, l.Render(), getSnap(t, "layout"))
	})

	t.Run("it renders partials", func(t *testing.T) {
		p1 := NewView("partial", e)
		p2 := NewView("partial", e)
		l := NewView("layout.partials", e)
		l.WithPartial("content", p1)
		// Skips this one cause there is no slot for this
		l.WithPartial("navigation", p2)

		assertEqual(t, l.Render(), getSnap(t, "layout_partials"))
	})

	t.Run("it renders without partials", func(t *testing.T) {
		l := NewView("layout.partials", e)
		assertEqual(t, l.Render(), getSnap(t, "layout_partials_without"))
	})

	t.Run("it renders with context", func(t *testing.T) {
		ctx := make(AnyCtx)
		ctx["title"] = "Hello World!"

		l := NewView("layout.ctx", e)
		l.WithCtx(ctx)

		assertEqual(t, l.Render(), getSnap(t, "layout_ctx"))
	})

	t.Run("it renders with no ctx", func(t *testing.T) {
		l := NewView("layout.ctx", e)
		assertEqual(t, l.Render(), getSnap(t, "layout_ctx_without"))
	})

	t.Run("it supports HTMX", func(t *testing.T) {
		// Add HTMX support in LAYOUT
		content := NewView("htmx.content", e)
		navigation := NewView("htmx.navigation", e)
		l := NewView("layout.htmx", e)
		l.WithPartial("content", content)
		l.WithPartial("navigation", navigation)

		l.WithHTMX("content", "navigation")

		assertEqual(t, l.HTMX(), getSnap(t, "layout_htmx"))
	})
}
