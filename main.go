package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vitormelon/transactions/adapter/handler"
	"github.com/vitormelon/transactions/adapter/presenter"
	"github.com/vitormelon/transactions/repository/mysql"
	"github.com/vitormelon/transactions/usecase"
	"net/http"
)

func main() {
	engine := gin.Default()

	engine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	mysqlRepository := mysql.ConnectDataBase()
	accountRepository := mysql.NewMysqlAccountRepository(mysqlRepository)
	accountPresenter := presenter.NewAccountPresenter()
	accountUseCase := usecase.NewAccountUseCase(accountRepository, accountPresenter)

	transactionRepository := mysql.NewMysqlTransactionRepository(mysqlRepository)
	transactionPresenter := presenter.NewTransactionPresenter()
	transactionUseCase := usecase.NewTransactionUseCase(transactionRepository, transactionPresenter)

	handler.NewTransactionHandler(engine, transactionUseCase)
	handler.NewAccountHandler(engine, accountUseCase)

	engine.Run()
}
