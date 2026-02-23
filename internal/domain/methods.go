package domain

import (
	"fmt"
	"os"
	"text/tabwriter"
)

type DB interface {
	CreateTable(name string, field ...Field) Table
}

func NewDatabase(name string) *Database {
	return &Database{Name: name}
}

func (db *Database) CreateTable(name string, fields ...Field) Table {
	var table Table
	table.Name = name
	for _, f := range fields {
		table.Fields = append(table.Fields, f)
	}
	return table
}

func ShowTable(t *Table) {
	const padding = 3
	w := tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', 0)
	defer w.Flush()

	// headers
	for _, f := range t.Fields {
		if f.Name == nil {
			fmt.Fprint(w, "<nil>\t")
		} else {
			fmt.Fprintf(w, "%s\t", *f.Name)
		}
	}
	fmt.Fprintln(w)

	if t.Fields[0].Values == nil {
		return
	}
	rows := len(t.Fields[0].Values)

	// rows
	for i := 0; i < rows; i++ {
		for _, f := range t.Fields {
			if f.Values == nil || i >= len(f.Values) {
				fmt.Fprint(w, "\t")
				continue
			}
			fmt.Fprintf(w, "%v\t", (f.Values)[i])
		}
		fmt.Fprintln(w)
	}
}
