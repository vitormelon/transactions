package models

type Account struct {
	ID             uint `json:"id" gorm:"primaryKey;autoIncrement"`
	DocumentNumber string
}
