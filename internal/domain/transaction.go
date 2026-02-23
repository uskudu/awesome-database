package domain

type IsolationLevel int

const (
	ReadUncommited IsolationLevel = iota
	ReadCommitted
	RepeatableRead
	Serializable
)

func (l IsolationLevel) String() string {
	switch l {
	case ReadUncommited:
		return "READ UNCOMMITED"
	case ReadCommitted:
		return "READ COMMITTED"
	case RepeatableRead:
		return "REPEATABLE READ"
	case Serializable:
		return "SERIALIZABLE"
	default:
		return "UNKNOWN"
	}
}

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
