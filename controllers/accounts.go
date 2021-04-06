package controllers
//
//import (
//	"github.com/gin-gonic/gin"
//	"github.com/vitormelon/transactions/domain"
//	"github.com/vitormelon/transactions/repository/mysql"
//	"net/http"
//)
//
//func FindAccount(c *gin.Context) {
//	id := c.Param("accountId")
//
//	var account domain.Account
//	err := mysql.DB.First(&account, id).Error
//
//	if err != nil {
//		//TODO: colocar log
//		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
//		return
//	}
//
//	c.JSON(http.StatusOK, account)
//}
//
//func CreateAccount(c *gin.Context) {
//	var input CreateAccountInput
//	err := c.ShouldBindJSON(&input)
//	if err != nil {
//		//TODO: colocar log
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//	account := domain.Account{
//		DocumentNumber: input.DocumentNumber,
//	}
//	mysql.DB.Create(&account)
//
//	c.JSON(http.StatusOK, account)
//}
//
//type CreateAccountInput struct {
//	DocumentNumber string `json:"document_number" binding:"required"`
//}