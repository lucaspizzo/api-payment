package domains

import "time"

type Transaction struct {
	TransactionID uint64 `sql:"primary_key;"`

	AccountID uint64
	Account   Account

	OperationTypeID uint64
	OperationType   OperationType

	Amount  float64
	Balance float64

	EventDate time.Time
	DueDate   time.Time
}
