package models

import (
	"time"
)

type Transaction struct {
	ID        uint `json:"id" gorm:"primaryKey;autoIncrement"`
	AccountID uint
	Account   Account
	OperationID uint
	Operation OperationTypes
	Amount    float32
	CreatedAt time.Time
}
