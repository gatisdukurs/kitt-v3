package repository

type DriverValues = map[string]interface{}

const DRIVER_SQL = "sql"
const DRIVER_MONGO = "mongo"

type Driver[ID interface{}] interface {
	Insert(values DriverValues) (ID, error)
	Update(values DriverValues, id ID) error
	Delete(id ID) error

	ByID(id ID) (DriverValues, error)
	Find(query QueryBuilder) ([]DriverValues, error)
	First(query QueryBuilder) (DriverValues, error)

	DropCollection() error
	CreateCollection() error
	WithModelMeta(modelMeta ModelMeta) Driver[ID]
}
