package repository

type ID = int64

type Repository[T interface{}] interface {
	Create(m *T) (ID, error)
	ByID(id ID) (T, error)
	Update(m *T) error
	Delete(id ID) error
}

type repo[T interface{}] struct {
}

func (r repo[T]) Create(m *T) (ID, error) {
	return 0, nil
}

func (r repo[T]) ByID(id ID) (T, error) {
	var zero T
	return zero, nil
}

func (r repo[T]) Update(m *T) error {
	return nil
}

func (r repo[T]) Delete(id ID) error {
	return nil
}

func NewRepo[T interface{}]() Repository[T] {
	return &repo[T]{}
}
