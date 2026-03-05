package repository

type Repository[T interface{}, ID interface{}] interface {
	Create(m *T) (ID, error)
	ByID(id ID) (T, error)
	Update(m *T) error
	Delete(id ID) error
}

type repo[T interface{}, ID interface{}] struct {
	driver Driver[ID]
}

func (r repo[T, ID]) Create(m *T) (ID, error) {
	var zero ID
	return zero, nil
}

func (r repo[T, ID]) ByID(id ID) (T, error) {
	var zero T
	return zero, nil
}

func (r repo[T, ID]) Update(m *T) error {
	return nil
}

func (r repo[T, ID]) Delete(id ID) error {
	return nil
}

func NewRepo[T interface{}, ID interface{}](driver Driver[ID]) Repository[T, ID] {
	return &repo[T, ID]{}
}

// driver := sqlite.NewDriver[User](kitt.SQL(), "users")
// users := repo.New[User](driver)
