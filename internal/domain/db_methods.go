package domain

import (
	"errors"
	"fmt"
	"os"
	"text/tabwriter"
)

type DB interface {
	CreateTable(name string, cols []Column) error
	PrintTable(name string) error
	InsertRows(tableName string, values ...any) error
	GetRowByID(tableName string, id int) (Row, error)
}

func NewDatabase(name string) *Database {
	return &Database{
		Name:   name,
		Tables: make(map[string]*Table),
	}
}

func tableExists(m map[string]*Table, key string) bool {
	if _, yes := m[key]; yes {
		return true
	}
	return false
}

func (db *Database) CreateTable(name string, cols []Column) error {
	if tableExists(db.Tables, name) {
		return errors.New("table already exists")
	}

	t := &Table{
		Name:    name,
		Columns: cols,
		Rows:    []Row{},
	}
	db.Tables[name] = t

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

func (db *Database) InsertRows(tableName string, values ...any) error {
	if !tableExists(db.Tables, tableName) {
		return errors.New("table doesnt exist")
	}

	cols := db.Tables[tableName].Columns
	if len(values) != len(cols) {
		return fmt.Errorf("expected %d values, got %d", len(cols), len(values))
	}

	for i, v := range values {
		switch cols[i].Type {
		case Int:
			if _, ok := v.(int); !ok {
				return fmt.Errorf("column %s expects INT", cols[i].Name)
			}
		case String:
			if _, ok := v.(string); !ok {
				return fmt.Errorf("column %s expects STRING", cols[i].Name)
			}
		}
	}
	db.Tables[tableName].Rows = append(db.Tables[tableName].Rows, values)
	return nil
}

func (db *Database) GetRowByID(tableName string, id int) (Row, error) {
	if !tableExists(db.Tables, tableName) {
		return nil, errors.New("table doesnt exist")
	}
	for i := 0; i < len(db.Tables[tableName].Rows); i++ {
		if id == db.Tables[tableName].Rows[i][0] {
			return db.Tables[tableName].Rows[i], nil
		}
	}

	return nil, fmt.Errorf("row with id %d doesnt exist", id)
}

//func (db *Database) Filter() (Row, error)
// func (db *Database) DeleteRow(name string) (string, error)
