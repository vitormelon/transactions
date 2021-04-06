package models

type OperationTypes struct {
	ID          uint `json:"id" gorm:"primaryKey;autoIncrement"`
	Description string
}

