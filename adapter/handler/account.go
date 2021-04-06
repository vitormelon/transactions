package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/vitormelon/transactions/domain"
	"net/http"
	"strconv"
)

type AccountHandler struct {
	AccountUseCase domain.AccountUseCase
}

func (accountHandler AccountHandler) Find(context *gin.Context) {
	ID, err :=  strconv.Atoi(context.Param("accountId"))
	if err != nil {
		//TODO: add log
		context.JSON(http.StatusBadRequest, gin.H{"error" : "Invalid account ID"})
		return
	}

	account, err := accountHandler.AccountUseCase.Find(ID)
	if err != nil {
		//TODO: add log
		context.JSON(http.StatusNotFound, gin.H{"error" : "Account not found"})
		return
	}

	context.JSON(http.StatusOK, account)
}

func (accountHandler AccountHandler) Create(context *gin.Context) {
	panic("implement-me")
}

func NewAccountHandler(engine *gin.Engine, AccountUseCase domain.AccountUseCase)  {
	handler := &AccountHandler{AccountUseCase: AccountUseCase}
	engine.GET("/accounts/:accountId", handler.Find)
	engine.POST("/accounts", handler.Create)
}
