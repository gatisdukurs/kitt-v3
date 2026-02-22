package render

import (
	"bytes"
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
