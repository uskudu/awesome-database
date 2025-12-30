package _interface

import (
	"awesomeDatabase/internal/domain"
	"fmt"
	"os"

	"text/tabwriter"
)

type Database interface {
	// ddl
	//InitNewDB() (*domain.Database, error) // create new db
	//TruncateDB() error                    // clear all db data
	//DropDB() error                        // delete db

	//// tcl
	//BeginTransaction(level *domain.TransactionLevel) (*domain.Transaction, error)
	//Commit() error   // commit transaction
	//Rollback() error // cancel transaction
	//// todo add SavePoint
}

type TableIfc interface {
	CreateTable(name string, fields ...domain.Field) error // create new table
	// todo del this Show()
	Show()

	//// todo add Alter
	//TruncateTable() error
	//DropTable() error
	//
	//// dml
	//Select() // get data from db
	//Insert() // add data to db
	//Update() // edit data in db
	//Delete() // delete data from db
	//
}

//type database struct {
//	Tables *[]domain.Table
//}

//func NewDatabase() Database {
//	return &database{}
//}

func CreateTable(name string, fields ...domain.Field) domain.Table {
	var table domain.Table
	table.Name = name
	for _, f := range fields {
		table.Fields = append(table.Fields, f)
	}
	return table
}

func TableMetadata(t *domain.Table) {
	if t == nil {
		return
	}

	fmt.Println("Table:", t.Name)

	for _, f := range t.Fields {
		name := "<nil>"
		typ := "<nil>"

		if f.Name != nil {
			name = *f.Name
		}
		if f.Type != nil {
			typ = *f.Type
		}

		fmt.Printf("  Field: %s, Type: %s\n", name, typ)
	}
}

func ShowTable(t *domain.Table) {
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
