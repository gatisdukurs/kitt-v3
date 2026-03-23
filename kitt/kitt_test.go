package kitt

import (
	"context"
	"kitt/kitt/render"
	"kitt/kitt/router"
	"testing"
)

func Test_Kitt(t *testing.T) {
	t.Run("it provides Layout", func(t *testing.T) {
		K().InTesting()
		l := K().View("none")
		if _, ok := l.(render.View); !ok {
			t.Fatalf("not providing Layout")
		}
	})

	t.Run("it provides router", func(t *testing.T) {
		K().InTesting()
		r := K().Router()
		if _, ok := r.(router.Router); !ok {
			t.Fatalf("not providing router")
		}
	})

	t.Run("it allows to add templates", func(t *testing.T) {
		K().InTesting()
		K().WithTemplates(TemplatePatterns{
			"testdata/template.html",
		})
		l := K().View("template")
		assertEqual(t, l.Render(), getSnap(t, "template"))
	})

	t.Run("it allows to add string templates", func(t *testing.T) {
		K().InTesting()
		K().WithTemplate("partial", "<h1>Hello World!</h1>")
		p := K().View("partial")
		assertEqual(t, p.Render(), "<h1>Hello World!</h1>")
	})

	t.Run("it allows to add funcs", func(t *testing.T) {
		K().InTesting()
		K().WithTemplateFuncs(render.Funcs{
			"hw": func() string {
				return "World!"
			},
		})
		K().WithTemplate("partial", "<h1>Hello {{ hw }}</h1>")
		p := K().View("partial")
		assertEqual(t, p.Render(), "<h1>Hello World!</h1>")
	})

	t.Run("it serves", func(t *testing.T) {
		K().InTesting()
		addr := ":3000"
		fakeServer := newFakeHttpServer()
		K().WithHttpServer(fakeServer)
		err := K().ServeHttp(context.Background(), addr)

		assertNoError(t, err)
		assertEqual(t, fakeServer.Addr, addr)
		assertEqual(t, fakeServer.Handler, K().Router())
	})

	t.Run("it serves and returns error", func(t *testing.T) {
		K().InTesting()
		addr := ":3000"
		fakeServer := newFakeHttpServer()
		fakeServer.Error = true
		K().WithHttpServer(fakeServer)
		err := K().ServeHttp(context.Background(), addr)

		assertError(t, err)
	})

	t.Run("it supports modules", func(t *testing.T) {
		k := K()
		m := NewFakeModule()
		k.WithModule(m)
		k.Boot()

		assertEqual(t, m.Booted, true)
	})

}
