package domain

type Field struct {
	Name *string
	Type *string

	Values []any
}
type Table struct {
	Name   string
	Fields []Field
}

type Database struct {
	Tables *[]Table
}
