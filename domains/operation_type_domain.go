package domains

type OperationType struct {
	OperationTypeID uint64 `sql:"primary_key;"`
	Description     OperationTypeDescription
	ChargeOrder     uint64
}
