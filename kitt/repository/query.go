package repository

import (
	"fmt"
	"strings"
)

type Sqlizer interface {
	Build() (string, []any)
}

type Clause interface {
	Build() (string, []any)
}

const (
	ASC  = "ASC"
	DESC = "DESC"
)

// Query Builders
// SELECT
type SelectBuilder struct {
	Sqlizer

	table   string
	columns []string

	orderBy []string
	limit   *int
	offset  *int

	where Clause
}

func (sb *SelectBuilder) Columns(columns ...string) *SelectBuilder {
	sb.columns = columns
	return sb
}

func (sb *SelectBuilder) Where(where Clause) *SelectBuilder {
	sb.where = where
	return sb
}

func (sb *SelectBuilder) OrderBy(column string, direction string) *SelectBuilder {
	sb.orderBy = append(sb.orderBy, fmt.Sprintf(`%s %s`, column, direction))
	return sb
}

func (sb *SelectBuilder) Limit(limit int) *SelectBuilder {
	sb.limit = &limit
	return sb
}

func (sb *SelectBuilder) Offset(offset int) *SelectBuilder {
	sb.offset = &offset
	return sb
}

func (sb SelectBuilder) Build() (string, []any) {
	args := []any{}
	query := "SELECT"

	// Columns
	query += BuildColumns(sb.columns)

	// From
	query += BuildFrom(sb.table)

	// Where
	if sb.where != nil {
		where, whereArgs := sb.where.Build()
		query += fmt.Sprintf(" WHERE %s", where)

		args = append(args, whereArgs...)
	}

	// Order by
	if len(sb.orderBy) > 0 {
		query += fmt.Sprintf(" ORDER BY %s", strings.Join(sb.orderBy, ", "))
	}

	// Limit
	if sb.limit != nil {
		query += fmt.Sprintf(" LIMIT %d", *sb.limit)
	}

	// Offset
	if sb.offset != nil {
		query += fmt.Sprintf(" OFFSET %d", *sb.offset)
	}

	return query, args
}

func SELECT(table string) *SelectBuilder {
	return &SelectBuilder{
		table: table,
	}
}

// INSERT
type InsertBuilder struct {
	Sqlizer

	table   string
	columns []string
	rows    [][]any
}

func (ib *InsertBuilder) Columns(columns ...string) *InsertBuilder {
	ib.columns = columns
	return ib
}

func (ib *InsertBuilder) Row(values ...any) *InsertBuilder {
	if len(values) != len(ib.columns) {
		panic("columns != rows")
	}

	ib.rows = append(ib.rows, values)
	return ib
}

func (ib *InsertBuilder) Build() (string, []any) {
	args := []any{}

	query := fmt.Sprintf("INSERT INTO %s", ib.table)
	query += fmt.Sprintf(" (%s)", strings.Join(ib.columns, ", "))

	query += " VALUES "

	placeholders := make([]string, len(ib.columns))
	for i := range placeholders {
		placeholders[i] = "?"
	}

	valuePlaceholders := fmt.Sprintf(`(%s)`, strings.Join(placeholders, ", "))
	valueRows := []string{}

	if len(ib.rows) == 0 {
		valueRows = append(valueRows, valuePlaceholders)
	}

	for _, row := range ib.rows {
		valueRows = append(valueRows, valuePlaceholders)
		args = append(args, row...)
	}

	query += strings.Join(valueRows, ", ")

	return query, args
}

func INSERT(table string) *InsertBuilder {
	return &InsertBuilder{
		table: table,
	}
}

// UPDATE
type UpdateBuilder struct {
	Sqlizer

	table   string
	columns []string
	values  []any
	where   Clause
}

func (ub *UpdateBuilder) Set(column string, value any) *UpdateBuilder {
	ub.columns = append(ub.columns, column)
	ub.values = append(ub.values, value)

	return ub
}

