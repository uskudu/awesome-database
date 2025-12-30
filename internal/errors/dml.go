package errors

import (
	"errors"
)

var DBNameOccupied = errors.New("database name is occupied")
