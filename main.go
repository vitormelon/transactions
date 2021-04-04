package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vitormelon/transactions/models"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	models.ConnectDataBase()

	// transaction := models.Transaction{
	// 	AccountId:   10,
	// 	OperationId: 2,
	// 	Ammount:     2.53,
	// }
	// models.DB.Create(&transaction)

	r.Run()
}
