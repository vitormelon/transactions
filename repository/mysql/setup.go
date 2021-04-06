package mysql

import (
	"fmt"
	"github.com/vitormelon/transactions/domain"
	"github.com/vitormelon/transactions/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func ConnectDataBase() *gorm.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DATABASE"),
	)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	//TODO: add log
	if err != nil {
		panic("Failed to connect to database!")
	}

	//TODO: remover isso
	database.AutoMigrate(&models.Transaction{}, &domain.Account{}, &models.OperationTypes{})

	return database
}
