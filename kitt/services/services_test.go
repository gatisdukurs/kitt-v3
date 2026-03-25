package services

import (
	"kitt/kitt/router"
	"reflect"
	"strings"
	"testing"
)

func Test_Services(t *testing.T) {
	t.Run("it sets and gets services", func(t *testing.T) {
		have := router.NewRouter()
		container := NewServices()
		container.Set(have)
		rtr := strings.ToLower(reflect.TypeOf(have).String())

		want := container.Get(rtr)

		if have != want {
			t.Fatalf("!eq: %s -> %s", have, want)
		}

		rtr2 := GetService[router.Router](container)

		if have != rtr2 {
			t.Fatalf("!eq: %s -> %s", have, rtr2)
		}
	})

	t.Run("it sets with key too", func(t *testing.T) {
		want := router.NewRouter()
		container := NewServices()
		container.SetWithKey("router", want)

		have := GetServiceWithKey[router.Router]("router", container)

		if want != have {
			t.Fatalf("!eq: %s -> %s", have, want)
		}
	})
}
