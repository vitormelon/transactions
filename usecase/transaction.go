package usecase

import (
	"github.com/vitormelon/transactions/domain"
)

type transactionUseCase struct {
	transactionRepository domain.TransactionRepository
	transactionPresenter domain.TransactionPresenter
}

func (transactionUseCase transactionUseCase) Find(id int) (domain.TransactionOutput, error) {
	transaction, err := transactionUseCase.transactionRepository.Find(id)
	transactionOutput := transactionUseCase.transactionPresenter.Output(transaction)
	return transactionOutput, err
}

func (transactionUseCase transactionUseCase) Create(transactionInput domain.TransactionInput) (domain.TransactionOutput, error) {
	transaction := domain.Transaction{
		AccountID:       transactionInput.AccountID,
		OperationTypeID: transactionInput.OperationTypeID,
		Amount:          transactionInput.Amount,
	}
	err := transactionUseCase.transactionRepository.Create(&transaction)
	output := transactionUseCase.transactionPresenter.Output(transaction)
	return output, err
}

func NewTransactionUseCase (transactionRepository domain.TransactionRepository, transactionPresenter domain.TransactionPresenter) domain.TransactionUseCase {
	return &transactionUseCase{
		transactionRepository: transactionRepository,
		transactionPresenter:  transactionPresenter,
	}
}
