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

		for _, r := range router.Routes() {
			if r == route {
				found = true
			}
		}

		if !found {
			t.Fatal("route not found")
		}
	})

	t.Run("it routes", func(t *testing.T) {
		handled404 := false
		handler404 := func(ctx RouteCtx) RouteResponse {
			handled404 = true

			return nil
		}
		router := NewRouter()
		router.With404(handler404)
		homeRouted := false
		aboutRouted := false
		home := NewRoute("/home")
		home.GET(func(ctx RouteCtx) RouteResponse {
			homeRouted = true
			return nil
		})

		about := NewRoute("/about")
		about.GET(func(ctx RouteCtx) RouteResponse {
			aboutRouted = true
			return nil
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

	t.Run("it can serve static files", func(t *testing.T) {
		router := NewRouter()
		static := NewStaticRoute("/assets", "./testdata")
		router.To(static)

		r := httptest.NewRequest("GET", "/assets/static.txt", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, r)

		if w.Body.String() != "static" {
			t.Fatalf("unexpected body: %q", w.Body.String())
		}
	})

	t.Run("it handles with route response if its returned", func(t *testing.T) {
		str := "Hello World!"
		router := NewRouter()
		home := NewRoute("/home")
		home.GET(func(ctx RouteCtx) RouteResponse {
			sendable := newFakeRenderable(str)
			response := NewRouteResponse(sendable)
			return response
		})
		router.To(home)

		r := httptest.NewRequest("GET", "/home", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, r)

		if w.Body.String() != str {
			t.Fatalf("unexpected body: %q", w.Body.String())
		}
	})
}
