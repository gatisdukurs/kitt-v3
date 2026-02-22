package kitt

import "testing"

func Test_Kitt_Context(t *testing.T) {
	t.Run("it sets values and returns basic context", func(t *testing.T) {
		kctx := NewKittCtx()
		kctx.Set("foo", "bar")

		ctx := kctx.Basic()

		assertEqual(t, ctx["foo"], "bar")
	})
}
