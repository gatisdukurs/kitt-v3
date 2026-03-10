package repository

import (
	"fmt"
	"strings"
)

type QueryBuilder interface {
	Select(fields ...string) QueryBuilder
	From(table string) QueryBuilder
	Where(field string, op string, value any) QueryBuilder
	OrderAscBy(field string) QueryBuilder
	OrderDescBy(field string) QueryBuilder
	Limit(limit int) QueryBuilder
	Offset(offset int) QueryBuilder
	Build() (string, []any)
}

type Query struct {
	Fields  []string
	Table   string
	Where   []Condition
	OrderBy []Order
	Limit   int
	Offset  int
}

type Condition struct {
	Field string
	Op    string
	Value any
}

type Order struct {
	Field string
	Desc  bool
}

type builder struct {
	query       Query
	queryString string
	queryArgs   []any
}

func (b *builder) Select(fields ...string) QueryBuilder {
	b.query.Fields = fields
	return b
}

func (b *builder) From(table string) QueryBuilder {
	b.query.Table = table
	return b
}

func (b *builder) Where(field string, op string, value any) QueryBuilder {
	condition := Condition{
		Field: field,
		Op:    op,
		Value: value,
	}

	b.query.Where = append(b.query.Where, condition)

	return b
}

func (b *builder) OrderAscBy(field string) QueryBuilder {
	order := Order{
		Field: field,
		Desc:  false,
	}
	b.query.OrderBy = append(b.query.OrderBy, order)

	return b
}

func (b *builder) OrderDescBy(field string) QueryBuilder {
	order := Order{
		Field: field,
		Desc:  false,
	}
	b.query.OrderBy = append(b.query.OrderBy, order)

	return b
}

func (b *builder) Limit(limit int) QueryBuilder {
	b.query.Limit = limit
	return b
}

func (b *builder) Offset(offset int) QueryBuilder {
	b.query.Offset = offset
	return b
}

func (b builder) Build() (string, []any) {
	str, args, err := BuildQuery(b.query)

	if err != nil {
		panic(err)
	}

	return str, args
}

func NewQueryBuilder() QueryBuilder {
	return &builder{
		query:       Query{},
		queryString: "",
		queryArgs:   []any{},
	}
}

func BuildCondition(condition Condition) (string, any) {
	return fmt.Sprintf(`%s %s ?`, condition.Field, condition.Op), condition.Value
}

func BuildWhere(query Query) (string, []any) {
	conditions := query.Where
	where := []string{}
	args := []any{}

	for _, condition := range conditions {
		condition, arg := BuildCondition(condition)
		where = append(where, condition)
		args = append(args, arg)
	}

	return fmt.Sprintf(`WHERE %s`, strings.Join(where, " AND ")), args
}

func BuildOrderBy(query Query) string {
	orderBys := query.OrderBy
	conditions := []string{}

	for _, condition := range orderBys {
		order := "ASC"
		if condition.Desc {
			order = "DESC"
		}
		conditions = append(conditions, fmt.Sprintf(`%s %s`, condition.Field, order))
	}

	str := fmt.Sprintf(`ORDER BY %s`, strings.Join(conditions, ", "))

	return str
}

func BuildLimitOffset(query Query) string {
	limit := query.Limit
	offset := query.Offset

	if limit > 0 && offset > 0 {
		return fmt.Sprintf(`LIMIT %d OFFSET %d`, limit, offset)
	}

	if limit > 0 {
		return fmt.Sprintf(`LIMIT %d`, limit)
	}

	if offset > 0 {
		return fmt.Sprintf(`LIMIT -1 OFFSET %d`, offset)
	}

	return ""
}

func BuildQuery(query Query) (string, []any, error) {
	fields := "*"
	table := query.Table
	where := ""
	args := []any{}
	q := ""

	// Table
	if table == "" {
		return "", nil, fmt.Errorf("No table")
	}

	// Fields
	if len(query.Fields) > 0 {
		fields = strings.Join(query.Fields, ", ")
	}

	q = fmt.Sprintf(`SELECT %s FROM %s`, fields, table)

	// Where
	if len(query.Where) > 0 {
		where, args = BuildWhere(query)
		q += " " + where
	}

	// Order by
	if len(query.OrderBy) > 0 {
		orderBy := BuildOrderBy(query)
		q += " " + orderBy
	}

	// Limit, Offset
	limitOffset := BuildLimitOffset(query)

	if limitOffset != "" {
		q += " " + limitOffset
	}

	return q, args, nil
}
