package render

import "testing"

func Test_View_Stack(t *testing.T) {
	e := NewEngine()
	e.WithTemplates("testdata/partial.html")

	t.Run("it renders empty", func(t *testing.T) {
		stack := NewViewStack()
		assertEqual(t, stack.Render(), "")
	})

	t.Run("it renders stack", func(t *testing.T) {
		v1 := NewView("partial", e)
		v2 := NewView("partial", e)
		stack := NewViewStack()
		stack.WithRenderable(v1)
		stack.WithRenderable(v2)

		assertEqual(t, stack.Render(), `<h1>Partial</h1><h1>Partial</h1>`)
	})
}
