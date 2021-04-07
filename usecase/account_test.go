package usecase

import (
	"github.com/vitormelon/transactions/adapter/presenter"
	"github.com/vitormelon/transactions/domain"
	"testing"
)

type accountRepositoryMock struct {}

var defaultDocumentNumber = "123456789"
var defaultId = uint(123)

func (accountRepositoryMock accountRepositoryMock) Find(id int) (domain.Account, error) {
	return domain.Account{
		ID:             uint(id),
		DocumentNumber: defaultDocumentNumber,
	}, nil
}

func (accountRepositoryMock accountRepositoryMock) Create(account *domain.Account) error {
	account.ID = defaultId
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

	if account.DocumentNumber != defaultDocumentNumber {
		t.Errorf("docummentNumber was incorrect, got: %s, want: %s.", account.DocumentNumber, defaultDocumentNumber)
	}
}

func TestAccountUseCase_Create(t *testing.T) {
	accountUseCase := getAccountUseCase()
	documentNumberTest := "321654987"
	accountInput := domain.AccountInput{
		DocumentNumber: documentNumberTest,
	}
	account, err := accountUseCase.Create(accountInput)
	if err != nil {
		t.Error("No error should occur")
	}

	if account.ID != defaultId {
		t.Errorf("Id was incorrect, got: %d, want: %d.", account.ID, defaultId)
	}

	if account.DocumentNumber != documentNumberTest {
		t.Errorf("docummentNumber was incorrect, got: %s, want: %s.", account.DocumentNumber, documentNumberTest)
	}
}

func getAccountUseCase() domain.AccountUseCase {
	accountPresenter := presenter.NewAccountPresenter()
	accountUseCase := NewAccountUseCase(&accountRepositoryMock{}, accountPresenter)
	return accountUseCase
}

