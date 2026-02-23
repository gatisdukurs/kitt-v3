package render

import "testing"

func Test_Layout(t *testing.T) {
	e := NewEngine()
	e.WithTemplates("testdata/layout.html")
	e.WithTemplates("testdata/layout_partials.html")
	e.WithTemplates("testdata/partial.html")
	e.WithTemplates("testdata/layout_ctx.html")

	t.Run("it renders", func(t *testing.T) {
		l := NewLayout("layout", e)
		assertNotNil(t, l)
		assertEqual(t, l.Render(), getSnap(t, "layout"))
	})

	t.Run("it renders partials", func(t *testing.T) {
		p1 := NewPartial("partial", e)
		p2 := NewPartial("partial", e)
		l := NewLayout("layout.partials", e)
		l.WithPartial("content", p1)
		// Skips this one cause there is no slot for this
		l.WithPartial("navigation", p2)

		assertEqual(t, l.Render(), getSnap(t, "layout_partials"))
	})

	t.Run("it renders without partials", func(t *testing.T) {
		l := NewLayout("layout.partials", e)
		assertEqual(t, l.Render(), getSnap(t, "layout_partials_without"))
	})

	t.Run("it renders with context", func(t *testing.T) {
		ctx := make(AnyCtx)
		ctx["title"] = "Hello World!"

		l := NewLayout("layout.ctx", e)
		l.WithCtx(ctx)

		assertEqual(t, l.Render(), getSnap(t, "layout_ctx"))
	})

	t.Run("it renders with no ctx", func(t *testing.T) {
		l := NewLayout("layout.ctx", e)
		assertEqual(t, l.Render(), getSnap(t, "layout_ctx_without"))
	})
}
