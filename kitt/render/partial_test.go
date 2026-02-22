package render

import "testing"

func Test_Partial(t *testing.T) {
	e := NewEngine()
	e.WithTemplates("testdata/partial.html")
	e.WithTemplates("testdata/partial_ctx.html")

	t.Run("it renders", func(t *testing.T) {
		p := NewPartial("partial", e)
		assertEqual(t, p.Render(), getSnap(t, "partial"))
	})

	t.Run("it renders ctx", func(t *testing.T) {
		ctx := make(AnyCtx)
		ctx["world"] = "World!"
		p := NewPartial("partial.ctx", e)
		p.WithCtx(ctx)
		assertEqual(t, p.Render(), getSnap(t, "partial_ctx"))
	})

	t.Run("it renders without ctx", func(t *testing.T) {
		p := NewPartial("partial.ctx", e)
		assertEqual(t, p.Render(), getSnap(t, "partial_ctx_without"))
	})
}
