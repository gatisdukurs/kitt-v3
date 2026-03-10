package repository

type DriverValues = map[string]interface{}

type Driver[ID interface{}] interface {
	Insert(values DriverValues) (ID, error)
	Update(values DriverValues, id ID) error
	Delete(id ID) error
	ByID(id ID) (DriverValues, error)
	DropCollection() error
	CreateCollection() error
	WithModelMeta(modelMeta ModelMeta) Driver[ID]
}
