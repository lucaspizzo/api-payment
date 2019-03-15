package services

import (
	"github.com/lucaspizzo/api-payment/forms"
)

type PaymentContract interface {
	RegisterPayments(form *[]*forms.PaymentForm) error
}

type PaymentService struct {
	AccountService       AccountContract       `inject:""`
	OperationTypeService OperationTypeContract `inject:""`
	TransactionService   TransactionContract   `inject:""`
}

func (ps *PaymentService) RegisterPayments(form *[]*forms.PaymentForm) error {

	var operationTypeID uint64 = 4
	// TODO: IMPLEMENTAR ROLLBACK
	for _, payment := range *form {
		_, err := ps.TransactionService.process(payment.AccountID, operationTypeID, payment.Amount)
		if err != nil {
			return err
		}
	}

	return nil
}
