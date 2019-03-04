package migrations

import (
	"github.com/jinzhu/gorm"
	"github.com/lucaspizzo/api-payment/domains"
	"gopkg.in/gormigrate.v1"
)

type Account struct {
	AccountID uint64 `sql:"primary_key;"`
	AvailableCreditLimit float64
	AvailableWithdrawalLimit float64
}

type OperationType struct {
	OperationTypeID uint64 `sql:"primary_key;"`
	Description domains.OperationTypeDescription
	ChargeOrder uint64
}


var migration201903041722 = gormigrate.Migration{
	ID: "201903041722",
	Migrate: func(tx *gorm.DB) error {
		tx.AutoMigrate(
			&Account{},
			&OperationType{},
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

		for _, operationType := range operationTypes{
			tx.Save(&operationType)
		}

		return nil
	},
	Rollback: func(tx *gorm.DB) error {
		return nil
	},
}