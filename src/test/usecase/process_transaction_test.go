package usecase

import (
	"github.com/a2martins/golang-tdd-clean-architecture/src/main/usecase"
	"github.com/a2martins/golang-tdd-clean-architecture/src/main/usecase/dto"
	mock_entity "github.com/a2martins/golang-tdd-clean-architecture/src/test/entity/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProcessTransactionWhenItsValid(t *testing.T) {
	input := dto.TransactionDTOInput{
		ID:        "1",
		AccountID: "1",
		Amount:    200,
	}

	expectedOutput := dto.TransactionDTOOutput{
		ID:           "1",
		Status:       "approved",
		ErrorMessage: "",
	}

	ctrl := gomock.NewController(t)
	repository := mock_entity.NewMockTransactionRepository(ctrl)
	repository.EXPECT().Insert(
		input.ID,
		input.AccountID,
		input.Amount,
		"approved",
		"",
	).Return(nil)
	defer ctrl.Finish()

	processTransaction := usecase.NewProcessTransaction(repository)
	output, err := processTransaction.Execute(input)

	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)
}
