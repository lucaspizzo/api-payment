package migrations

import (
	"github.com/jinzhu/gorm"
	"github.com/lucaspizzo/api-payment/domains"
	"gopkg.in/gormigrate.v1"
	"time"
)

var migration201903041722 = gormigrate.Migration{
	ID: "201903041722",
	Migrate: func(tx *gorm.DB) error {

		type Account struct {
			AccountID                uint64 `sql:"primary_key;"`
			AvailableCreditLimit     float64
			AvailableWithdrawalLimit float64
		}

		type OperationType struct {
			OperationTypeID uint64 `sql:"primary_key;"`
			Description     domains.OperationTypeDescription
			ChargeOrder     uint64
		}

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

		tx.AutoMigrate(
			&Account{},
			&OperationType{},
			&Transaction{},
		)

		operationTypes := []OperationType{
			{
				Description: domains.OperationTypePurchaseAtSight,
				ChargeOrder: 2,
			}, {
				Description: domains.OperationTypePurchaseInInstallment,
				ChargeOrder: 1,
			}, {
				Description: domains.OperationTypeBankDraft,
				ChargeOrder: 0,
			}, {
				Description: domains.OperationTypePayment,
				ChargeOrder: 0,
			},
		}

		for _, operationType := range operationTypes {
			tx.Save(&operationType)
		}

		// FK
		tx.Exec("alter table transaction add constraint FK_TRANSACTION_ACCOUNT foreign key (account_id) references account(account_id)")
		tx.Exec("alter table transaction add constraint FK_TRANSACTION_OPERATION_TYPE foreign key (operation_type_id) references operation_type(operation_type_id)")

		// INDEX
		tx.Exec("create index if not exists idx_operation_type_description on operation_type (description)")

		return nil
	},
	Rollback: func(tx *gorm.DB) error {
		return nil
	},
}
