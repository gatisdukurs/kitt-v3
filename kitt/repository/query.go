package repository

type Query struct {
	Where   []Filter
	OrderBy []Order
	Limit   int
	Offset  int
}

type Filter struct {
	Field string
	Op    string
	Value any
}

type Order struct {
	Field string
	Desc  bool
}
