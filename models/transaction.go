package models

import (
	"time"
)

type Transaction struct {
	ID          uint `json:"id" gorm:"primaryKey;autoIncrement"`
	AccountId   uint
	OperationId uint
	Ammount     float64
	CreatedAt   time.Time
}
