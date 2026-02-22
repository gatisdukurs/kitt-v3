package router

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Router(t *testing.T) {
	t.Run("it adds routes", func(t *testing.T) {
		router := NewRouter()
		route := NewRoute("/home")

		router.To(route)
		found := false

		for pattern, r := range router.Routes() {
			if r == route {
				found = true
				assertEqual(t, r.Pattern(), pattern)
			}
		}

		if !found {
			t.Fatal("route not found")
		}
	})

	t.Run("it routes", func(t *testing.T) {
		handled404 := false
		handler404 := func(ctx RouteCtx) {
			handled404 = true
		}
		router := NewRouter()
		router.With404(handler404)
		homeRouted := false
		aboutRouted := false
		home := NewRoute("/home")
		home.GET(func(ctx RouteCtx) {
			homeRouted = true
		})

		about := NewRoute("/about")
		about.GET(func(ctx RouteCtx) {
			aboutRouted = true
		})

		router.
			To(home).
			To(about)

		r := httptest.NewRequest("GET", "/home", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, r)

		assertEqual(t, homeRouted, true)

		r = httptest.NewRequest("GET", "/about", nil)
		w = httptest.NewRecorder()
		router.ServeHTTP(w, r)

		assertEqual(t, aboutRouted, true)

		r = httptest.NewRequest("GET", "/404", nil)
		w = httptest.NewRecorder()
		router.ServeHTTP(w, r)

		assertEqual(t, handled404, true)

		// Check for regular 404
		router.With404(nil)

		r = httptest.NewRequest("GET", "/404", nil)
		w = httptest.NewRecorder()
		router.ServeHTTP(w, r)

		if w.Code != http.StatusNotFound {
			t.Fatalf("expected status 404, got %d", w.Code)
		}

		if w.Body.String() != "404 page not found\n" {
			t.Fatalf("unexpected body: %q", w.Body.String())
		}
	})
}
