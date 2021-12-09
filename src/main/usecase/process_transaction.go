package usecase

import (
	"github.com/a2martins/golang-tdd-clean-architecture/src/main/entity"
	"github.com/a2martins/golang-tdd-clean-architecture/src/main/usecase/dto"
)

type ProcessTransaction struct {
	Repository entity.TransactionRepository
}

func NewProcessTransaction(repository entity.TransactionRepository) *ProcessTransaction {
	return &ProcessTransaction{Repository: repository}
}

func (p *ProcessTransaction) Execute(input dto.TransactionDTOInput) (dto.TransactionDTOOutput, error) {
	transaction := entity.NewTransaction()
	transaction.ID = input.ID
	transaction.AccountID = input.AccountID
	transaction.Amount = input.Amount
	transaction.Status = "approved"
	isValidTransaction := transaction.IsValid()

	if isValidTransaction != nil {
		transaction.Status = "rejected"
		transaction.ErrorMessage = isValidTransaction.Error()
	}

	return p.insertTransaction(transaction)

}

func (p *ProcessTransaction) insertTransaction(transaction *entity.Transaction) (dto.TransactionDTOOutput, error) {
	err := p.Repository.Insert(
		transaction.ID,
		transaction.AccountID,
		transaction.Amount,
		transaction.Status,
		transaction.ErrorMessage,
	)
	if err != nil {
		return dto.TransactionDTOOutput{}, err
	}

	output := dto.TransactionDTOOutput{
		ID:           transaction.ID,
		Status:       transaction.Status,
		ErrorMessage: transaction.ErrorMessage,
	}

	return output, nil
}
