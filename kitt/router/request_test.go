package router

import (
	"net/http"
	"net/http/httptest"
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
}
