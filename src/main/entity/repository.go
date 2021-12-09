package entity

type TransactionRepository interface {
	Insert(ID string, AccountID string, Amount float64, Status string, ErrorMessage string) error
}
