package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/vitormelon/transactions/models"
	"github.com/vitormelon/transactions/repository/mysql"
	"gorm.io/gorm/clause"
	"net/http"
	"time"
)

func FindTransaction(c *gin.Context) {
	id := c.Param("transactionId")

	var transaction models.Transaction
	err := mysql.DB.Preload(clause.Associations).First(&transaction, id).Error

	if err != nil {
		//TODO: colocar log
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, transaction)
}

func CreateTransaction(c *gin.Context) {
	var input CreateTransactionInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		//TODO: colocar log
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	account := models.Transaction{
		AccountID:   input.AccountId,
		OperationID: input.OperationTypeId,
		Amount:      input.Amount,
		CreatedAt:   time.Time{},
	}
	mysql.DB.Preload(clause.Associations).Create(&account)

	c.JSON(http.StatusOK, account)
}

type CreateTransactionInput struct {
	AccountId uint `json:"account_id" binding:"required"`
	OperationTypeId uint `json:"operation_type_id" binding:"required"`
	Amount float32 `json:"amount" binding:"required"`
}