package repository

type sqliteDriver[ID interface{}] struct{}

func (d sqliteDriver[ID]) Insert(collection string, values DriverValues) (ID, error) {
	var zero ID
	return zero, nil
}

func (d sqliteDriver[ID]) Update(collection string, values DriverValues, id ID) error {
	return nil
}

func (d sqliteDriver[ID]) Delete(collection string, id ID) error {
	return nil
}

func (d sqliteDriver[ID]) ByID(collection string, id ID) (DriverValues, error) {
	var zero DriverValues
	return zero, nil
}

func NewSqliteDriver[ID interface{}]() Driver[ID] {
	return &sqliteDriver[ID]{}
}
