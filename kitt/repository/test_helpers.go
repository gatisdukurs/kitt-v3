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
	ID    int    `db:"id,pk,required"`
	Name  string `db:"name"`
	NoTag string
}

type testFakeDriver[ID interface{}] struct {
	InsertCalled bool
	InsertValues DriverValues
	UpdateValues DriverValues
	UpdateCalled bool
	DeleteCalled bool
	ByIDCalled   bool
	InsertID     ID
	UpdateID     ID
}

func (d *testFakeDriver[ID]) Insert(collection string, values DriverValues) (ID, error) {
	d.InsertCalled = true
	d.InsertValues = values
	return d.InsertID, nil
}

func (d *testFakeDriver[ID]) Update(collection string, values DriverValues, id ID) error {
	d.UpdateCalled = true
	d.UpdateValues = values
	d.UpdateID = id
	return nil
}

func (d *testFakeDriver[ID]) Delete(collection string, id ID) error {
	d.DeleteCalled = true
	return nil
}

func (d *testFakeDriver[ID]) ByID(collection string, id ID) (DriverValues, error) {
	d.ByIDCalled = true
	var zero = DriverValues{}

	zero["id"] = 10
	zero["name"] = "Gatis"

	return zero, nil
}

func NewTestFakeDriver[ID interface{}]() *testFakeDriver[ID] {
	return &testFakeDriver[ID]{
		InsertCalled: false,
		UpdateCalled: false,
		DeleteCalled: false,
		ByIDCalled:   false,
	}
}
