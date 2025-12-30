package main

import (
	"awesomeDatabase/internal/domain"
	"awesomeDatabase/internal/interface"
	"awesomeDatabase/internal/utils"
)

func main() {
	t := _interface.CreateTable(
		"users",
		domain.Field{
			Name:   utils.GetPtr("id"),
			Type:   utils.GetPtr("int"),
			Values: []any{1, 2, 3},
		},
		domain.Field{
			Name:   utils.GetPtr("name"),
			Type:   utils.GetPtr("string"),
			Values: []any{"alice", "john", "david"},
		},
		domain.Field{
			Name:   utils.GetPtr("email"),
			Type:   utils.GetPtr("string"),
			Values: []any{"alice@mail.com", "john22011@mail.com", "supersmartdavid@mail.com"},
		},
	)
	_interface.ShowTable(&t)
}
