package mysql

import (
	"github.com/vitormelon/transactions/domain"
	"gorm.io/gorm"
)

type mysqlAccountRepository struct {
	DB *gorm.DB
}

func (mysqlAccountRepository mysqlAccountRepository) Create(account *domain.Account) error {
	err := mysqlAccountRepository.DB.Create(&account)
	return err.Error
}

func (mysqlAccountRepository *mysqlAccountRepository) Find(id int) (domain.Account, error) {
	var account domain.Account
	err := mysqlAccountRepository.DB.First(&account, id)
	return account, err.Error
}

func NewMysqlAccountRepository(DB *gorm.DB) domain.AccountRepository {
	return &mysqlAccountRepository{DB}
}



