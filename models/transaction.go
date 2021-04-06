package models

import (
	"github.com/vitormelon/transactions/domain"
	"time"
)

type Transaction struct {
	ID          uint `json:"id" gorm:"primaryKey;autoIncrement"`
	AccountID   uint
	Account     domain.Account
	OperationID uint
	Operation   OperationTypes
	Amount      float32
	CreatedAt   time.Time
}
