package presenter

import (
	"github.com/vitormelon/transactions/domain"
)

type accountPresenter struct {}

func (accountPresenter accountPresenter) Output(account domain.Account) domain.AccountOutput {
	return domain.AccountOutput{
		ID:             account.ID,
		DocumentNumber: account.DocumentNumber,
	}
}

func NewAccountPresenter() domain.AccountPresenter {
	return accountPresenter{}
}