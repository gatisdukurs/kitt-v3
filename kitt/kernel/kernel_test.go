package kernel

import (
	"testing"
)

func Test_Kitt(t *testing.T) {
	t.Run("it supports modules", func(t *testing.T) {
		k := NewKernel()
		app := NewFakeApp()
		k.WithApp(app)
		k.Boot()

		assertEqual(t, app.Booted, true)
	})

}
