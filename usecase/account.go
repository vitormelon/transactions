package usecase

import "github.com/vitormelon/transactions/domain"

type AccountUseCase struct {
	accountRepository domain.AccountRepository
	accountPresenter domain.AccountPresenter
}

func (accountUseCase *AccountUseCase) Create(accountInput domain.AccountInput) (domain.Account, error) {
	account := domain.Account{
		DocumentNumber: accountInput.DocumentNumber,
	}
	err := accountUseCase.accountRepository.Create(&account)
	return account, err
}

func (accountUseCase *AccountUseCase) Find(id int) (domain.AccountOutput, error) {
	account, err := accountUseCase.accountRepository.Find(id)
	output := accountUseCase.accountPresenter.Output(account)
	return output, err
}

func NewAccountUseCase(repository domain.AccountRepository, presenter domain.AccountPresenter) domain.AccountUseCase {
	return &AccountUseCase{accountRepository: repository, accountPresenter: presenter}
}


