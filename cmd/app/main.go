package main

import (
	"awesomeDatabase/internal/domain"
)

func main() {
	db := domain.NewDatabase("company")

	usersTable, _ := db.CreateTable(
		"users", []domain.Column{
			{Name: "id", Type: domain.Int},
			{Name: "name", Type: domain.String},
		})

	_ = usersTable.InsertRows(0, "henry")
	_ = usersTable.InsertRows(1, "brain")
	_ = usersTable.InsertRows(1, "herman")
	_ = usersTable.InsertRows(1, "hosas")

	db.PrintTable("users")
}
