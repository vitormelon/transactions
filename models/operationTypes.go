package models

type OperationTypes struct {
	ID          uint `json:"id" gorm:"primaryKey;autoIncrement"`
	Description string
}

func InsertDefaultOperationTypes() {
	insertOperationType("COMPRA A VISTA")
	insertOperationType("COMPRA PARCELADA")
	insertOperationType("SAQUE")
	insertOperationType("PAGAMENTO")
}

func insertOperationType(description string) {
	operationType := OperationTypes{
		Description: description,
	}
	DB.Create(&operationType)
}

