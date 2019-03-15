package domains

import "time"

type Transaction struct {
	TransactionID uint64 `sql:"primary_key;"`

	AccountID uint64  `json:"account_id"`
	Account   Account `json:"-"`

	OperationTypeID uint64        `json:"operation_type_id"`
	OperationType   OperationType `gorm:"foreignkey:OperationTypeID;association_foreignkey:OperationTypeID" json:"-"`

	Amount  float64 `json:"amount"`
	Balance float64 `json:"balance"`

	EventDate time.Time `json:"event_date"`
	DueDate   time.Time `json:"due_date"`
}
