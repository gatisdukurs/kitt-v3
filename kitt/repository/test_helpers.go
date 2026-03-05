package repository

import (
	"bytes"
	"reflect"
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
	if !reflect.DeepEqual(have, want) {
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

type TestUser struct {
	Id   int64
	Name string
}

type testFakeDriver[ID interface{}] struct {
	InsertCalled bool
	UpdateCalled bool
	DeleteCalled bool
	ByIDCalled   bool
}

func (d *testFakeDriver[ID]) Insert(collection string, values DriverValues) error {
	d.InsertCalled = true
	return nil
}

func (d *testFakeDriver[ID]) Update(collection string, values DriverValues, id ID) error {
	d.UpdateCalled = true
	return nil
}

func (d *testFakeDriver[ID]) Delete(collection string, id ID) error {
	d.DeleteCalled = true
	return nil
}

func (d *testFakeDriver[ID]) ByID(collection string, id ID) {
	d.ByIDCalled = true
}

func NewTestFakeDriver[ID interface{}]() *testFakeDriver[ID] {
	return &testFakeDriver[ID]{
		InsertCalled: false,
		UpdateCalled: false,
		DeleteCalled: false,
		ByIDCalled:   false,
	}
}
