package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vitormelon/transactions/adapter/handler"
	"github.com/vitormelon/transactions/adapter/presenter"
	"github.com/vitormelon/transactions/models"
	"github.com/vitormelon/transactions/repository/mysql"
	"github.com/vitormelon/transactions/usecase"
	"gorm.io/gorm"
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
	accountUsecase := usecase.NewAccountUseCase(accountRepository, accountPresenter)

	handler.NewAccountHandler(engine, accountUsecase)
	//InsertDefaultOperationTypes(mysqlRepository)

	engine.Run()
}

func InsertDefaultOperationTypes(mysqlRepository *gorm.DB) {
	insertOperationType("COMPRA A VISTA", mysqlRepository)
	insertOperationType("COMPRA PARCELADA", mysqlRepository)
	insertOperationType("SAQUE", mysqlRepository)
	insertOperationType("PAGAMENTO", mysqlRepository)
}

func insertOperationType(description string, mysqlRepository *gorm.DB) {
	operationType := models.OperationTypes{
		Description: description,
	}
	mysqlRepository.Create(&operationType)
}
