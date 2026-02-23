package domain

type StorableType int

const (
	Int StorableType = iota
	String
)

type Column struct {
	Name string
	Type StorableType
}

type Row []any

type Table struct {
	Name    string
	Columns []Column
	Rows    []Row
}

type Database struct {
	Name   string
	Tables map[string]*Table
}
