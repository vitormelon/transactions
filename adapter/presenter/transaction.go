package presenter

import "github.com/vitormelon/transactions/domain"

type transactionPresenter struct {}

func (transactionPresenter transactionPresenter) Output(transaction domain.Transaction) domain.TransactionOutput {
	return domain.TransactionOutput{
		Amount:          transaction.Amount,
		AccountID:       transaction.AccountID,
		OperationTypeID: transaction.OperationTypeID,
	}
}

func NewTransactionPresenter() domain.TransactionPresenter {
	return &transactionPresenter{}
}
