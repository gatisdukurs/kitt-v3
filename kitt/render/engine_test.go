package render

import (
	"testing"
)

func Test_Engine(t *testing.T) {
	t.Run("it adds templates", func(t *testing.T) {
		e := NewEngine()
		e.WithTemplates("testdata/engine.html")
	})

	t.Run("allows to add string templates", func(t *testing.T) {
		e := NewEngine()
		tpl := "<h1>Hello World!</h1>"
		e.WithTemplate("partial", tpl)
		buf := newBuf()
		e.Render(buf, "partial", nil)
		assertEqual(t, getBufStr(buf), tpl)
	})

	t.Run("it adds funcs", func(t *testing.T) {
		e := NewEngine()
		e.WithFuncs(Funcs{
			"hw": func() string {
				return "Hello World!"
			},
		})
		e.WithTemplate("partial", "{{ hw }}")
		buf := newBuf()
		e.Render(buf, "partial", nil)
		assertEqual(t, getBufStr(buf), "Hello World!")
	})

	t.Run("is rendering", func(t *testing.T) {
		e := NewEngine()
		e.WithTemplates("testdata/engine.html")
		buf := newBuf()
		err := e.Render(buf, "engine", nil)

		assertNoError(t, err)

		if buf.Len() == 0 {
			t.Fatal("no output")
		}

		assertEqual(t, getBufStr(buf), getSnap(t, "engine"))
	})
}
