package domain

type TransactionLevel string
type ReadUncommitedLevel TransactionLevel
type ReadCommitedLevel TransactionLevel
type RepeatableReadLevel TransactionLevel
type SerializableLevel TransactionLevel

type Transaction struct {
}

//func (t *Transaction) ReadUncommited() (*Transaction, error){
//}
//
//func (t *Transaction) ReadCommited() (*Transaction, error){
//}
//
//func (t *Transaction) RepeatableRead() (*Transaction, error){
//}
//
//func (t *Transaction) Serializable() (*Transaction, error){
//}
