package dto

type TransactionDTOInput struct {
	ID        string
	AccountID string
	Amount    float64
}

type TransactionDTOOutput struct {
	ID           string
	Status       string
	ErrorMessage string
}
