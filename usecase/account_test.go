package usecase

import (
	"github.com/vitormelon/transactions/adapter/presenter"
	"github.com/vitormelon/transactions/domain"
	"testing"
)

type accountRepositoryMock struct {}

var documentNumber = "123456789"

func (accountRepositoryMock accountRepositoryMock) Find(id int) (domain.Account, error) {
	return domain.Account{
		ID:             uint(id),
		DocumentNumber: documentNumber,
	}, nil
}

func (accountRepositoryMock accountRepositoryMock) Create(account *domain.Account) error {
	account.ID = 1
	return nil
}


func TestAccountUseCase_Find(t *testing.T) {
	accountUseCase := getAccountUseCase()
	id := 10
	account, err := accountUseCase.Find(id)
	if err != nil {
		t.Error("No error should occur")
	}

	if account.ID != uint(id) {
		t.Errorf("Id was incorrect, got: %d, want: %d.", account.ID, id)
	}

	if account.DocumentNumber != documentNumber {
		t.Errorf("docummentNumber was incorrect, got: %s, want: %s.", account.DocumentNumber, documentNumber)
	}

}

func getAccountUseCase() domain.AccountUseCase {
	accountPresenter := presenter.NewAccountPresenter()
	accountUseCase := NewAccountUseCase(&accountRepositoryMock{}, accountPresenter)
	return accountUseCase
}

