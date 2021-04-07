package domain

import (
	"github.com/gin-gonic/gin"
	"time"
)

type Transaction struct {
	ID int
	AccountID int
	OperationTypeID int
	Amount float32
	CreatedAt   time.Time
	OperationType
	Account
}

type TransactionOutput struct {
	AccountID int `json:"account_id"`
	OperationTypeID int `json:"operation_type_id"`
	Amount float32 `json:"amount"`
}

type TransactionInput struct {
	AccountID int `json:"account_id" binding:"required"`
	OperationTypeID int `json:"operation_type_id"  binding:"required"`
	Amount float32 `json:"amount"  binding:"required"`
}

type TransactionRepository interface {
	Find(id int) (Transaction, error)
	Create(transaction *Transaction) error
}

type TransactionUseCase interface {
	Find(id int) (TransactionOutput, error)
	Create(transaction TransactionInput) (TransactionOutput, error)
}

type TransactionHandler interface {
	Find(context *gin.Context)
	Create(context gin.Context)
}

type TransactionPresenter interface {
	Output(transaction Transaction) TransactionOutput
}