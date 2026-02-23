package main

import (
	"awesomeDatabase/internal/domain"
)

func main() {
	db := domain.NewDatabase("horses")

	_ = db.CreateTable(
		"names", []domain.Column{
			{Name: "id", Type: domain.Int},
			{Name: "name", Type: domain.String},
		})

	_ = db.InsertRows("names", 0, "henry")
	_ = db.InsertRows("names", 1, "brain")
	_ = db.InsertRows("names", 2, "herman")
	_ = db.InsertRows("names", 3, "hosas")

	_ = db.PrintTable("names")
}
