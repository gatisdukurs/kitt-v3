package router

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func Test_Request(t *testing.T) {
	t.Run("it returns path", func(t *testing.T) {
		path := "/home"
		httpRequest := httptest.NewRequest(http.MethodGet, path, nil)
		request := NewRequest()
		request.WithHttpRequest(httpRequest)

		assertEqual(t, path, request.Path())
	})

	t.Run("it returens for values", func(t *testing.T) {
		email := "test@example.com"
		password := "secret"

		form := url.Values{}
		form.Set("email", email)
		form.Set("password", password)

		httpRequest := httptest.NewRequest(http.MethodPost, "/post", strings.NewReader(form.Encode()))
		httpRequest.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		request := NewRequest()
		request.WithHttpRequest(httpRequest)

		values := request.FormValues()

		assertEqual(t, values.Get("email"), email)
	})

	t.Run("it returns http request", func(t *testing.T) {
		path := "/home"
		httpRequest := httptest.NewRequest(http.MethodGet, path, nil)
		request := NewRequest()
		request.WithHttpRequest(httpRequest)

		assertEqual(t, request.HttpRequest(), httpRequest)
	})

	t.Run("it detects HTMX", func(t *testing.T) {
		path := "/home"
		httpRequest := httptest.NewRequest(http.MethodGet, path, nil)
		httpRequest.Header.Set("HX-Request", "true")
		request := NewRequest()
		request.WithHttpRequest(httpRequest)

		assertEqual(t, request.HTMX(), true)

		httpRequest = httptest.NewRequest(http.MethodGet, path, nil)
		request = NewRequest()
		request.WithHttpRequest(httpRequest)

		assertEqual(t, request.HTMX(), false)
	})
}
