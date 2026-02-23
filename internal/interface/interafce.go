package _interface

//type Database interface {
//	// ddl
//	CreateTable() error // create new table in this db
//	TruncateDB() error  // clear all db data
//	DropDB() error      // delete db
//
//	// tcl
//	BeginTransaction(level *domain.TransactionLevel) (*domain.Transaction, error)
//	Commit() error   // commit transaction
//	Rollback() error // cancel transaction
//	// todo add SavePoint
//}
//
//type Table interface {
//	CreateTable(name string, fields ...domain.Field) error // create new table
//	// todo del this Show()
//	Show()
//
//	// todo add Alter
//	TruncateTable() error
//	DropTable() error
//
//	// dml
//	Select() // get data from db
//	Insert() // add data to db
//	Update() // edit data in db
//	Delete() // delete data from db
//
//}
//
//type database struct {
//	Tables *[]domain.Table
//}

//func NewDatabase() Database {
//	return &database{}
//}
