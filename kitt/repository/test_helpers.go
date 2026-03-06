package repository

import (
	"bytes"
	"context"
	"database/sql"
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

type mockSqlConn struct {
}

func (c mockSqlConn) Exec(ctx context.Context, q string, args ...any) (sql.Result, error) {
	var zero sql.Result
	return zero, nil
}

func (c mockSqlConn) Query(ctx context.Context, q string, args ...any) (*sql.Rows, error) {
	var zero sql.Rows

	return &zero, nil
}

func (c mockSqlConn) QueryRow(ctx context.Context, q string, args ...any) *sql.Row {
	var zero sql.Row
	return &zero
}

func (c mockSqlConn) Close() error {
	return nil
}

func (c mockSqlConn) WithDB(path string) SqlConnection {
	return c
}

func NewMockSqlConnection() *mockSqlConn {
	return &mockSqlConn{}
}
