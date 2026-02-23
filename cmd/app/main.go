package main

import (
	"awesomeDatabase/internal/domain"
	"fmt"
)

func main() {
	db := domain.NewDatabase("business")

	_ = db.CreateTable(
		"horses", []domain.Column{
			{Name: "id", Type: domain.Int},
			{Name: "name", Type: domain.String},
			{Name: "age", Type: domain.Int},
		})

	_ = db.InsertRows("horses", 0, "henry", 3)
	_ = db.InsertRows("horses", 1, "brain", 2)
	_ = db.InsertRows("horses", 2, "herman", 4)
	_ = db.InsertRows("horses", 3, "hosas", 4)

	//_ = db.PrintTable("names")

	d, _ := db.GetRowByID("horses", 0)
	fmt.Println(d)
}
