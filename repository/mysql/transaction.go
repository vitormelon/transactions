package mysql

import (
	"github.com/vitormelon/transactions/domain"
	"gorm.io/gorm"
)

type mysqlTransactionRepository struct {
	DB *gorm.DB
}

func (mysqlTransactionRepository mysqlTransactionRepository) Find(id int) (domain.Transaction, error) {
	var transaction domain.Transaction
	err := mysqlTransactionRepository.DB.First(&transaction, id)
	return transaction, err.Error
}

func (mysqlTransactionRepository mysqlTransactionRepository) Create(transaction *domain.Transaction) error {
	err := mysqlTransactionRepository.DB.Create(&transaction)
	return err.Error
}

func NewMysqlTransactionRepository (DB *gorm.DB) domain.TransactionRepository{
	return &mysqlTransactionRepository{DB}
}