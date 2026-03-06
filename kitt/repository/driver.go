package repository

type DriverValues = map[string]interface{}

type Driver[ID interface{}] interface {
	Insert(collection string, values DriverValues) (ID, error)
	Update(collection string, values DriverValues, id ID) error
	Delete(collection string, id ID) error
	ByID(collection string, id ID) (DriverValues, error)
}
