package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/vitormelon/transactions/models"
	"net/http"
)

func FindAccount(c *gin.Context) {
	id := c.Param("accountId")

	var account models.Account
	err := models.DB.First(&account, id).Error

	if err != nil {
		//TODO: colocar log
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, account)
}

func CreateAccount(c *gin.Context) {
	var input CreateAccountInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		//TODO: colocar log
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	account := models.Account{
		DocumentNumber: input.DocumentNumber,
	}
	models.DB.Create(&account)

	c.JSON(http.StatusOK, account)
}

type CreateAccountInput struct {
	DocumentNumber string `json:"document_number" binding:"required"`
}