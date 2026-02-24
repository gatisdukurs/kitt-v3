package router

import (
	"bytes"
	"net/http"
	"os"
	"strings"
	"testing"
)

func newBuf() *bytes.Buffer {
	var buf bytes.Buffer
	return &buf
}

func getBufStr(buf *bytes.Buffer) string {
	return strings.TrimSpace(buf.String())
}

func assertEqual(t *testing.T, have interface{}, want interface{}) {
	if have != want {
		t.Fatal("not equal", have, "\n\n", want)
	}
}

func assertNotNil(t *testing.T, i interface{}) {
	if i == nil {
		t.Fatal("is nil")
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err.Error())
	}
}

func getSnap(t *testing.T, snap string) string {
	t.Helper()

	path := "testdata/" + snap + ".snap.html"

	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("failed to read file %q: %v", path, err)
	}

	return strings.TrimSpace(string(b))
}

// Transition to these later
// req := httptest.NewRequest("GET", "/", nil)
// w := httptest.NewRecorder()
// mux.ServeHTTP(w, req)
func newFakeResponseWriter() *responseWriterForTesting {
	return &responseWriterForTesting{
		Buf:    newBuf(),
		Status: 0,
	}
}

type responseWriterForTesting struct {
	Buf    *bytes.Buffer
	Status int
}

func (r responseWriterForTesting) Write(bytes []byte) (int, error) {
	return r.Buf.Write(bytes)
}

func (r responseWriterForTesting) Header() http.Header {
	return make(http.Header)
}

func (r *responseWriterForTesting) WriteHeader(status int) {
	r.Status = status
}

func (r responseWriterForTesting) Sent() string {
	return getBufStr(r.Buf)
}

type fakeHttpHandler struct {
	Served bool
}

func (fh *fakeHttpHandler) ServeHTTP(http.ResponseWriter, *http.Request) {
	fh.Served = true
}

func newFakeHttpHandler() *fakeHttpHandler {
	return &fakeHttpHandler{
		Served: false,
	}
}

type fakeRenderable struct {
	String string
}

func (fr fakeRenderable) Render() string {
	return fr.String
}

func (fr fakeRenderable) HTMX() string {
	return fr.String
}

func newFakeRenderable(str string) *fakeRenderable {
	return &fakeRenderable{
		String: str,
	}
}
