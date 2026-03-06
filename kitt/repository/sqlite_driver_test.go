package repository

import "testing"

func Test_Sqlite_Driver(t *testing.T) {
	t.Run("it inserts", func(t *testing.T) {
		conn := NewMockSqlConnection()
		driver := NewSqliteDriver[int](conn)
		values := DriverValues{}

		driver.Insert("users", values)
	})
}
