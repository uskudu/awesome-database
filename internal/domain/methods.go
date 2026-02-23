package domain

import (
	"errors"
	"fmt"
	"os"
	"text/tabwriter"
)

type DB interface {
	CreateTable(name string, cols []Column) (*Table, error)
	PrintTable(name string, table *Table) error
}

func NewDatabase(name string) *Database {
	return &Database{
		Name:   name,
		Tables: make(map[string]*Table),
	}
}

func (db *Database) CreateTable(name string, cols []Column) (*Table, error) {
	if _, exists := db.Tables[name]; exists {
		return nil, errors.New("table already exists")
	}

	t := &Table{
		Name:    name,
		Columns: cols,
		Rows:    []Row{},
	}
	db.Tables[name] = t

	return t, nil
}

//func (db *Database) SelectRow(name string) (Row, error)
//func (db *Database) DeleteRow(name string) (string, error)

func (t *Table) InsertRows(values ...any) error {
	if len(values) != len(t.Columns) {
		return fmt.Errorf("expected %d values, got %d", len(t.Columns), len(values))
	}

	for i, v := range values {
		switch t.Columns[i].Type {
		case Int:
			if _, ok := v.(int); !ok {
				return fmt.Errorf("column %s expects INT", t.Columns[i].Name)
			}
		case String:
			if _, ok := v.(string); !ok {
				return fmt.Errorf("column %s expects STRING", t.Columns[i].Name)
			}
		}
	}
	t.Rows = append(t.Rows, values)
	return nil
}

func (db *Database) PrintTable(name string) error {
	t, exists := db.Tables[name]
	if !exists {
		return errors.New("table doesnt exist")
	}
	const padding = 3
	w := tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', 0)
	defer w.Flush()

	// headers
	for _, c := range t.Columns {
		fmt.Fprintf(w, "%s\t", c.Name)
	}
	fmt.Fprintln(w)

	// rows
	for _, r := range t.Rows {
		for _, v := range r {
			fmt.Fprintf(w, "%v\t", v)
		}
		fmt.Fprintln(w)
	}

	return nil
}