func (ub *UpdateBuilder) Where(where Clause) *UpdateBuilder {
	ub.where = where

	return ub
}

func (ub UpdateBuilder) Build() (string, []any) {
	args := []any{}
	query := fmt.Sprintf("UPDATE %s", ub.table)

	// Set
	set := []string{}
	for i, column := range ub.columns {
		set = append(set, fmt.Sprintf(`%s = ?`, column))
		args = append(args, ub.values[i])
	}
	query += " SET " + strings.Join(set, ", ")

	// Where
	if ub.where == nil {
		panic("no where clause")
	}
	where, whereArgs := ub.where.Build()
	query += fmt.Sprintf(" WHERE %s", where)
	args = append(args, whereArgs...)

	return query, args
}

func UPDATE(table string) *UpdateBuilder {
	return &UpdateBuilder{
		table: table,
	}
}

// DELETE
type DeleteBuilder struct {
	Sqlizer

	table string

	where Clause
}

func (db *DeleteBuilder) Where(c Clause) *DeleteBuilder {
	db.where = c
	return db
}

func (db DeleteBuilder) Build() (string, []any) {
	args := []any{}
	query := "DELETE"
	query += BuildFrom(db.table)

	if db.where == nil {
		panic("no where clause")
	}

	where, args := db.where.Build()
	query += fmt.Sprintf(" WHERE %s", where)

	return query, args
}

func DELETE(table string) *DeleteBuilder {
	return &DeleteBuilder{
		table: table,
	}
}

// Build Functions
func BuildColumns(columns []string) string {
	if len(columns) == 0 {
		return " *"
	}

	return " " + strings.Join(columns, ", ")
}

func BuildFrom(table string) string {
	return " FROM " + table
}

// Ops
type GroupClause struct {
	operator string
	clauses  []Clause
}

func And(clauses ...Clause) Clause {
	return GroupClause{
		operator: "AND",
		clauses:  clauses,
	}
}

func Or(clauses ...Clause) Clause {
	return GroupClause{
		operator: "OR",
		clauses:  clauses,
	}
}

func (g GroupClause) Build() (string, []any) {
	var parts []string
	var args []any

	for _, clause := range g.clauses {
		if clause == nil {
			continue
		}

		sql, clauseArgs := clause.Build()
		if sql == "" {
			continue
		}

		parts = append(parts, sql)
		args = append(args, clauseArgs...)
	}

	if len(parts) == 0 {
		return "", nil
	}

	if len(parts) == 1 {
		return parts[0], args
	}

	return "(" + strings.Join(parts, " "+g.operator+" ") + ")", args
}

type ExprClause struct {
	sql  string
	args []any
}

func Expr(sql string, args ...any) Clause {
	return ExprClause{
		sql:  sql,
		args: args,
	}
}

func (e ExprClause) Build() (string, []any) {
	return e.sql, e.args
}

func Eq(column string, value any) Clause {
	return Expr(fmt.Sprintf("%s = ?", column), value)
}

func Ne(column string, value any) Clause {
	return Expr(fmt.Sprintf("%s != ?", column), value)
}

func Gt(column string, value any) Clause {
	return Expr(fmt.Sprintf("%s > ?", column), value)
}

func Gte(column string, value any) Clause {
	return Expr(fmt.Sprintf("%s >= ?", column), value)
}

func Lt(column string, value any) Clause {
	return Expr(fmt.Sprintf("%s < ?", column), value)
}

func Lte(column string, value any) Clause {
	return Expr(fmt.Sprintf("%s <= ?", column), value)
}

func Like(column string, value any) Clause {
	return Expr(fmt.Sprintf("%s LIKE ?", column), value)
}

func IsNull(column string) Clause {
	return Expr(fmt.Sprintf("%s IS NULL", column))
}

func IsNotNull(column string) Clause {
	return Expr(fmt.Sprintf("%s IS NOT NULL", column))
}
