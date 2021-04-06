package usecase

import "github.com/vitormelon/transactions/domain"

type AccountUsecase struct {
	accountRepository domain.AccountRepository
	accountPresenter domain.AccountPresenter
}

func (accountUsecase *AccountUsecase) Create(account *domain.Account) error {
	panic("implement me")
}

func (accountUsecase *AccountUsecase) Find(id int) (domain.AccountOutput, error) {
	account, err := accountUsecase.accountRepository.Find(id)
	output := accountUsecase.accountPresenter.Output(account)
	return output, err
}

func NewAccountUsecase(repository domain.AccountRepository, presenter domain.AccountPresenter) domain.AccountUseCase {
	return &AccountUsecase{accountRepository: repository, accountPresenter: presenter}
}


