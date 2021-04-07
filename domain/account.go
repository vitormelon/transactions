package domain

import "github.com/gin-gonic/gin"


type (
	Account struct {
		ID             uint `gorm:"primaryKey;autoIncrement"`
		DocumentNumber string
	}

	AccountOutput struct {
		ID uint `json:"id"`
		DocumentNumber string `json:"document_number"`
	}

	AccountInput struct {
		DocumentNumber string `json:"document_number" binding:"required"`
	}

	AccountRepository interface {
		Create(account *Account) error
		Find(id int) (Account, error)
	}

	AccountUseCase interface {
		Create(account AccountInput) (Account, error)
		Find(id int) (AccountOutput, error)
	}

	AccountHandler interface {
		Create(c *gin.Context)
		Find(c *gin.Context)
	}

	AccountPresenter interface {
		Output(account Account) AccountOutput
	}
)
