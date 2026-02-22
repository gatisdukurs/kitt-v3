package kitt

import (
	"kitt/kitt/render"
	"testing"
)

func Test_Kitt(t *testing.T) {
	t.Run("it provides Layout", func(t *testing.T) {
		K().InTesting()
		l := K().Layout("none")
		if _, ok := l.(render.Layout); !ok {
			t.Fatalf("not providing Layout")
		}
	})

	t.Run("it provides Partial", func(t *testing.T) {
		K().InTesting()
		p := K().Partial("none")
		if _, ok := p.(render.Partial); !ok {
			t.Fatalf("not providing Partial")
		}
	})

	t.Run("it allows to add templates", func(t *testing.T) {
		K().InTesting()
		K().WithTemplates(KittTemplatePatterns{
			"testdata/template.html",
		})
		l := K().Layout("template")
		assertEqual(t, l.Render(), getSnap(t, "template"))
	})

	t.Run("it allows to add string templates", func(t *testing.T) {
		K().InTesting()
		K().WithTemplate("partial", "<h1>Hello World!</h1>")
		p := K().Partial("partial")
		assertEqual(t, p.Render(), "<h1>Hello World!</h1>")
	})

	t.Run("it allows to add funcs", func(t *testing.T) {
		K().InTesting()
		K().WithTemplateFuncs(KittTemplateFuncs{
			"hw": func() string {
				return "World!"
			},
		})
		K().WithTemplate("partial", "<h1>Hello {{ hw }}</h1>")
		p := K().Partial("partial")
		assertEqual(t, p.Render(), "<h1>Hello World!</h1>")
	})

	t.Run("it provides basic context", func(t *testing.T) {
		K().InTesting()
		ctx := K().Ctx().Basic()
		ctx["foo"] = "bar"
		assertEqual(t, ctx["foo"], "bar")
	})
}
