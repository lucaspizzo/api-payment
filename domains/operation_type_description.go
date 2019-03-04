package domains

import "github.com/pkg/errors"

type OperationTypeDescription string

const (
	OperationTypePurchaseAtSight       OperationTypeDescription = "COMPRA A VISTA"
	OperationTypePurchaseInInstallment OperationTypeDescription = "COMPRA PARCELADA"
	OperationTypeBankDraft             OperationTypeDescription = "SAQUE"
	OperationTypePayment               OperationTypeDescription = "PAGAMENTO"
)

func (*OperationTypeDescription) ValueOfOperationTypeDescription(value string) (OperationTypeDescription, error) {
	switch value {
	case "COMPRA A VISTA":
		return OperationTypePurchaseAtSight, nil
	case "COMPRA PARCELADA":
		return OperationTypePurchaseInInstallment, nil
	case "SAQUE":
		return OperationTypeBankDraft, nil
	case "PAGAMENTO":
		return OperationTypePayment, nil
	default:
		var operationTypeDescription OperationTypeDescription
		return operationTypeDescription, errors.New("Operation not found")
	}
}

func (o OperationTypeDescription) Name() string {
	switch o {
	case OperationTypePurchaseAtSight:
		return "COMPRA A VISTA"
	case OperationTypePurchaseInInstallment:
		return "COMPRA PARCELADA"
	case OperationTypeBankDraft:
		return "SAQUE"
	case OperationTypePayment:
		return "PAGAMENTO"
	default:
		return ""
	}
}

