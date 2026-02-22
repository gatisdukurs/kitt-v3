package kitt

import (
	"bytes"
	"context"
	"fmt"
	"kitt/kitt/router"
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

func assertError(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Fatal("should throw error")
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

type fakeHttpServer struct {
	Error          bool
	Addr           string
	Handler        router.HttpHandler
	ShutdownCalled bool
}

func (f *fakeHttpServer) ListenAndServe(ctx context.Context, addr string, handler router.HttpHandler) error {
	f.Addr = addr
	f.Handler = handler

	if f.Error {
		return fmt.Errorf("fake error")
	}
	return nil
}

func (f *fakeHttpServer) Shutdown() error {
	f.ShutdownCalled = true
	if f.Error {
		return fmt.Errorf("fake error")
	}
	return nil
}

func newFakeHttpServer() *fakeHttpServer {
	return &fakeHttpServer{}
}
