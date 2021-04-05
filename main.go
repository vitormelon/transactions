package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vitormelon/transactions/controllers"
	"github.com/vitormelon/transactions/models"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	r.POST("/accounts", controllers.CreateAccount)
	r.GET("/accounts/:accountId", controllers.FindAccount)
	r.POST("/transactions", controllers.CreateTransaction)
	r.GET("/transactions/:transactionId", controllers.FindTransaction)

	models.ConnectDataBase()
	models.InsertDefaultOperationTypes()

	r.Run()
}
