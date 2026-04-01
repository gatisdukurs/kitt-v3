package kernel

import (
	"context"
	"kitt/kitt/services"
	"testing"
	"time"
)

func Test_Kitt(t *testing.T) {
	t.Run("it supports apps", func(t *testing.T) {
		k := NewKernel()
		app := NewFakeApp()
		k.WithApp(app)
		k.Boot()

		assertEqual(t, app.Booted, true)
	})

	t.Run("it supports services", func(t *testing.T) {
		c := services.NewContainer()
		k := NewKernel()
		k.WithServices(c)

		have := k.Services()
		want := c

		if have != want {
			t.Fatalf("!eq %s -> %s", have, want)
		}
	})

	t.Run("it supports runnables", func(t *testing.T) {
		r := &fakeRunnable{}
		k := NewKernel()
		k.WithRunnable(r)
		k.Boot()

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		go func(cancel context.CancelFunc) {
			time.Sleep(time.Millisecond * 200)
			cancel()
		}(cancel)

		err := k.Run(ctx)

		if err != nil {
			t.Fatal(err)
		}

		if !r.DidRun {
			t.Fatal("it did not run")
		}
	})

}
