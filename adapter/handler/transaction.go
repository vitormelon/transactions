package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/vitormelon/transactions/domain"
	"net/http"
	"strconv"
)

type transactionHandler struct {
	TransactionUseCase domain.TransactionUseCase
}

func (transactionHandler transactionHandler) Find(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("transactionId"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction ID"})
		return
	}

	transactionOutput, err := transactionHandler.TransactionUseCase.Find(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error" : "Transaction not found"})
		return
	}
	context.JSON(http.StatusOK, transactionOutput)
}

func (transactionHandler transactionHandler) Create(context *gin.Context) {
	var transactionInput domain.TransactionInput
	err := context.BindJSON(&transactionInput)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error" : "Invalid Parameters"})
		return
	}

	transactionOutput, err := transactionHandler.TransactionUseCase.Create(transactionInput)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, transactionOutput)
}

func NewTransactionHandler(engine *gin.Engine, useCase domain.TransactionUseCase) *transactionHandler {
	transactionHandler := &transactionHandler{useCase}
	engine.GET("/transactions/:transactionId", transactionHandler.Find)
	engine.POST("/transactions", transactionHandler.Create)
	return transactionHandler
}
